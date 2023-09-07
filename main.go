package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Response struct {
	SlackName     string    `json:"slack_name"`
	CurrentDay    string    `json:"current_day"`
	UTCTime       time.Time `json:"utc_time"`
	Track         string    `json:"track"`
	GithubFileURL string    `json:"github_file_url"`
	GithubRepoURL string    `json:"github_repo_url"`
	StatusCode    int       `json:"status_code"`
}

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

	githubFile := "https://github.com/oluu-web/hngx-stage1/blob/main/main.go"
	githubRepo := "https://github.com/oluu-web/hngx-stage1"

	// JSON response
	resp := Response{
		SlackName:     slackName,
		CurrentDay:    currentDay,
		UTCTime:       currentTime,
		Track:         track,
		GithubFileURL: githubFile,
		GithubRepoURL: githubRepo,
		StatusCode:    200,
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
