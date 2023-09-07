package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func getInformation(w http.ResponseWriter, r *http.Request) {
	//query parameters
	params := r.URL.Query()
	slackName := params.Get("slack_name")
	track := params.Get("track")

	//current day
	currentDay := time.Now().UTC().Format("Monday")

	//current UTC time
	currentTime := time.Now().UTC()
	timeDifference := time.Until(currentTime)
	if timeDifference.Minutes() > 2 || timeDifference.Minutes() < -2 {
		http.Error(w, "Invalid Time", http.StatusBadRequest)
		return
	}

	githubFile := "your_github_file_url"
	githubRepo := "your_github_repo_url"

	// JSON response
	resp := map[string]interface{}{
		"slack_name":      slackName,
		"current_day":     currentDay,
		"utc_time":        currentTime,
		"track":           track,
		"github_file_url": githubFile,
		"github_repo_url": githubRepo,
		"status_code":     http.StatusOK,
	}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "Error creating JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api", getInformation).Methods("GET")
	port := "8080"
	fmt.Printf("Server is running on port %s...\n", port)
	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}

}
