package csv

import (
	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	Name        string `gorm:"column:name"`
	LastName    string `gorm:"column:last_name"`
	Email       string `gorm:"column:email;uniqueIndex"`
	Phone       string `gorm:"column:phone"`
	City        string `gorm:"column:city"`
	Street      string `gorm:"column:street"`
	HouseNumber string `gorm:"column:house_number"`
	State       string `gorm:"column:state"`
}

func (Client) TableName() string {
	return "clients"
}
