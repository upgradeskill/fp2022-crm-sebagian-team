package user

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

func (Position) TableName() string {
	return "m_position"
}

type Position struct {
	gorm.Model
	ID          int64        `gorm:"primary_key"`
	Name        string       `gorm:"column:name"`
	Description string       `gorm:"column:description"`
	CreatedAt   time.Time    `gorm:"column:created_at"`
	CreatedBy   string       `gorm:"column:created_by"`
	UpdatedAt   time.Time    `gorm:"column:updated_at"`
	UpdatedBy   string       `gorm:"column:updated_by"`
	DeletedAt   sql.NullTime `gorm:"column:deleted_at"`
	DeletedBy   string       `gorm:"column:deleted_by"`
}
