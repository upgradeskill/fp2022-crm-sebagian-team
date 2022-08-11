package user

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

func (User) TableName() string {
	return "m_user"
}

type User struct {
	gorm.Model
	ID         int64        `gorm:"primaryKey;autoIncrement"`
	Name       string       `gorm:"column:name"`
	Email      string       `gorm:"column:email"`
	Password   string       `gorm:"column:password"`
	Address    string       `gorm:"column:address"`
	IdPosition int64        `gorm:"column:id_position"`
	CreatedAt  time.Time    `gorm:"column:created_at"`
	CreatedBy  string       `gorm:"column:created_by"`
	UpdatedAt  time.Time    `gorm:"column:updated_at"`
	UpdatedBy  string       `gorm:"column:updated_by"`
	DeletedAt  sql.NullTime `gorm:"column:deleted_at"`
	DeletedBy  string       `gorm:"column:deleted_by"`
}
