package Jobs

import (
	"context"
	"fmt"
	"google.golang.org/api/youtube/v3"
	"log"
	"os"
	"time"
	"youtube/constants"
	"youtube/helper"
	"youtube/types"
)

func PollVideos(youtubeService *youtube.Service) {
	// Get MongoDB Database and Collection
	dbName := os.Getenv("DB_NAME")
	if len(dbName) == 0 {
		dbName = constants.DefaultDbName
	}
	collectionName := os.Getenv("COLLECTION_NAME")
	if len(collectionName) == 0 {
		collectionName = constants.CollectionName
	}
	pollIntervalStr := os.Getenv("POLL_INTERVAL")
	if len(pollIntervalStr) == 0 {
		pollIntervalStr = constants.DefaultPollInterval.String()
	}
	pollInterval, err := time.ParseDuration(pollIntervalStr)
	if err != nil {
		// Use the default poll interval if parsing fails
		pollInterval = constants.DefaultPollInterval
	}

	database := helper.Client.Database(dbName)
	collection := database.Collection(collectionName)

	//Started for loop for polling videos every  PollInterval time
	for {
		videos, err := pollLatestVideos(youtubeService)
		if err != nil {
			log.Printf("Error fetching videos: %v", err)
			continue
		}

		// Store videos in MongoDB
		for _, video := range videos {
			_, err := collection.InsertOne(context.Background(), video)
			if err != nil {
				log.Printf("Error inserting video into MongoDB: %v", err)
			}
		}

		time.Sleep(pollInterval)
	}
}

func pollLatestVideos(service *youtube.Service) ([]*types.Video, error) {
	oneMinuteAgo := time.Now().Add(-1 * time.Minute)
	// Make a search API call to get the latest videos
	searchQuery := os.Getenv("SEARCH_QUERY")
	if len(searchQuery) == 0 {
		searchQuery = constants.SearchQuery
	}
	call := service.Search.List([]string{"snippet"}).
		Q(searchQuery).
		MaxResults(10).
		Order("date").
		Type("video").
		PublishedAfter(oneMinuteAgo.Format(time.RFC3339))

	// Execute the API call
	response, err := call.Do()
	if err != nil {
		return nil, fmt.Errorf("error making search API call: %v", err)
	}
	// Parse the response and create a slice of Video objects
	var videos []*types.Video
	for _, item := range response.Items {
		publishedAt, err := time.Parse(time.RFC3339, item.Snippet.PublishedAt)
		if err != nil {
			fmt.Printf("Error parsing published time: %v\n", err)
			continue
		}
		video := &types.Video{
			Title:        item.Snippet.Title,
			Description:  item.Snippet.Description,
			ChannelId:    item.Snippet.ChannelId,
			ChannelTitle: item.Snippet.ChannelTitle,
			VideoId:      item.Id.VideoId,
			PublishedAt:  publishedAt,
			Thumbnails: types.Thumbnails{
				Default: &types.Thumbnail{Url: item.Snippet.Thumbnails.Default.Url},
				Medium:  &types.Thumbnail{Url: item.Snippet.Thumbnails.Medium.Url},
				High:    &types.Thumbnail{Url: item.Snippet.Thumbnails.High.Url},
			},
		}
		videos = append(videos, video)
	}

	return videos, nil
}
