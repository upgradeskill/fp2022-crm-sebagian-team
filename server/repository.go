package server

import (
	"crm-sebagian-team/domain"
	_categoryRepo "crm-sebagian-team/modules/product/repository"
	_userRepo "crm-sebagian-team/modules/user/repository"
	"crm-sebagian-team/utils"
)

type Repository struct {
	UserRepo     domain.UserRepository
	CategoryRepo domain.CategoryRepository
}

// NewRepository will create an object that represent all repos interface
func NewRepository(conn *utils.Conn) *Repository {
	return &Repository{
		UserRepo:     _userRepo.NewUserRepository(conn.GORM),
		CategoryRepo: _categoryRepo.NewCategoryRepository(conn.GORM),
	}
}
