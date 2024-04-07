package constants

import (
	"time"
)

const (
	ApiKey              = "AIzaSyAVZk7pKHah7CL3QxMHG8ZNWx7S4vCwt7s"
	SearchQuery         = "cricket"
	DefaultPollInterval = 60 * time.Second
	DbName              = "youtube"
	CollectionName      = "videos"

	ContextPath = "/api/v1"

	POST = "POST"
	GET  = "GET"

	// GetVideosDataURI get videos
	GetVideosDataURI = ContextPath + "/videos"
)
