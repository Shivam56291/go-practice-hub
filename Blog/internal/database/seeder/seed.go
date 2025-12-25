package seeder

import (
	articleModel "Blog/internal/modules/article/models"
	userModel "Blog/internal/modules/user/models"
	"Blog/pkg/database"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func Seed() {
	db := database.Connection()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Hashed password error", err)
		return
	}

	user := userModel.User{
		Name:     "Random name",
		Email:    "random@email.com",
		Password: string(hashedPassword),
	}
	db.Create(&user)

	log.Printf("User seeded successfully with email address %s", user.Email)

	for i := 0; i < 10; i++ {
		article := articleModel.Article{
			Title:   fmt.Sprintf("Random title %d", i),
			Content: fmt.Sprintf("Random content %d", i),
			UserID:  user.ID,
		}
		db.Create(&article)
		log.Printf("Article seeded successfully with title %s", article.Title)
	}

	log.Println("Seeding completed successfully")
}
