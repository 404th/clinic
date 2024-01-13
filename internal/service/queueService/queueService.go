package queueservice

import (
	"context"

	"github.com/404th/clinic/config"
	"github.com/404th/clinic/internal/service"
	"github.com/404th/clinic/internal/storage"
	"github.com/404th/clinic/model"
	"github.com/sirupsen/logrus"
)

type queueService struct {
	cfg  *config.Config
	log  *logrus.Logger
	strg storage.QueueI
}

func NewQueueService(cfg *config.Config, log *logrus.Logger, strg storage.QueueI) service.QueueServiceI {
	return &queueService{
		cfg:  cfg,
		log:  log,
		strg: strg,
	}
}

func (qs *queueService) CreateQueue(ctx context.Context, req *model.CreateQueueRequest) (resp *model.IDTracker, err error) {
	resp = &model.IDTracker{}

	qs.log.Infof("CreateQueue() => customer_id: %s => req: %+v", req.CustomerID, req)
	resp, err = qs.strg.CreateQueue(ctx, req)
	if err != nil {
		qs.log.Errorf("CreateQueue() => customer_id: %s => err: %+v", req.CustomerID, err)
		return resp, err
	}

	qs.log.Infof("CreateQueue() => id: %s => resp: %+v", resp.ID, resp)
	return resp, err
}

func (qs *queueService) MakePurchase(ctx context.Context, req *model.MakePurchaseRequest) (resp *model.IDTracker, err error) {
	resp = &model.IDTracker{}

	qs.log.Infof("MakePurchase() => queue_id: %s => req: %+v", req.QueueID, req)
	resp, err = qs.strg.MakePurchase(ctx, req)
	if err != nil {
		qs.log.Errorf("MakePurchase() => queue_id: %s => err: %+v", req.QueueID, err)
		return resp, err
	}

	qs.log.Infof("MakePurchase() => id: %s => resp: %+v", resp.ID, resp)
	return resp, err
}

func (qs *queueService) GetAllQueues(ctx context.Context, req *model.GetAllQueuesRequest) (resp *model.GetAllQueuesResponse, err error) {
	resp = &model.GetAllQueuesResponse{}

	qs.log.Infof("GetAllQueues() => page: %d => req: %+v", req.Page, req)
	resp, err = qs.strg.GetAllQueues(ctx, req)
	if err != nil {
		qs.log.Errorf("GetAllQueues() => page: %d => err: %+v", req.Page, err)
		return resp, err
	}

	qs.log.Infof("GetAllQueues() => metadata: %+v => resp: %+v", resp.Metadata, resp)
	return resp, err
}
