package handler

import (
	"net/http"

	"github.com/404th/clinic/model"
	"github.com/404th/clinic/pkg/helper"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateQueue(c *gin.Context) {
	var data model.CreateQueueRequest

	if err := c.ShouldBindJSON(&data); err != nil {
		helper.SendResponse(c, http.StatusBadRequest, model.ErrorResponse{
			Message: "Invalid arguments",
			Data:    err,
		})
		return
	}

	resp, err := h.service.QueueService().CreateQueue(c.Request.Context(), &data)
	if err != nil {
		helper.SendResponse(c, http.StatusInternalServerError, model.ErrorResponse{
			Message: err.Error(),
			Data:    err,
		})
		return
	}

	helper.SendResponse(c, http.StatusOK, model.SuccessResponse{
		Message: "OK",
		Data:    resp,
	})
}

func (h *Handler) MakePurchase(c *gin.Context) {
	var data model.MakePurchaseRequest

	if err := c.ShouldBindJSON(&data); err != nil {
		helper.SendResponse(c, http.StatusBadRequest, model.ErrorResponse{
			Message: "Invalid arguments",
			Data:    err,
		})
		return
	}

	if data.Amount <= 0 {
		helper.SendResponse(c, http.StatusBadRequest, model.ErrorResponse{
			Message: "Amount is not valid",
			Data:    nil,
		})
		return
	}

	if data.QueueID == "" {
		helper.SendResponse(c, http.StatusBadRequest, model.ErrorResponse{
			Message: "QueueID is not valid",
			Data:    nil,
		})
		return
	}

	resp, err := h.service.QueueService().MakePurchase(c.Request.Context(), &data)
	if err != nil {
		helper.SendResponse(c, http.StatusInternalServerError, model.ErrorResponse{
			Message: "error happened",
			Data:    err,
		})
		return
	}

	helper.SendResponse(c, http.StatusOK, model.SuccessResponse{
		Message: "OK",
		Data:    resp,
	})
	return
}
