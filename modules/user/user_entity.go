package user

import "gorm.io/gorm"

func (User) TableName() string {
	return "m_user"
}

type User struct {
	gorm.Model
	Name     string `gorm:"column:name"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`
}
