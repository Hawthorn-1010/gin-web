package model

import (
	"github.com/jinzhu/gorm"
)

//import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	//ID        uint   `json:"id" gorm:"primary key"`
	Name string `json:"name" gorm:"type:varchar(50);not null;unique"`
	//CreatedAt time.Time
	//UpdatedAt time.Time
}
