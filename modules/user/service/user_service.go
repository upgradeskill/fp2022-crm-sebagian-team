package service

import (
	"context"
	"time"

	"crm-sebagian-team/domain"
)

type userService struct {
	userRepo       domain.UserRepository
	contextTimeout time.Duration
}

func NewUserService(br domain.UserRepository, timeout time.Duration) domain.UserService {
	return &userService{
		userRepo:       br,
		contextTimeout: timeout,
	}
}

func (bs *userService) Get(c context.Context, params *domain.Request) (res []domain.User, total int64, err error) {

	ctx, cancel := context.WithTimeout(c, bs.contextTimeout)
	defer cancel()

	res, total, err = bs.userRepo.Get(ctx, params)
	if err != nil {
		return nil, 0, err
	}

	return
}
