package user

import (
	"database/sql"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name             string          `gorm:"column:name"`
	Email            string          `gorm:"column:email;unique"`
	Password         string          `gorm:"column:password"`
	Age              sql.NullInt64   `gorm:"column:age"`
	Gender           sql.NullString  `gorm:"column:gender"`
	Weight           sql.NullFloat64 `gorm:"column:weight"`
	Height           sql.NullFloat64 `gorm:"column:height"`
	RestingHeartRate sql.NullInt64   `gorm:"column:resting_heart_rate"`
	Units            sql.NullString  `gorm:"column:units"`
}

func (User) TableName() string {
	return "users"
}
