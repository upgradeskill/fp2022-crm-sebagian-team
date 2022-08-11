package product

import (
	"crm-sebagian-team/helpers"

	"gorm.io/gorm"
)

func (Category) TableName() string {
	return "m_category"
}

type Category struct {
	ID        uint   `gorm:"primarykey"`
	Name      string `gorm:"column:name"`
	IsActive  bool   `gorm:"column:is_active"`
	CreatedAt helpers.Date
	CreatedBy string `gorm:"column:created_by"`
	UpdatedAt helpers.Date
	UpdatedBy string         `gorm:"column:updated_by"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	DeletedBy string         `gorm:"column:deleted_by"`
}
