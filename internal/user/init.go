package user

import (
	"github.com/dimapog/jwt-microservice/utils"
)

func init() {
	if err := Migrate(); err != nil {
		panic(err)
	}
}

func Migrate() error {
	return utils.DB.AutoMigrate(&User{})
}
