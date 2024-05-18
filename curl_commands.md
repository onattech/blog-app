### List All Articles

curl -X GET http://localhost:8080/articles

### Get a Single Article by ID

curl -X GET http://localhost:8080/articles/{id}

### Create a New Article

curl -X POST http://localhost:8080/articles -H "Content-Type: application/json" -d '{
"Title": "Sample Title",
"Author": "Author Name",
"Content": "This is the content of the article."
}'

### Update an Article by ID

curl -X PUT http://localhost:8080/articles/{id} -H "Content-Type: application/json" -d '{
"Title": "Updated Title",
"Author": "Updated Author",
"Content": "Updated content of the article."
}'

### Delete an Article by ID

curl -X DELETE http://localhost:8080/articles/{id}

### List All Comments for a Specific Article

curl -X GET http://localhost:8080/articles/{id}/comments

### Create a New Comment for a Specific Article

curl -X POST http://localhost:8080/articles/{id}/comments -H "Content-Type: application/json" -d '{
"Name": "Commenter Name",
"Comment": "This is a comment."
}'

### Delete a Comment by ID

curl -X DELETE http://localhost:8080/comments/{id}
