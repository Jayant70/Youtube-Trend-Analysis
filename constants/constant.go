package constants

import (
	"time"
)

const (
	DefaultApiKey       = "AIzaSyAVZk7pKHah7CL3QxMHG8ZNWx7S4vCwt7s"
	SearchQuery         = "cricket"
	DefaultPollInterval = 60 * time.Second
	DefaultDbName       = "youtube"
	CollectionName      = "videos"
	DefaultHttpPort     = "8080"
	DefaultMongoURI     = "mongodb://localhost:27017"

	ContextPath = "/api/v1"

	POST = "POST"
	GET  = "GET"

	// GetVideosDataURI get videos
	GetVideosDataURI        = ContextPath + "/videos"
	GetVideosDataByQueryURI = ContextPath + "/videos/search"
)
