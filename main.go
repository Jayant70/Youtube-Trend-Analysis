package main

import (
	"context"
	"fmt"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"log"
	"os"
	"os/signal"
	"syscall"
	"youtube/Jobs"
	"youtube/constants"
	"youtube/helper"
)

func main() {
	//create a new mongoDb
	helper.GetClient()
	defer helper.Client.Disconnect(context.Background())
	//create a new signal channel
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	//create a new youtube service
	youtubeService, err := youtube.NewService(context.Background(), option.WithAPIKey(constants.ApiKey))
	if err != nil {
		log.Fatalf("Error creating YouTube client: %v", err)
	}

	fmt.Println("YouTube service initialized successfully")
	//Started go routines for polling videos
	go func() {
		Jobs.PollVideos(youtubeService)
	}()

	sig := <-sigChan
	fmt.Printf("Received signal: %v\n", sig)
	os.Exit(0)

}
