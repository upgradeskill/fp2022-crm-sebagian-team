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

func beforeCreate(u *domain.UserRequest) (err error) {
	hashedPassword, errHash := encryptPassword(u.Password)
	if errHash != nil {
		return
	}
	u.Password = string(hashedPassword)
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return
}

func beforeUpdate(u *domain.UserUpdate) (err error) {
	hashedPassword, errHash := encryptPassword(u.Password)
	if errHash != nil {
		return
	}
	u.Password = string(hashedPassword)
	u.UpdatedAt = time.Now()
	return
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

func (svc *userService) GetByEmail(c context.Context, email string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, svc.contextTimeout)
	defer cancel()

	res, err := svc.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}

	return res, nil
}

func (svc *userService) Store(c context.Context, request *domain.UserRequest) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, svc.contextTimeout)
	defer cancel()

	err := beforeCreate(request)
	if err != nil {
		return domain.User{}, err
	}

	usr := domain.User{
		Name:       request.Name,
		Email:      request.Email,
		Password:   request.Password,
		Address:    request.Address,
		IdPosition: request.IdPosition,
		CreatedAt:  request.CreatedAt,
		CreatedBy:  request.CreatedBy,
		UpdatedAt:  request.UpdatedAt,
	}
	usrRes, err := svc.userRepo.Store(ctx, &usr)

	return usrRes, nil
}

func (svc *userService) GetByID(c context.Context, id int64) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, svc.contextTimeout)
	defer cancel()

	res, err := svc.userRepo.GetByID(ctx, id)
	if err != nil {
		return domain.User{}, err
	}

	return res, nil
}

func (svc *userService) UpdateUser(c context.Context, request *domain.UserUpdate) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, svc.contextTimeout)
	defer cancel()

	err := beforeUpdate(request)
	if err != nil {
		return domain.User{}, err
	}

	usr := domain.User{
		ID:         int64(request.ID),
		Name:       request.Name,
		Email:      request.Email,
		Password:   request.Password,
		Address:    request.Address,
		IdPosition: request.IdPosition,
		UpdatedAt:  request.UpdatedAt,
		UpdatedBy:  request.UpdatedBy,
	}
	usrRes, err := svc.userRepo.UpdateUser(ctx, &usr)

	return usrRes, nil
}

func (svc *userService) DeleteUser(c context.Context, id int64, deletedBy string) error {
	ctx, cancel := context.WithTimeout(c, svc.contextTimeout)
	defer cancel()

	err := svc.userRepo.DeleteUser(ctx, id, deletedBy)
	if err != nil {
		return err
	}

	return nil
}
