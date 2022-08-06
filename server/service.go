package server

import (
	"time"

	"crm-sebagian-team/config"
	_authService "crm-sebagian-team/modules/auth/service"
	_userService "crm-sebagian-team/modules/user/service"
	"crm-sebagian-team/utils"

	"crm-sebagian-team/domain"
)

type Service struct {
	UserService domain.UserService
	AuthService domain.AuthService
}

func NewService(cfg *config.Config, conn *utils.Conn, r *Repository, timeoutContext time.Duration) *Service {
	return &Service{
		AuthService: _authService.NewAuthService(cfg, r.UserRepo, timeoutContext),
		UserService: _userService.NewUserService(r.UserRepo, timeoutContext),
	}
}
