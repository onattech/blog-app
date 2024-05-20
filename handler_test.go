package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Setup an in-memory database for testing
func setupTestDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&Article{}, &Comment{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

// Test ListArticles handler
func TestListArticles(t *testing.T) {
	// Setup the database
	db, err := setupTestDB()
	assert.NoError(t, err)

	// Pre-test database seed
	db.Create(&Article{Title: "Test Article", Author: "Test Author", Content: "Test Content"})

	// Setup the HTTP request
	req, err := http.NewRequest("GET", "/articles", nil)
	assert.NoError(t, err)

	// Setup the response recorder to record the HTTP response
	rr := httptest.NewRecorder()

	// Setup the router and handler
	r := http.NewServeMux()
	r.HandleFunc("GET /articles", ListArticles(db))

	// Serve the HTTP request
	r.ServeHTTP(rr, req)

	// Assert
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "Test Article")
}

// Test GetArticle handler
func TestGetArticle(t *testing.T) {
	// Setup the database
	db, err := setupTestDB()
	assert.NoError(t, err)

	// Pre-test database seed
	article := Article{Title: "Test Article", Author: "Test Author", Content: "Test Content"}
	db.Create(&article)

	// Setup the HTTP request
	req, err := http.NewRequest("GET", "/articles/"+strconv.Itoa(int(article.ID)), nil)
	assert.NoError(t, err)

	// Setup the response recorder to record the HTTP response
	rr := httptest.NewRecorder()

	// Setup the router and handler
	r := http.NewServeMux()
	r.HandleFunc("GET /articles/{id}", GetArticle(db))

	// Serve the HTTP request
	r.ServeHTTP(rr, req)

	// Assert
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "Test Article")
}

// Test CreateArticle handler
func TestCreateArticle(t *testing.T) {
	// Setup the database
	db, err := setupTestDB()
	assert.NoError(t, err)

	// HTTP request payload
	newArticle := Article{Title: "New Article", Author: "New Author", Content: "New Content"}
	body, err := json.Marshal(newArticle)
	assert.NoError(t, err)

	// Setup the HTTP request
	req, err := http.NewRequest("POST", "/articles", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	assert.NoError(t, err)

	// Setup the response recorder to record the HTTP response
	rr := httptest.NewRecorder()

	// Setup the router and handler
	r := http.NewServeMux()
	r.HandleFunc("POST /articles", CreateArticle(db))

	// Serve the HTTP request
	r.ServeHTTP(rr, req)

	// Assert
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "New Article")
}

// Test UpdateArticle handler
func TestUpdateArticle(t *testing.T) {
	// Setup the database
	db, err := setupTestDB()
	assert.NoError(t, err)

	// Pre-test database seed
	article := Article{Title: "Test Article", Author: "Test Author", Content: "Test Content"}
	db.Create(&article)

	// HTTP request payload
	updatedFields := map[string]interface{}{
		"title": "Updated Article",
	}
	body, err := json.Marshal(updatedFields)
	assert.NoError(t, err)

	// Setup the HTTP request
	req, err := http.NewRequest("PATCH", "/articles/"+strconv.Itoa(int(article.ID)), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	assert.NoError(t, err)

	// Setup the response recorder to record the HTTP response
	rr := httptest.NewRecorder()

	// Setup the router and handler
	r := http.NewServeMux()
	r.HandleFunc("PATCH /articles/{id}", UpdateArticle(db))

	// Serve the HTTP request
	r.ServeHTTP(rr, req)

	// Assert
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "Updated Article")
}

// Test DeleteArticle handler
func TestDeleteArticle(t *testing.T) {
	// Setup the database
	db, err := setupTestDB()
	assert.NoError(t, err)

	// Pre-test database seed
	article := Article{Title: "Test Article", Author: "Test Author", Content: "Test Content"}
	db.Create(&article)

	// Setup the HTTP request
	req, err := http.NewRequest("DELETE", "/articles/"+strconv.Itoa(int(article.ID)), nil)
	assert.NoError(t, err)

	// Setup the response recorder to record the HTTP response
	rr := httptest.NewRecorder()

	// Setup the router and handler
	r := http.NewServeMux()
	r.HandleFunc("DELETE /articles/{id}", DeleteArticle(db))

	// Serve the HTTP request
	r.ServeHTTP(rr, req)

	// Assert
	assert.Equal(t, http.StatusNoContent, rr.Code)
}

// Test ListComments handler
func TestListComments(t *testing.T) {
	// Setup the database
	db, err := setupTestDB()
	assert.NoError(t, err)

	// Pre-test database seed
	article := Article{Title: "Test Article", Author: "Test Author", Content: "Test Content"}
	db.Create(&article)
	comment := Comment{ArticleID: article.ID, Name: "Test Commenter", Comment: "Test Comment", CreatedAt: time.Now()}
	db.Create(&comment)

	// Setup the HTTP request
	req, err := http.NewRequest("GET", "/articles/"+strconv.Itoa(int(article.ID))+"/comments", nil)
	assert.NoError(t, err)

	// Setup the response recorder to record the HTTP response
	rr := httptest.NewRecorder()

	// Setup the router and handler
	r := http.NewServeMux()
	r.HandleFunc("GET /articles/{id}/comments", ListComments(db))

	// Serve the HTTP request
	r.ServeHTTP(rr, req)

	// Assert
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "Test Comment")
}

// Test CreateComment handler
func TestCreateComment(t *testing.T) {
	// Setup the database
	db, err := setupTestDB()
	assert.NoError(t, err)

	// Pre-test database seed
	article := Article{Title: "Test Article", Author: "Test Author", Content: "Test Content"}
	db.Create(&article)

	// HTTP request payload
	newComment := Comment{Name: "New Commenter", Comment: "New Comment"}
	body, err := json.Marshal(newComment)
	assert.NoError(t, err)

	// Setup the HTTP request
	req, err := http.NewRequest("POST", "/articles/"+strconv.Itoa(int(article.ID))+"/comments", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	assert.NoError(t, err)

	// Setup the response recorder to record the HTTP response
	rr := httptest.NewRecorder()

	// Setup the router and handler
	r := http.NewServeMux()
	r.HandleFunc("POST /articles/{id}/comments", CreateComment(db))

	// Serve the HTTP request
	r.ServeHTTP(rr, req)

	// Assert
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "New Comment")
}

// Test DeleteComment handler
func TestDeleteComment(t *testing.T) {
	// Setup the database
	db, err := setupTestDB()
	assert.NoError(t, err)

	// Pre-test database seed
	article := Article{Title: "Test Article", Author: "Test Author", Content: "Test Content"}
	db.Create(&article)
	comment := Comment{ArticleID: article.ID, Name: "Test Commenter", Comment: "Test Comment", CreatedAt: time.Now()}
	db.Create(&comment)

	// Setup the HTTP request
	req, err := http.NewRequest("DELETE", "/comments/"+strconv.Itoa(int(comment.ID)), nil)
	assert.NoError(t, err)

	// Setup the response recorder to record the HTTP response
	rr := httptest.NewRecorder()

	// Setup the router and handler
	r := http.NewServeMux()
	r.HandleFunc("DELETE /comments/{id}", DeleteComment(db))

	// Serve the HTTP request
	r.ServeHTTP(rr, req)

	// Assert
	assert.Equal(t, http.StatusNoContent, rr.Code)
}
