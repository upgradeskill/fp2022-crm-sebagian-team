package product

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

func (Product) TableName() string {
	return "m_product"
}

type Product struct {
	gorm.Model
	ID        int64        `gorm:"primaryKey;autoIncrement"`
	Name      string       `gorm:"column:name"`
	Qty       int64        `gorm:"column:qty"`
	CreatedAt time.Time    `gorm:"column:created_at"`
	CreatedBy string       `gorm:"column:created_by"`
	UpdatedAt time.Time    `gorm:"column:updated_at"`
	UpdatedBy string       `gorm:"column:updated_by"`
	DeletedAt sql.NullTime `gorm:"column:deleted_at"`
	DeletedBy string       `gorm:"column:deleted_by"`
}
