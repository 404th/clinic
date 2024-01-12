package handler

import (
	"net/http"

	"github.com/404th/clinic/internal/jwt"
	"github.com/404th/clinic/model"
	"github.com/404th/clinic/pkg/helper"
	"github.com/404th/clinic/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

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

	c.SetCookie("refresh_token", refresh_token, h.cfg.RefreshTokenExpiryHour, "", "*", true, false)

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

	helper.SendResponse(c, http.StatusOK, model.SuccessResponse{
		Message: "Successfully created",
		Data:    resp,
	})
}

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
		Message: "Successfully created",
		Data:    usr,
	})
}

func (h *Handler) UpdateUser(c *gin.Context) {
}

func (h *Handler) DeleteUser(c *gin.Context) {
}

func (h *Handler) TransferMoney(c *gin.Context) {
}
