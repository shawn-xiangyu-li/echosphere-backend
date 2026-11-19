# EchoSphere Backend

> 🚧 **Work in progress.** This is the backend of EchoSphere, a social media app where users share image and video posts.
> Frontend: [echosphere-frontend](https://github.com/shawn-xiangyu-li/echosphere-frontend)

A REST API written in Go. It handles sign-up and login with JWT, stores posts in Elasticsearch for search, and uploads media files to Google Cloud Storage.

## Tech stack

- **Language:** Go
- **Routing:** gorilla/mux, with CORS from gorilla/handlers
- **Auth:** JWT (HS256), checked by auth0/go-jwt-middleware
- **Search and storage:** Elasticsearch (self-hosted on Google Compute Engine)
- **Media:** Google Cloud Storage
- **Hosting:** Google App Engine (flexible environment)

## Architecture

```
Next.js frontend ──HTTP/JSON──> Go REST API (App Engine)
                                   ├── Elasticsearch (GCE)  users + posts, full-text search
                                   └── Google Cloud Storage  images and videos
```

```
├── main.go          # starts the server on :8080
├── handler/         # HTTP handlers and router (JWT middleware, CORS)
├── service/         # business logic
├── backend/         # Elasticsearch and GCS clients
├── model/           # Post and User structs
└── constants/       # index names and settings read from environment variables
```

## API

| Method | Endpoint | Auth | Description |
|---|---|---|---|
| POST | `/users` | No | Sign up (`{"username", "password"}`) |
| POST | `/auth/login` | No | Log in, returns a JWT (valid for 24 hours) |
| POST | `/posts` | JWT | Create a post: multipart form with `message` and `media_file` |
| GET | `/posts/search?user=NAME` | JWT | Posts by one user |
| GET | `/posts/search?keywords=WORDS` | JWT | Full-text search on post messages (all posts if empty) |

## Configuration

All connection settings and secrets come from environment variables. The server stops at startup if any are missing.

| Variable | Description |
|---|---|
| `ES_URL` | Elasticsearch URL, e.g. `http://host:9200` |
| `ES_USERNAME` / `ES_PASSWORD` | Elasticsearch basic auth |
| `GCS_BUCKET` | Google Cloud Storage bucket for media |
| `JWT_SECRET` | Key used to sign JWTs |

GCS uses Google Application Default Credentials (on App Engine this works automatically; locally, run `gcloud auth application-default login`).

## Run locally

```bash
cp .env.example .env      # then fill in real values
set -a; source .env; set +a
go run .
```

## Deploy to App Engine

```bash
cp app.yaml.example app.yaml   # fill in env_variables (app.yaml is git-ignored)
gcloud app deploy
```

## Roadmap

- [ ] AI-assisted post creation (image generation in the post editor)
- [ ] Hash passwords with bcrypt instead of storing them as plain text
- [ ] Delete posts (`DELETE /posts/{id}`)
- [ ] Restrict CORS to the frontend's domain
- [ ] Tests for handlers and services
