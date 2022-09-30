package database

import (
	"github.com/Cambara/pokedex-service/dto/models"
)

func Migrate() {
	db := GetInstance().Db
	db.AutoMigrate(models.Pokemon{})
}
