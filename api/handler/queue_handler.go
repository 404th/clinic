package handler

import (
	"net/http"

	"github.com/404th/clinic/model"
	"github.com/404th/clinic/pkg/helper"
	"github.com/gin-gonic/gin"
)

// Create CreateQueue
// @ID				create_queue
// @Security		ApiKeyAuth
// @Router			/queue [POST]
// @Summary			create queue
// @Description		create queue
// @Tags			queue
// @Accept			json
// @Produce			json
// @Param			data	body		model.CreateQueueRequest						true	"data"
// @Success			200		{object}	model.SuccessResponse{data=model.IDTracker}				"body"
// @Response		400		{object}	model.ErrorResponse{message=string}						"Invalid Argument"
// @Failure			500		{object}	model.ErrorResponse{message=string}						"Server Error"
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

// Create MakePurchase
// @ID				make_purchase
// @Security		ApiKeyAuth
// @Router			/queue [PATCH]
// @Summary			make purchase
// @Description		make purchase
// @Tags			queue
// @Accept			json
// @Produce			json
// @Param			data	body		model.MakePurchaseRequest						true	"data"
// @Success			200		{object}	model.SuccessResponse{data=model.IDTracker}				"body"
// @Response		400		{object}	model.ErrorResponse{message=string}						"Invalid Argument"
// @Failure			500		{object}	model.ErrorResponse{message=string}						"Server Error"
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
