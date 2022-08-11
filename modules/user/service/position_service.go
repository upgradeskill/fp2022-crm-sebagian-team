package service

import (
	"context"
	"crm-sebagian-team/domain"
	"time"
)

type positionService struct {
	positionRepo   domain.PositionRepository
	contextTimeout time.Duration
}

func NewPositionService(br domain.PositionService, timeout time.Duration) domain.PositionService {
	return &positionService{
		positionRepo:   br,
		contextTimeout: timeout,
	}
}

func (svc *positionService) GetByID(c context.Context, id int64) (domain.Position, error) {
	ctx, cancel := context.WithTimeout(c, svc.contextTimeout)
	defer cancel()

	res, err := svc.positionRepo.GetByID(ctx, id)
	if err != nil {
		return domain.Position{}, err
	}

	return res, nil
}
