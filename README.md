# Forum Backend API

A simple forum backend application built with Golang.

## Features
- User authentication (JWT-based)
- Create, read, update, and delete forum posts
- Comment on posts
- Basic user roles (admin, user)
- RESTful API with JSON responses

## Tech Stack
- Language: Go
- Framework: Gin
- Database: PostgreSQL/MySQL
- Authentication: JWT
- Migrations: Golang Migrate

## Installation

### Prerequisites
- Go 1.20+ installed
- MySQL database setup
- Environment variables configured

### Clone the Repository
```sh
git clone https://github.com/rifkifajarramadhani/golang-forum-app.git
cd golang-forum-app
```

### Install Dependencies
```sh
go mod tidy
```

### Configure Environment Variables
Create a `.env` file and update the following variables:
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=youruser
DB_PASSWORD=yourpassword
DB_NAME=forum_db
JWT_SECRET=your_secret_key
```

### Run Database Migrations
```sh
migrate -database "postgres://youruser:yourpassword@localhost:5432/forum_db?sslmode=disable" -path ./migrations up
```

### Run the Application
```sh
go run main.go
```

## API Endpoints

### Authentication
- `POST /register` - Register a new user
- `POST /login` - Authenticate user and get token

### Forum Posts
- `GET /posts` - Get all posts
- `GET /posts/{id}` - Get a single post
- `POST /posts` - Create a new post (auth required)
- `PUT /posts/{id}` - Update a post (auth required)
- `DELETE /posts/{id}` - Delete a post (auth required)

### Comments
- `POST /posts/{id}/comments` - Add a comment
- `GET /posts/{id}/comments` - Get all comments for a post

## License
This project is licensed under the MIT License.

