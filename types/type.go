package types

import (
	"time"
)

type (
	Video struct {
		Title       string     `json:"title"`
		Description string     `json:"description"`
		PublishTime time.Time  `json:"publish_time"`
		Thumbnails  Thumbnails `json:"thumbnails"`
	}

	Thumbnails struct {
		Default *Thumbnail `json:"default"`
		Medium  *Thumbnail `json:"medium"`
		High    *Thumbnail `json:"high"`
	}

	Thumbnail struct {
		Url string `json:"url"`
	}
)
