package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"youtube/service"
)

func GetVideosHandler(w http.ResponseWriter, r *http.Request) {
	// Pagination parameters
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 0 {
		page = 0
	}

	pageSize, err := strconv.Atoi(r.URL.Query().Get("pageSize"))
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}

	svc := service.NewVideosService()
	videosResponse, err := svc.GetVideosDetails(page, pageSize)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return

	}
	// Convert response to JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(videosResponse)
}
