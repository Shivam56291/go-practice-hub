package migration

import (
	articleModels "Blog/internal/modules/article/models"
	userModels "Blog/internal/modules/user/models"
	"Blog/pkg/database"
	"fmt"
	"log"
)

func Migrate() {
	db := database.Connection()

	err := db.AutoMigrate(&userModels.User{})
	if err != nil {
		log.Fatal("failed to migrate user")
	}
	err = db.AutoMigrate(&articleModels.Article{})
	if err != nil {
		log.Fatal("failed to migrate article")
	}

	fmt.Println("migrated successfully")
}
