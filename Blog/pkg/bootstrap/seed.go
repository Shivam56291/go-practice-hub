package bootstrap

import (
	"Blog/internal/database/seeder"
	"Blog/pkg/config"
	"Blog/pkg/database"
)

func Seed() {
	config.Set()

	database.Connect()

	seeder.Seed()
}
