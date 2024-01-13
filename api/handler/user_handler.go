package handler

import (
	"errors"
	"net/http"

	"github.com/404th/clinic/internal/jwt"
	"github.com/404th/clinic/model"
	"github.com/404th/clinic/pkg/helper"
	"github.com/404th/clinic/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Create User
// @ID				create_user
// @Router			/user [POST]
// @Summary			create user
// @Description		create user
// @Tags			user
// @Accept			json
// @Produce			json
// @Param			data	body		model.CreateUserRequest						true	"data"
// @Success			200		{object}	model.SuccessResponse{data=model.IDTracker}			"body"
// @Response		400		{object}	model.ErrorResponse{message=string}					"Invalid Argument"
// @Failure			500		{object}	model.ErrorResponse{message=string}					"Server Error"
func (h *Handler) CreateUser(c *gin.Context) {
	var data model.CreateUserRequest

	if err := c.ShouldBindJSON(&data); err != nil {
		helper.SendResponse(c, http.StatusBadRequest, model.ErrorResponse{
			Message: "Invalid arguments",
			Data:    err,
		})
		return
	}

	if !util.IsEmail(data.Email) {
		helper.SendResponse(c, http.StatusBadRequest, model.ErrorResponse{
			Message: "Invalid email address",
			Data:    nil,
		})
		return
	}

	hashed_password, err := util.HashPassword(data.Password)
	if err != nil {
		helper.SendResponse(c, http.StatusInternalServerError, model.ErrorResponse{
			Message: "Cannot hash the password",
			Data:    err,
		})
		return
	}

	data.Password = hashed_password

	resp, err := h.service.UserService().CreateUser(c.Request.Context(), &data)
	if err != nil {
		helper.SendResponse(c, http.StatusInternalServerError, model.ErrorResponse{
			Message: err.Error(),
			Data:    err,
		})
		return
	}

	helper.SendResponse(c, http.StatusOK, model.SuccessResponse{
		Message: "Successfully created",
		Data:    resp,
	})
}

// Login
// @ID				login
// @Router			/login [POST]
// @Summary			login
// @Description		login
// @Tags			user
// @Accept			json
// @Produce			json
// @Param			data	body		model.LoginRequest							true	"data"
// @Success			200		{object}	model.SuccessResponse{data=model.LoginResponse}		"body"
// @Response		400		{object}	model.ErrorResponse{message=string}					"Invalid Argument"
// @Failure			500		{object}	model.ErrorResponse{message=string}					"Server Error"
func (h *Handler) Login(c *gin.Context) {
	var data model.LoginRequest

	if err := c.ShouldBindJSON(&data); err != nil {
		helper.SendResponse(c, http.StatusBadRequest, model.ErrorResponse{
			Message: "Invalid arguments",
			Data:    err,
		})
		return
	}

	if data.Password == "" || data.Username == "" {
		helper.SendResponse(c, http.StatusBadRequest, model.ErrorResponse{
			Message: "Username and Password should be provided",
			Data:    nil,
		})
		return
	}

	resp, err := h.service.UserService().Login(c.Request.Context(), &data)
	if err != nil {
		helper.SendResponse(c, http.StatusInternalServerError, model.ErrorResponse{
			Message: err.Error(),
			Data:    err,
		})
		return
	}

	if resp.ID == "" {
		helper.SendResponse(c, http.StatusInternalServerError, model.ErrorResponse{
			Message: "user not found",
			Data:    errors.New("user not found"),
		})
		return
	}

	refresh_token, err := jwt.CreateRefreshToken(&model.RefreshTokenData{
		ID: resp.ID,
	}, h.cfg.RefreshTokenSecret, h.cfg.RefreshTokenExpiryHour)
	if err != nil {
		helper.SendResponse(c, http.StatusInternalServerError, model.ErrorResponse{
			Message: err.Error(),
			Data:    err,
		})
		return
	}

	access_token, err := jwt.CreateAccessToken(&model.AccessTokenData{
		ID:       resp.ID,
		Username: data.Username,
	}, h.cfg.AccessTokenSecret, h.cfg.AccessTokenExpiryHour)
	if err != nil {
		helper.SendResponse(c, http.StatusInternalServerError, model.ErrorResponse{
			Message: err.Error(),
			Data:    err,
		})
		return
	}

	resp.AccessToken = access_token
	resp.RefreshToken = refresh_token

	helper.SendResponse(c, http.StatusOK, model.SuccessResponse{
		Message: "Authorized",
		Data:    resp,
	})
}

// GetUserByID
// @ID				get_user_by_id
// @Router			/user/{id} [GET]
// @Summary			get_user_by_id
// @Description		get_user_by_id
// @Tags			user
// @Accept			json
// @Produce			json
// @Param			id		path		string										true	"id"
// @Success			200		{object}	model.SuccessResponse{data=model.User}				"user"
// @Response		400		{object}	model.ErrorResponse{message=string}					"Invalid Argument"
// @Failure			500		{object}	model.ErrorResponse{message=string}					"Server Error"
func (h *Handler) GetUserByID(c *gin.Context) {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		helper.SendResponse(c, http.StatusBadRequest, model.ErrorResponse{
			Message: "ID is invalid",
			Data:    nil,
		})
		return
	}

	usr, err := h.service.UserService().GetUserByID(c.Request.Context(), &model.IDTracker{
		ID: id,
	})
	if err != nil {
		helper.SendResponse(c, http.StatusBadRequest, model.ErrorResponse{
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	helper.SendResponse(c, http.StatusOK, model.SuccessResponse{
		Message: "OK",
		Data:    usr,
	})
}

// TransferMoney
// @ID				transfer
// @Security		ApiKeyAuth
// @Router			/user/transfer [PUT]
// @Summary			transfer
// @Description		transfer
// @Tags			user
// @Accept			json
// @Produce			json
// @Param			data	body		model.TransferMoneyRequest					true	"data"
// @Success			200		{object}	model.SuccessResponse{data=model.IDTracker}	"body"
// @Response		400		{object}	model.ErrorResponse{message=string}					"Invalid Argument"
// @Failure			500		{object}	model.ErrorResponse{message=string}					"Server Error"
func (h *Handler) TransferMoney(c *gin.Context) {
	var data model.TransferMoneyRequest
	if err := c.ShouldBindJSON(&data); err != nil {
		helper.SendResponse(c, http.StatusBadRequest, model.ErrorResponse{
			Message: "invalid arguments",
			Data:    err.Error(),
		})
		return
	}

	// ...1: getting id from url param
	if data.ID == "" {
		helper.SendResponse(c, http.StatusBadRequest, model.ErrorResponse{
			Message: "id is required",
			Data:    nil,
		})
		return
	}

	if _, err := uuid.Parse(data.ID); err != nil {
		helper.SendResponse(c, http.StatusBadRequest, model.ErrorResponse{
			Message: "id is invalid",
			Data:    nil,
		})
		return
	}

	resp, err := h.service.UserService().TransferMoney(c.Request.Context(), &data)
	if err != nil {
		helper.SendResponse(c, http.StatusInternalServerError, model.ErrorResponse{
			Message: "some error happened",
			Data:    err.Error(),
		})
		return
	}

	helper.SendResponse(c, http.StatusOK, model.SuccessResponse{
		Message: "OK",
		Data:    resp,
	})
}

// RefreshToken
// @ID				refresh-token
// @Security		ApiKeyAuth
// @Router			/refresh-token [PUT]
// @Summary			refresh-token
// @Description		refresh-token
// @Tags			user
// @Accept			json
// @Produce			json
// @Param			data	body		model.RefreshTokenRequest					true		"data"
// @Success			200		{object}	model.SuccessResponse{data=model.RefreshTokenResponse}	"body"
// @Response		400		{object}	model.ErrorResponse{message=string}						"Invalid Argument"
// @Failure			500		{object}	model.ErrorResponse{message=string}						"Server Error"
func (h *Handler) RefreshToken(c *gin.Context) {
	var request model.RefreshTokenRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		helper.SendResponse(c, http.StatusBadRequest, model.ErrorResponse{
			Message: "Invalid arguments",
			Data:    err,
		})
		return
	}

	id, err := jwt.ExtractIDFromToken(request.RefreshToken, h.cfg.RefreshTokenSecret)
	if err != nil {
		helper.SendResponse(c, http.StatusBadRequest, model.ErrorResponse{
			Message: "Invalid token",
			Data:    err,
		})
		return
	}

	user, err := h.service.UserService().GetUserByID(c, &model.IDTracker{ID: id})
	if err != nil {
		helper.SendResponse(c, http.StatusBadRequest, model.ErrorResponse{
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	accessToken, err := jwt.CreateAccessToken(&model.AccessTokenData{
		Username: user.Username,
		ID:       user.ID,
	}, h.cfg.AccessTokenSecret, h.cfg.AccessTokenExpiryHour)
	if err != nil {
		helper.SendResponse(c, http.StatusBadRequest, model.ErrorResponse{
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	refreshToken, err := jwt.CreateRefreshToken(&model.RefreshTokenData{ID: user.ID}, h.cfg.RefreshTokenSecret, h.cfg.RefreshTokenExpiryHour)
	if err != nil {
		helper.SendResponse(c, http.StatusBadRequest, model.ErrorResponse{
			Message: "ID is invalid",
			Data:    nil,
		})
		return
	}

	refreshTokenResponse := model.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	helper.SendResponse(c, http.StatusOK, model.SuccessResponse{
		Message: "Successfully created",
		Data:    refreshTokenResponse,
	})
}
