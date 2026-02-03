# Freesound API Wrapper

A simple Go REST API that wraps the Freesound API to fetch sound/song information.

## Prerequisites

- Go 1.24 or later
- A Freesound API key

## Getting an API Key

1. Create an account at https://freesound.org
2. Go to https://freesound.org/apiv2/apply/
3. Fill out the application form
4. Copy your API key

## Setup

1. Clone the repository

2. Create a `.env` file in the project root:
```
FREESOUND_API_KEY=your_api_key_here
```

3. Install dependencies:
```bash
go mod tidy
```

## Running

```bash
go run main.go
```

The server starts on port 8080 by default. Set the `PORT` environment variable to change it.

## Endpoints

- `GET /` - Info page
- `GET /song/{id}/` - Fetch a single sound by ID
- `GET /songs?ids=123,456,789` - Fetch multiple sounds by comma-separated IDs
