package model

type CreateQueueRequest struct {
	RecipientID string `json:"recipient_id" binding:"required"`
	CustomerID  string `json:"customer_id" binding:"required"`
}

type MakePurchaseRequest struct {
	QueueID string  `json:"queue_id" binding:"required"`
	Amount  float64 `json:"amount" binding:"required"`
}

type GetAllQueuesRequest struct {
	Page  int32 `json:"page"`
	Limit int32 `json:"limit"`
}

type Queue struct {
	ID            string  `json:"id"`
	RecipientID   string  `json:"recipient_id"`
	CustomerID    string  `json:"customer_id"`
	PaidMoney     float64 `json:"paid_money"`
	QueueNumber   int     `json:"queue_number"`
	PaymentStatus int     `json:"payment_status"`
}

type Metadata struct {
	Limit int32 `json:"limit"`
	Page  int32 `json:"page"`
	Count int   `json:"count"`
}

type GetAllQueuesResponse struct {
	Metadata Metadata `json:"metadata"`
	Queues   []Queue  `json:"queues"`
}
