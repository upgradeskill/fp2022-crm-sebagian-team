package product

import "gorm.io/gorm"

func (User) TableName() string {
	return "m_product"
}

type Product struct {
	gorm.Model
	Name	string `gorm:"column:name"`
	Qty		string `gorm:"column:qty"`
}
