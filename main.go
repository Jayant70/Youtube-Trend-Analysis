package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"youtube/Jobs"
	"youtube/helper"
	"youtube/router"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error while getting env")
	}

	//create a new mongoDb
	helper.GetClient()
	defer helper.Client.Disconnect(context.Background())
	//create a new signal channel
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	apiKey := os.Getenv("API_KEY")
	//create a new youtube service
	youtubeService, err := youtube.NewService(context.Background(), option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Error creating YouTube client: %v", err)
	}

	fmt.Println("YouTube service initialized successfully")
	//Started go routines for polling videos
	go func() {
		Jobs.PollVideos(youtubeService)
	}()

	httpPort := os.Getenv("HTTP_PORT")

	r := router.NewRouter()

	go func() {
		err = http.ListenAndServe(":"+httpPort, r)
		if err != nil {
			fmt.Println("Error while listening to port")
			return
		}
	}()

	sig := <-sigChan
	fmt.Printf("Received signal: %v\n", sig)
	os.Exit(0)

}
