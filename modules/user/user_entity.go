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
	Name       string       `gorm:"column:name;not null;type:varchar(50);"`
	Email      string       `gorm:"column:email;not null;type:varchar(50);"`
	Password   string       `gorm:"column:password;not null;type:varchar(100);"`
	Address    string       `gorm:"column:address;not null;type:varchar(255);"`
	IdPosition int64        `gorm:"column:id_position;not null;"`
	CreatedAt  time.Time    `gorm:"column:created_at;not null;"`
	CreatedBy  string       `gorm:"column:created_by;not null;type:varchar(50);"`
	UpdatedAt  time.Time    `gorm:"column:updated_at;not null;"`
	UpdatedBy  string       `gorm:"column:updated_by;type:varchar(50);"`
	DeletedAt  sql.NullTime `gorm:"column:deleted_at"`
	DeletedBy  string       `gorm:"column:deleted_by;type:varchar(50);"`
}
