package repository

import (
	"context"

	"crm-sebagian-team/domain"
	"crm-sebagian-team/helpers"
	"crm-sebagian-team/modules/product"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type productRepository struct {
	Conn *gorm.DB
}