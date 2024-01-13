package handler

import (
	"net/http"
	"strconv"

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

// Get All Queues
// @ID				get_all_queues
// @Security		ApiKeyAuth
// @Router			/queue [GET]
// @Summary			get_all_queues
// @Description		get_all_queues
// @Tags			queue
// @Accept			json
// @Produce			json
// @Param     	 	limit   query	    string	  										false	"limit"
// @Param        	page    query     	string  										false  	"page"
// @Success			200		{object}	model.SuccessResponse{data=model.GetAllQueuesResponse}	"body"
// @Response		400		{object}	model.ErrorResponse{message=string}						"Invalid Argument"
// @Failure			500		{object}	model.ErrorResponse{message=string}						"Server Error"
func (h *Handler) GetAllQueues(c *gin.Context) {
	var data model.GetAllQueuesRequest

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

	resp, err := h.service.QueueService().GetAllQueues(c.Request.Context(), &data)
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
