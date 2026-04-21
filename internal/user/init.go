package user

import (
	"github.com/dimapog/jwt-microservice/utils"
)

func Migrate() {
	utils.DB.AutoMigrate(&User{})
}
