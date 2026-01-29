package main

import (
    "fmt"

    "github.com/gocolly/colly/v2"
)

type JobMatch struct {
    Title    string `json:"title"`
    Company  string `json:"company"`
    Location string `json:"location"`
    ApplyURL string `json:"applyUrl"`
}

func scrapeJobs(skills []string) ([]JobMatch, error) {
    matches := []JobMatch{}
    c := colly.NewCollector()

    c.OnHTML(".job_seen_beacon", func(e *colly.HTMLElement) {
        title := e.ChildText("h2.jobTitle")
        company := e.ChildText(".companyName")
        location := e.ChildText(".companyLocation")
        link := e.ChildAttr("a", "href")

        if title != "" {
            matches = append(matches, JobMatch{
                Title:    title,
                Company:  company,
                Location: location,
                ApplyURL: "https://indeed.com" + link,
            })
        }
    })

    searchURL := fmt.Sprintf(
        "https://www.indeed.com/jobs?q=%s",
        skills[0],
    )

    c.Visit(searchURL)
    return matches, nil
}
