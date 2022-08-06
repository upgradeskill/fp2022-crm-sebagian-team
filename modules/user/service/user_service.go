package service

import (
	"context"
	"time"

	"crm-sebagian-team/domain"

	"golang.org/x/crypto/bcrypt"
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

func encryptPassword(password string) (string, error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return "", err
	}

	return string(encryptedPassword), nil
}

func (svc *userService) Get(c context.Context, params *domain.Request) ([]domain.User, int64, error) {

	ctx, cancel := context.WithTimeout(c, svc.contextTimeout)
	defer cancel()

	res, total, err := svc.userRepo.Get(ctx, params)
	if err != nil {
		return []domain.User{}, 0, err
	}

	return res, total, nil
}

func (svc *userService) Store(c context.Context, usr *domain.User) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, svc.contextTimeout)
	defer cancel()

	encryptedPassword, err := encryptPassword(usr.Password)
	if err != nil {
		return domain.User{}, err
	}

	usr.Password = string(encryptedPassword)

	usrRes, err := svc.userRepo.Store(ctx, usr)

	return usrRes, nil
}
