package dto

import (
	"fp2022-crm-sebagian-team/config"
)

type User struct {
	Id		 string `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"-" form:"password"`
}

var Users []User

func All() ([]User, error) {
	if err := config.DB.Find(&Users).Error; err != nil {
		return nil, err
	}
	return Users, nil
}

func GetUserById(id string) error {
	if err := config.DB.Where("id = ?", id).Take(&Users).Error; err != nil {
		return err
	}
	return nil
}