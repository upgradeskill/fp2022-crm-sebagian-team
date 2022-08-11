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
	ID          int64        `gorm:"primary_key;autoIncrement;"`
	Name        string       `gorm:"column:name;type:varchar(15);not null;"`
	Description string       `gorm:"column:description;type:text;not null;"`
	CreatedAt   time.Time    `gorm:"column:created_at;not null;"`
	CreatedBy   string       `gorm:"column:created_by;type:varchar(50);not null;"`
	UpdatedAt   time.Time    `gorm:"column:updated_at;not null;"`
	UpdatedBy   string       `gorm:"column:updated_by;type:varchar(50);"`
	DeletedAt   sql.NullTime `gorm:"column:deleted_at"`
	DeletedBy   string       `gorm:"column:deleted_by;type:varchar(50);"`
}
