package model

type CreateQueueRequest struct {
	RecipientID string `json:"recipient_id" binding:"required"`
	CustomerID  string `json:"customer_id" binding:"required"`
}

type MakePurchaseRequest struct {
	QueueID string  `json:"queue_id" binding:"required"`
	Amount  float64 `json:"amount" binding:"required"`
}
