package csv

import (
	"github.com/dimapog/jwt-microservice/utils"
)

func Migrate() error {
	return utils.DB.AutoMigrate(&Client{})
}
