package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Customer struct {
	gorm.Model
	FirstName string
	LastName string
	Age int
	Email string
}

