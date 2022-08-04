package server

import (
	"time"

	_userService "crm-sebagian-team/modules/user/service"
	"crm-sebagian-team/utils"

	"crm-sebagian-team/domain"
)

type Service struct {
	UserService domain.UserService
	AuthService domain.AuthService
}

func NewService(conn *utils.Conn, r *Repository, timeoutContext time.Duration) *Service {
	return &Service{
		// AuthService:    _authUcase.NewAuthService(cfg, r.UserRepo, r.UnitRepo, r.RoleRepo, r.RolePermRepo, timeoutContext),
		UserService: _userService.NewUserService(r.UserRepo, timeoutContext),
	}
}
