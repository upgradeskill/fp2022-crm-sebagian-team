package server

import (
	"time"

	"crm-sebagian-team/config"
	_authService "crm-sebagian-team/modules/auth/service"
	_categoryService "crm-sebagian-team/modules/product/service"
	_userService "crm-sebagian-team/modules/user/service"

	"crm-sebagian-team/utils"

	"crm-sebagian-team/domain"

	"github.com/go-playground/validator/v10"
)

type Service struct {
	UserService     domain.UserService
	AuthService     domain.AuthService
	CategoryService domain.CategoryService
	PositionService domain.PositionService
}

func NewService(cfg *config.Config, conn *utils.Conn, r *Repository, timeoutContext time.Duration) *Service {
	validate := validator.New()

	return &Service{
		AuthService:     _authService.NewAuthService(cfg, r.UserRepo, timeoutContext),
		UserService:     _userService.NewUserService(r.UserRepo, timeoutContext),
		CategoryService: _categoryService.NewCategoryService(r.CategoryRepo, validate, timeoutContext),
		PositionService: _userService.NewPositionService(r.PositionRepo, timeoutContext),
	}
}
