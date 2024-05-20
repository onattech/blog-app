package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

// Handler to list all articles
func ListArticles(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Retrieve all articles from the database, including their comments
		var articles []Article
		if err := db.Preload("Comments").Find(&articles).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Respond
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(articles)
	}
}

// Handler to get an article by ID
func GetArticle(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract the ID from the URL path
		articleId := r.PathValue("id")

		// Retrieve the article from the database by ID
		var article Article
		if err := db.First(&article, articleId).Error; err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		// Respond
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(article)
	}
}

// Handler to create a new article
func CreateArticle(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Decode the request body into a new struct
		var article Article
		if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Insert the new article into the database
		if err := db.Create(&article).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Respond
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(article)
	}
}

// Handler to update an article by ID
func UpdateArticle(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Decode the request body into a map for partial updates
		var updates map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Extract the ID from the URL path
		articleId := r.PathValue("id")

		// Update the article in the database
		if err := db.Model(&Article{}).Where("id = ?", articleId).Updates(updates).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Retrieve the updated article to send in the response
		var updatedArticle Article
		if err := db.First(&updatedArticle, articleId).Error; err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		// Respond
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(updatedArticle)
	}
}

// Handler to delete an article by ID
func DeleteArticle(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract the ID from the URL path
		articleId := r.PathValue("id")

		// Delete the article in the database
		if err := db.Delete(&Article{}, articleId).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Respond
		w.WriteHeader(http.StatusNoContent)
	}
}

// Handler to list all comments for a specific article
func ListComments(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract the ID from the URL path
		articleID := r.PathValue("id")

		// Retrieve all comments for the specified article from the database
		var comments []Comment
		if err := db.Where("article_id = ?", articleID).Unscoped().Find(&comments).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Respond
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(comments)
	}
}

// Handler to create a new comment for a specific article
func CreateComment(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Decode the request body into a new struct
		var comment Comment
		if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Extract the article ID from the URL path and set it in the comment
		articleID := r.PathValue("id")
		id, err := strconv.Atoi(articleID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		comment.ArticleID = uint(id)

		// Insert the new comment into the database
		if err := db.Create(&comment).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Respond
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(comment)
	}
}

// Handler to delete a comment by ID (soft delete)
func DeleteComment(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract the ID from the URL path
		commentId := r.PathValue("id")

		// Soft delete the comment in the database
		if err := db.Delete(&Comment{}, commentId).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Respond
		w.WriteHeader(http.StatusNoContent)
	}
}
