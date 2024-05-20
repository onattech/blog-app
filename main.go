package main

import (
	"log"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Initialize the database
	db, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Auto-migrate models
	err = db.AutoMigrate(&Article{}, &Comment{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	// Setup router
	r := http.NewServeMux()

	// Article routes
	r.HandleFunc("GET /articles", ListArticles(db))
	r.HandleFunc("GET /articles/{id}", GetArticle(db))
	r.HandleFunc("POST /articles", CreateArticle(db))
	r.HandleFunc("PATCH /articles/{id}", UpdateArticle(db))
	r.HandleFunc("DELETE /articles/{id}", DeleteArticle(db))

	// Comment routes
	r.HandleFunc("GET /articles/{id}/comments", ListComments(db))
	r.HandleFunc("POST /articles/{id}/comments", CreateComment(db))
	r.HandleFunc("DELETE /comments/{id}", DeleteComment(db))

	// Setup server
	server := http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	// Start server
	log.Println("Server running on port 8080")
	log.Fatal(server.ListenAndServe())
}
