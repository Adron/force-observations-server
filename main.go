package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/adron/force-observations-server/internal/logging"
	"github.com/kbinani/screenshot"
	"go.uber.org/zap"
)

type healthResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}

type camerasResponse struct {
	Count int `json:"count"`
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("."))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	logger := logging.GetLogger()

	response := healthResponse{
		Status:    "Service live",
		Timestamp: time.Now(),
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		logger.Error("Failed to encode response", zap.Error(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func camerasHandler(w http.ResponseWriter, r *http.Request) {
	logger := logging.GetLogger()

	// Get the number of active displays (which includes cameras on Windows)
	n := screenshot.NumActiveDisplays()

	// On Windows, each camera typically appears as a display device
	// We subtract 1 to account for the actual display
	cameraCount := n - 1
	if cameraCount < 0 {
		cameraCount = 0
	}

	response := camerasResponse{
		Count: cameraCount,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		logger.Error("Failed to encode response", zap.Error(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func main() {
	logger := logging.GetLogger()

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/cameras", camerasHandler)

	port := ":8080"
	logger.Info("Starting server", zap.String("port", port))
	if err := http.ListenAndServe(port, nil); err != nil {
		logger.Fatal("Server failed to start", zap.Error(err))
	}
}
