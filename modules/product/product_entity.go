package product

import "gorm.io/gorm"

func (Product) TableName() string {
	return "m_product"
}

type Product struct {
	gorm.Model
	Name string `gorm:"column:name"`
	Qty  string `gorm:"column:qty"`
}
