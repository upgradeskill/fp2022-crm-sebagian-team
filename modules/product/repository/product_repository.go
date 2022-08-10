package repository

import (
	"gorm.io/gorm"
)

type productRepository struct {
	Conn *gorm.DB
}
