package user

import (
	"github.com/dimapog/d4-microservice/utils"
)

func Migrate() error {
	return utils.DB.AutoMigrate(&User{})
}
