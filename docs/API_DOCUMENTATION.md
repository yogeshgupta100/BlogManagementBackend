# Blog Management API Documentation

## Overview
A comprehensive blog management system with CRUD operations built with Go Fiber and PostgreSQL.

## Base URL
```
http://localhost:8080
```

## API Endpoints

### 1. Create Blog Post
**POST** `/api/blog-post`

Creates a new blog post.

#### Request Body
```json
{
  "title": "My First Blog Post",
  "description": "This is a brief description of my blog post",
  "body": "This is the main content of my blog post..."
}
```

#### Response (201 Created)
```json
{
  "message": "Blog post created successfully",
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "title": "My First Blog Post",
    "description": "This is a brief description of my blog post",
    "body": "This is the main content of my blog post...",
    "created_at": "2023-01-01T00:00:00Z",
    "updated_at": "2023-01-01T00:00:00Z"
  }
}
```

#### Error Response (400 Bad Request)
```json
{
  "error": "Failed to create blog post",
  "message": "title is required"
}
```

---

### 2. Get All Blog Posts
**GET** `/api/blog-post`

Retrieves all blog posts from the database.

#### Response (200 OK)
```json
{
  "message": "Blog posts retrieved successfully",
  "data": [
    {
      "id": "550e8400-e29b-41d4-a716-446655440000",
      "title": "My First Blog Post",
      "description": "This is a brief description of my blog post",
      "body": "This is the main content of my blog post...",
      "created_at": "2023-01-01T00:00:00Z",
      "updated_at": "2023-01-01T00:00:00Z"
    }
  ],
  "count": 1
}
```

#### Error Response (500 Internal Server Error)
```json
{
  "error": "Failed to retrieve blog posts",
  "message": "database error"
}
```

---

### 3. Get Blog Post by ID
**GET** `/api/blog-post/{id}`

Retrieves a specific blog post by its unique identifier.

#### Path Parameters
- `id` (string, required): The unique identifier of the blog post

#### Response (200 OK)
```json
{
  "message": "Blog post retrieved successfully",
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "title": "My First Blog Post",
    "description": "This is a brief description of my blog post",
    "body": "This is the main content of my blog post...",
    "created_at": "2023-01-01T00:00:00Z",
    "updated_at": "2023-01-01T00:00:00Z"
  }
}
```

#### Error Response (404 Not Found)
```json
{
  "error": "Blog post not found",
  "message": "The requested blog post does not exist"
}
```

#### Error Response (400 Bad Request)
```json
{
  "error": "Blog ID is required",
  "message": "Please provide a valid blog ID"
}
```

---

### 4. Update Blog Post
**PATCH** `/api/blog-post/{id}`

Updates an existing blog post with partial data.

#### Path Parameters
- `id` (string, required): The unique identifier of the blog post

#### Request Body
```json
{
  "title": "Updated Blog Post Title",
  "description": "Updated description",
  "body": "Updated blog post content..."
}
```

**Note:** All fields are optional. Only include the fields you want to update.

#### Response (200 OK)
```json
{
  "message": "Blog post updated successfully",
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "title": "Updated Blog Post Title",
    "description": "Updated description",
    "body": "Updated blog post content...",
    "created_at": "2023-01-01T00:00:00Z",
    "updated_at": "2023-01-02T00:00:00Z"
  }
}
```

#### Error Response (404 Not Found)
```json
{
  "error": "Blog post not found",
  "message": "The requested blog post does not exist"
}
```

#### Error Response (400 Bad Request)
```json
{
  "error": "Failed to update blog post",
  "message": "title cannot be empty"
}
```

---

### 5. Delete Blog Post
**DELETE** `/api/blog-post/{id}`

Deletes a blog post by its unique identifier.

#### Path Parameters
- `id` (string, required): The unique identifier of the blog post

#### Response (200 OK)
```json
{
  "message": "Blog post deleted successfully"
}
```

#### Error Response (404 Not Found)
```json
{
  "error": "Blog post not found",
  "message": "The requested blog post does not exist"
}
```

#### Error Response (400 Bad Request)
```json
{
  "error": "Blog ID is required",
  "message": "Please provide a valid blog ID"
}
```

---

### 6. Health Check
**GET** `/health`

Checks if the API is running.

#### Response (200 OK)
```json
{
  "status": "OK",
  "message": "Blog Management API is running"
}
```

---

## Data Models

### BlogCreateRequest
```json
{
  "title": "string (required, max 255 characters)",
  "description": "string (optional, max 1000 characters)",
  "body": "string (required, min 1 character)"
}
```

### BlogUpdateRequest
```json
{
  "title": "string (optional, max 255 characters)",
  "description": "string (optional, max 1000 characters)",
  "body": "string (optional, min 1 character)"
}
```

### BlogResponse
```json
{
  "id": "string (UUID)",
  "title": "string",
  "description": "string",
  "body": "string",
  "created_at": "datetime (ISO 8601)",
  "updated_at": "datetime (ISO 8601)"
}
```

---

## Error Handling

All endpoints return consistent error responses with the following structure:

```json
{
  "error": "Error title",
  "message": "Detailed error message"
}
```

### Common HTTP Status Codes
- `200` - Success
- `201` - Created
- `400` - Bad Request (validation errors)
- `404` - Not Found
- `500` - Internal Server Error

---

## Testing the API

### Using curl

#### Create a blog post
```bash
curl -X POST http://localhost:8080/api/blog-post \
  -H "Content-Type: application/json" \
  -d '{
    "title": "My First Blog Post",
    "description": "This is a brief description",
    "body": "This is the main content of my blog post..."
  }'
```

#### Get all blog posts
```bash
curl -X GET http://localhost:8080/api/blog-post
```

#### Get a specific blog post
```bash
curl -X GET http://localhost:8080/api/blog-post/550e8400-e29b-41d4-a716-446655440000
```

#### Update a blog post
```bash
curl -X PATCH http://localhost:8080/api/blog-post/550e8400-e29b-41d4-a716-446655440000 \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Updated Title"
  }'
```

#### Delete a blog post
```bash
curl -X DELETE http://localhost:8080/api/blog-post/550e8400-e29b-41d4-a716-446655440000
```

---

## Environment Variables

Configure the following environment variables in `config.env`:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=blog_management
SERVER_PORT=8080
```

---

## Running the Application

1. **Install dependencies:**
   ```bash
   go mod tidy
   ```

2. **Set up PostgreSQL database:**
   - Create a database named `blog_management`
   - Update the database credentials in `config.env`

3. **Run the application:**
   ```bash
   go run main.go
   ```

4. **Run tests:**
   ```bash
   go test ./... -cover
   ```

The API will be available at `http://localhost:8080` 