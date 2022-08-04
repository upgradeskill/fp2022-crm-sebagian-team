package user

import "gorm.io/gorm"

func (User) TableName() string {
	return "m_user"
}

type User struct {
	gorm.Model
	Code        string `gorm:"column:code"`
	Name        string `gorm:"column:name"`
	Description string `gorm:"column:description"`
}
