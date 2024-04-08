package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
	"youtube/helper"
	"youtube/types"
)

type VideosService interface {
	GetVideosDetails(page int, pageSize int) (types.VideosResponse, error)
	GetVideosDetailsByQuery(page int, pageSize int, query string) (types.VideosResponse, error)
}

type videosService struct {
}

func (s *videosService) GetVideosDetails(page int, pageSize int) (types.VideosResponse, error) {
	limit := pageSize
	// Sorting parameter
	sortBy := bson.D{{"publishedat", -1}}
	// Add validation for sortBy if needed
	findOptions := options.Find()
	findOptions.SetSkip(int64(page * limit))
	findOptions.SetLimit(int64(limit))
	findOptions.SetSort(sortBy)

	dbName := os.Getenv("DB_NAME")
	collectionName := os.Getenv("COLLECTION_NAME")
	database := helper.Client.Database(dbName)
	collection := database.Collection(collectionName)

	err := helper.Client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Could not connect to MongoDB:", err)
	}

	// Set the context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{}

	// Define a filter (can be empty to fetch all documents)
	cursor, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		log.Fatal("Error finding documents:", err)
	}
	defer cursor.Close(ctx)

	var videos []types.Video
	for cursor.Next(ctx) {
		var video types.Video
		if err := cursor.Decode(&video); err != nil {
			log.Fatal(err)
		}
		videos = append(videos, video)
	}
	// Total items count
	totalRecords, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	// Calculate pagination information
	totalPages := (totalRecords + int64(limit) - 1) / int64(limit)
	hasNext := page < int(totalPages-1)

	// Construct response
	response := types.VideosResponse{
		Videos:       videos,
		TotalPages:   int(totalPages),
		TotalRecords: int(totalRecords),
		PageNo:       page,
		PageSize:     limit,
		HasNext:      hasNext,
	}
	return response, nil
}

func (s *videosService) GetVideosDetailsByQuery(page int, pageSize int, query string) (types.VideosResponse, error) {
	limit := pageSize

	dbName := os.Getenv("DB_NAME")
	collectionName := os.Getenv("COLLECTION_NAME")
	database := helper.Client.Database(dbName)
	collection := database.Collection(collectionName)

	// Compile a regular expression to match the query in a fuzzy way
	fuzzyRegex := regexp.QuoteMeta(query)

	// Create a regex pattern for fuzzy search
	pattern := strings.Join([]string{".*", fuzzyRegex, ".*"}, "")

	// Create a filter for fuzzy search using the regex pattern
	//filter := bson.M{"title": bson.M{"$regex": pattern, "$options": "i"}}
	filter := bson.M{
		"$or": []bson.M{
			{"title": bson.M{"$regex": pattern, "$options": "i"}},
			{"description": bson.M{"$regex": pattern, "$options": "i"}},
		},
	}

	// Set options for pagination
	findOptions := options.Find()
	findOptions.SetSkip(int64(page * pageSize))
	findOptions.SetLimit(int64(pageSize))

	// Perform the fuzzy search
	cursor, err := collection.Find(context.Background(), filter, findOptions)
	if err != nil {
		log.Fatal("Error finding documents:", err)
		return types.VideosResponse{}, err
	}
	defer cursor.Close(context.Background())

	// Convert MongoDB cursor to Video slice
	var videos []types.Video
	for cursor.Next(context.Background()) {
		var video types.Video
		if err := cursor.Decode(&video); err != nil {
			log.Fatal(err)
			return types.VideosResponse{}, err
		}
		videos = append(videos, video)
	}

	// Total items count
	totalRecords, err := collection.CountDocuments(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	// Calculate pagination information
	totalPages := (totalRecords + int64(limit) - 1) / int64(limit)
	hasNext := page < int(totalPages-1)

	// Construct response
	response := types.VideosResponse{
		Videos:       videos,
		TotalPages:   int(totalPages),
		TotalRecords: int(totalRecords),
		PageNo:       page,
		PageSize:     limit,
		HasNext:      hasNext,
	}

	return response, nil
}

func NewVideosService() VideosService {
	return &videosService{}
}
