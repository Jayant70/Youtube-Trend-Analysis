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

func GetVideosDataByQueryHandler(w http.ResponseWriter, r *http.Request) {
	// Pagination parameters
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 0 {
		page = 0
	}

	pageSize, err := strconv.Atoi(r.URL.Query().Get("pageSize"))
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}

	query := r.URL.Query().Get("query")
	if query == "" {
		http.Error(w, "Query parameter is required", http.StatusBadRequest)
		return
	}

	svc := service.NewVideosService()

	videoResponse, err := svc.GetVideosDetailsByQuery(page, pageSize, query)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Convert response to JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(videoResponse)
}
