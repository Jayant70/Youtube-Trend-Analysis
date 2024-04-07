package types

import "time"

type (
	Video struct {
		Title        string     `json:"title"`
		ChannelId    string     `json:"channelId"`
		ChannelTitle string     `json:"channelTitle"`
		VideoId      string     `json:"videoId"`
		Description  string     `json:"description"`
		PublishedAt  time.Time  `json:"publishedAt"`
		Thumbnails   Thumbnails `json:"thumbnails"`
	}

	Thumbnails struct {
		Default *Thumbnail `json:"default"`
		Medium  *Thumbnail `json:"medium"`
		High    *Thumbnail `json:"high"`
	}

	Thumbnail struct {
		Url string `json:"url"`
	}

	VideosResponse struct {
		Videos       []Video `json:"videosData,omitempty"`
		TotalPages   int     `json:"totalPages"`
		TotalRecords int     `json:"totalRecords"`
		PageNo       int     `json:"pageNo"`
		PageSize     int     `json:"pageSize"`
		HasNext      bool    `json:"hasNext"`
	}

	GenericResponse struct {
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
)
