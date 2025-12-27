package seeder

import (
	articleModel "Blog/internal/modules/article/models"
	userModel "Blog/internal/modules/user/models"
	"Blog/pkg/database"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func Seed() {
	db := database.Connection()

	log.Println("Starting database seeding...")

	// -----------------------------
	// Clean existing data
	// -----------------------------
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")
	db.Exec("TRUNCATE TABLE articles")
	db.Exec("TRUNCATE TABLE users")
	db.Exec("SET FOREIGN_KEY_CHECKS = 1")

	log.Println("Existing users and articles removed")

	// -----------------------------
	// Create users
	// -----------------------------
	password, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Password hashing failed:", err)
	}

	users := []userModel.User{
		{
			Name:     "Shivam Kumar",
			Email:    "shivam@gmail.com",
			Password: string(password),
		},
		{
			Name:     "John Doe",
			Email:    "john.doe@gmail.com",
			Password: string(password),
		},
	}

	for i := range users {
		db.Create(&users[i])
		log.Printf("User created: %s", users[i].Name)
	}

	// -----------------------------
	// Create articles (long content)
	// -----------------------------
	articles := []articleModel.Article{
		{
			Title: "Getting Started with Go Programming",
			Content: `Go is a statically typed, compiled programming language designed at Google.
It focuses on simplicity, performance, and readability. One of Go's strongest
features is its excellent support for concurrency using goroutines and channels.
Developers choose Go for building fast web servers, APIs, and distributed systems.
If you are coming from languages like Java or Python, Go offers a refreshing balance
between performance and developer productivity.`,
			UserID: users[0].ID,
		},
		{
			Title: "Understanding MVC Architecture in Web Applications",
			Content: `The Model-View-Controller (MVC) pattern is one of the most commonly used
architectural patterns in web development. The model handles business logic,
the view manages UI rendering, and the controller acts as a bridge between them.
By separating concerns, MVC makes applications easier to maintain, test, and scale.
Most modern frameworks follow this pattern in one form or another.`,
			UserID: users[0].ID,
		},
		{
			Title: "Building REST APIs with Gin Framework",
			Content: `Gin is a high-performance HTTP web framework written in Go.
It provides a minimalistic API with excellent routing and middleware support.
Gin is widely used for building RESTful APIs due to its speed and simplicity.
With features like JSON validation, middleware chaining, and error handling,
Gin makes backend development both efficient and enjoyable.`,
			UserID: users[0].ID,
		},
		{
			Title: "Authentication and Authorization Explained",
			Content: `Authentication and authorization are core security concepts in web applications.
Authentication verifies the identity of a user, while authorization determines
what resources the user can access. Implementing these correctly prevents
unauthorized access and data breaches. Modern applications often use sessions,
JWT tokens, or OAuth for handling authentication securely.`,
			UserID: users[1].ID,
		},
		{
			Title: "Handling Forms and Validation in Go",
			Content: `Form handling is an essential part of any web application.
Proper validation ensures that user input is safe and meaningful.
In Go, developers typically use request structs combined with validation logic.
Displaying error messages and preserving old input significantly improves
the user experience and reduces frustration.`,
			UserID: users[1].ID,
		},
		{
			Title: "Repository and Service Layer Pattern",
			Content: `The repository and service layer pattern helps in structuring
large applications by separating data access logic from business rules.
Repositories handle database operations, while services orchestrate workflows.
This approach improves testability, scalability, and maintainability of code.`,
			UserID: users[1].ID,
		},
		{
			Title: "Middleware in Web Applications",
			Content: `Middleware acts as a pipeline between the request and response.
It is commonly used for authentication, logging, request modification,
and error handling. In Gin, middleware can be applied globally or per route.
Well-designed middleware simplifies cross-cutting concerns in applications.`,
			UserID: users[0].ID,
		},
		{
			Title: "Best Practices for Scalable Backend Development",
			Content: `Scalable backend systems require thoughtful design and clean architecture.
Using proper abstractions, following SOLID principles, and writing modular code
helps systems grow without becoming unmanageable. Monitoring, logging, and
proper error handling are equally important for long-term success.`,
			UserID: users[0].ID,
		},
	}

	for _, article := range articles {
		db.Create(&article)
		log.Printf("Article created: %s", article.Title)
	}

	log.Println("Database seeding completed successfully")
}
