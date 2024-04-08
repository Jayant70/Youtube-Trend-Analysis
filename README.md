## Backend Assignment | FamPay

### Project Goal

To make an API to fetch latest videos sorted in reverse chronological order of the publishing date-time from YouTube for a given search query in paginated response

### Basic Functionalities

- Cron Job to constantly fetch data in the background every minute
- GET API, `/videos` for fetching videos in a pagination format from DB.
- GET API, /videos/search for searching videos based on title and description
- Search API which also supports fuzzy matching for situations like `How to make a tea?` matched with `tea how`

### Development

1. Clone the project

`https://github.com/Jayant70/youtube-search-dashboard/tree/develop`

```
# For default values, refer `.env` file

# Server Properties
HTTP_PORT =

# MONGODB
MONGO_DB_URL = 
DB_NAME =
COLLECTION_NAME =

# YOUTUBE API
API_KEY =
SEARCH_QUERY =
```

You will need a YOUTUBE DATA API key in order to run this app. Follow the instructions on [this page](https://developers.google.com/youtube/v3/getting-started) to get one.

2. Install dependencies

```
go mod tidy
```
3. Run in development mode
```
go run main.go
```
### Running with Docker Compose
When using Docker Compose,

1. Create a `.env` file using the instructions mentioned above
2. Set the `MONGO_DB_URL` environment variable in your `.env` file to
```
MONGO_DB_URL = mongodb://mongo:27017
```
3. Run:

```
docker-compose up -d
```