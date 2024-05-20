# Blog API

This is a simple blog-like application built with Golang, using the `net/http` package for the API, `GORM` as the ORM, and `SQLite` for the database. The application allows for managing articles and comments, and tracks how often people start writing a comment but then delete it.

## Features

-   **Articles Management**

    -   Create, read, update, and delete articles.
    -   List all articles.

-   **Comments Management**

    -   Create, read, and delete comments for specific articles.
    -   List all comments for a specific article, including soft-deleted comments.

-   **Soft Deletion**
    -   Soft delete support for comments.

### Loom walk through

1. https://www.loom.com/share/79c12559812e4fae96cbe2ec4954763c
2. https://www.loom.com/share/c6a25e1cec4d407ba7ab894167e7aa8f
