# RSS Aggregator

This project is an RSS feed aggregation and management system built in Go. It provides features for handling RSS feeds, user authentication, feed following, and more. Below is an overview of the project and its functionality.

## Features

- **User Management**: Handles user authentication and management.
- **RSS Feed Management**: Enables scraping and managing RSS feeds.
- **Feed Following**: Allows users to follow specific feeds and retrieve their updates.
- **Error Handling**: Includes robust error-handling mechanisms.
- **SQL Integration**: Database integration for persistent storage.

## Project Structure

```
rssagg/
├── auth                  # Authentication logic and middleware
├── internal              # Internal utilities and helpers
├── sql                   # SQL migrations and queries
├── handler_feed.go       # Handlers for RSS feed operations
├── handler_feed_follows.go # Handlers for feed-following functionality
├── handler_user.go       # Handlers for user-related operations
├── main.go               # Entry point for the application
├── middleware_auth.go    # Middleware for authentication
├── models.go             # Data models used in the application
├── scraper.go            # RSS scraping logic
└── ...
```

## Prerequisites

- [Go](https://go.dev/dl/) (v1.19 or later)
- A database (e.g., PostgreSQL, MySQL) for storing data
- `sqlc` for generating database models and queries

## Getting Started

### Clone the Repository

```bash
git clone <repository-url>
cd rssagg
```

### Set Up Environment Variables

Copy the `.env` file and configure your environment variables:

```bash
cp .env.example .env
```

### Install Dependencies

Use `go mod` to install the necessary dependencies:

```bash
go mod tidy
```

### Set Up the Database

Run the SQL migration scripts in the `sql` directory to set up the database schema:

```bash
# Example for PostgreSQL
psql -U <username> -d <database> -f sql/migrations.sql
```

### Run the Application

Start the application:

```bash
go run main.go
```

## API Endpoints

### Authentication

- `POST /login`: Authenticate a user
- `POST /register`: Register a new user

### Feeds

- `GET /feeds`: Retrieve all feeds
- `POST /feeds`: Add a new feed

### Feed Follows

- `GET /follows`: Retrieve followed feeds
- `POST /follows`: Follow a feed
- `DELETE /follows/:id`: Unfollow a feed




