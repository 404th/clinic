package handler

import (
	"net/http"

	"github.com/404th/clinic/model"
	"github.com/404th/clinic/pkg/helper"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateRole(c *gin.Context) {
	var data model.CreateRoleRequest

	if err := c.ShouldBindJSON(&data); err != nil {
		helper.SendResponse(c, http.StatusBadRequest, model.ErrorResponse{
			Message: "Invalid arguments",
			Data:    err,
		})
		return
	}

	resp, err := h.service.RoleService().CreateRole(c.Request.Context(), &data)
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
