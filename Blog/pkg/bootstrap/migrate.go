package bootstrap

import (
	"Blog/internal/database/migration"
	"Blog/pkg/config"
	"Blog/pkg/database"
)

func Migrate() {
	config.Set()

	database.Connect()
	migration.Migrate()
}
