package handler

import (
	"net/http"
	"strconv"

	"github.com/404th/clinic/model"
	"github.com/404th/clinic/pkg/helper"
	"github.com/gin-gonic/gin"
)

// Create Role
// @ID				create_role
// @Router			/role [POST]
// @Summary			create role
// @Description		create role
// @Tags			role
// @Accept			json
// @Produce			json
// @Param			data	body		model.CreateRoleRequest						true	"data"
// @Success			200		{object}	model.SuccessResponse{data=model.IDTracker}			"body"
// @Response		400		{object}	model.ErrorResponse{message=string}					"Invalid Argument"
// @Failure			500		{object}	model.ErrorResponse{message=string}					"Server Error"
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

// Get All Roles
// @ID				get_all_roles
// @Security		ApiKeyAuth
// @Router			/role [GET]
// @Summary			get_all_roles
// @Description		get_all_roles
// @Tags			role
// @Accept			json
// @Produce			json
// @Param     	 	limit   query	    string	  										false	"limit"
// @Param        	page    query     	string  										false  	"page"
// @Success			200		{object}	model.SuccessResponse{data=model.GetAllRolesResponse}	"body"
// @Response		400		{object}	model.ErrorResponse{message=string}						"Invalid Argument"
// @Failure			500		{object}	model.ErrorResponse{message=string}						"Server Error"
func (h *Handler) GetAllRoles(c *gin.Context) {
	var data model.GetAllRolesRequest

	limit_str := c.Query("limit")
	if limit_str != "" {
		limit, err := strconv.Atoi(limit_str)
		if err != nil {
			helper.SendResponse(c, http.StatusBadRequest, model.ErrorResponse{
				Message: "limit is not valid",
				Data:    err,
			})
			return
		}
		data.Limit = int32(limit)
	} else {
		data.Limit = int32(10)
	}

	page_str := c.Query("page")
	if page_str != "" {
		page, err := strconv.Atoi(page_str)
		if err != nil {
			helper.SendResponse(c, http.StatusBadRequest, model.ErrorResponse{
				Message: "page is not valid",
				Data:    err,
			})
			return
		}
		data.Page = int32(page)
	} else {
		data.Page = int32(1)
	}

	resp, err := h.service.RoleService().GetAllRoles(c.Request.Context(), &data)
	if err != nil {
		helper.SendResponse(c, http.StatusInternalServerError, model.ErrorResponse{
			Message: "error happened",
			Data:    err.Error(),
		})
		return
	}

	helper.SendResponse(c, http.StatusOK, model.SuccessResponse{
		Message: "OK",
		Data:    resp,
	})
	return
}
