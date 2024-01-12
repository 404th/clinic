package handler

import (
	"net/http"

	"github.com/404th/clinic/model"
	"github.com/404th/clinic/pkg/helper"
	"github.com/404th/clinic/pkg/util"
	"github.com/gin-gonic/gin"
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

	if !util.IsEmail(data.Email) {
		helper.SendResponse(c, http.StatusBadRequest, model.ErrorResponse{
			Message: "Invalid email address",
			Data:    nil,
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
	}

	helper.SendResponse(c, http.StatusOK, model.SuccessResponse{
		Message: "Successfully created",
		Data:    resp,
	})
}

func (h *Handler) GetUserByID(c *gin.Context) {

}

func (h *Handler) UpdateUser(c *gin.Context) {
}

func (h *Handler) DeleteUser(c *gin.Context) {
}

func (h *Handler) TransferMoney(c *gin.Context) {
}
