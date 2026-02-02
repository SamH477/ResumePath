package main

import (
    "encoding/json"
    "net/http"
)

func handleJobMatches(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
    w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
    w.Header().Set("Content-Type", "application/json")

    if r.Method == http.MethodOptions {
        w.WriteHeader(http.StatusOK)
        return
    }

    if r.Method != http.MethodPost {
        http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
        return
    }

    err := r.ParseMultipartForm(10 << 20)
    if err != nil {
        http.Error(w, "Invalid form", http.StatusBadRequest)
        return
    }

    file, _, err := r.FormFile("resume")
    if err != nil {
        http.Error(w, "Resume missing", http.StatusBadRequest)
        return
    }
    defer file.Close()

	job := r.FormValue("job")
    city := r.FormValue("city")
    state := r.FormValue("state")

    // TEMP: mocked skill extraction
    skills := []string{"software engineer", "go", "react", "python"}

    jobs, err := scrapeJobs(job, city, state)
    if err != nil {
        http.Error(w, "Scraping failed", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(jobs)
}
