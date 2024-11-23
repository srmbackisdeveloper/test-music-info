# **Effective Mobile GO | Music Library API**

## **Features**

- **CRUD Operations**: Add, update, delete, and fetch songs.
- **Lyrics Pagination**: Retrieve song lyrics with pagination by verses.
- **Search and Filter**: Filter songs by group or title.
- **Caching**: Redis caching for frequently accessed data.
- **Swagger Documentation**: Comprehensive API documentation.

---

## **Technologies Used**

- **Go** (Golang)
- **Gin Framework** (Web framework)
- **PostgreSQL** (Database)
- **Redis** (Cache)
- **Swagger** (API Documentation)
- **GORM** (ORM for Go)

---

## Setup .env (example)

```
PORT=8080

POSTGRES_DSN=postgres://user:password@postgres:5432/db-name?sslmode=disable
REDIS_ADDRESS=redis:6379
REDIS_PASSWORD=
REDIS_DB=0

LOG_LEVEL=info

POSTGRES_USER=user
POSTGRES_PASSWORD=password
POSTGRES_DB=db-name
```
