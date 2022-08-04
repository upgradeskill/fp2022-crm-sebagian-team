package server

import (
	"crm-sebagian-team/domain"
	_userRepo "crm-sebagian-team/modules/user/repository"
	"crm-sebagian-team/utils"
)

type Repository struct {
	UserRepo domain.UserRepository
}

// NewRepository will create an object that represent all repos interface
func NewRepository(conn *utils.Conn) *Repository {
	return &Repository{
		UserRepo: _userRepo.NewUserRepository(conn.GORM),
	}
}
