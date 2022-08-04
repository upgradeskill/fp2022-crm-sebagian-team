package utils

import "gorm.io/gorm"

type Conn struct {
	GORM *gorm.DB
}
