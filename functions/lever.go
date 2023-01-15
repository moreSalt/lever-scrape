package functions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	http "net/http"
	"strings"
	"time"

	// "sync"

	t "github.com/moreSalt/lever-scrape/types"
)

// Makes a req to the companies api endpoint and then looks for matches, returns to slices: all jobs, filtered jobs
func ScrapeLever(link string, location string, requ []string, pos []string, neg []string) ([]t.Job, []t.Job, error) {
	companyName := strings.Split(link, "/")[3]
	url := fmt.Sprintf("https://api.lever.co/v0/postings/%v", companyName)
	log.Println(companyName, "- Searching")

	// Create request
	client := &http.Client{
		// set the time out
		Timeout: 30 * time.Second,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// log.Println("Error creating req:", companyName)
		return nil, nil, err
	}

	// Send request
	res, err := client.Do(req)
	if err != nil {
		// log.Println("Error sending req:", companyName)
		return nil, nil, err
	}
	defer res.Body.Close()

	// Invalid statuscode
	if res.StatusCode != 200 {
		// log.Println("Error reading res body:", companyName)
		return nil, nil, err
	}

	// Read response body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		// log.Println("Error reading res body:", companyName)
		return nil, nil, err
	}

	// turn response body into a struct
	var result t.LeverRes
	err = json.Unmarshal([]byte(body), &result)
	if err != nil {
		// log.Println("Error reading converting:", companyName)
		return nil, nil, err
	}

	// var wg sync.WaitGroup
	var allSlice []t.Job
	var filSlice []t.Job

	for i := 0; i < len(result); i++ {
		job := result[i]
		matched, err := KeywordsSearch(requ, pos, neg, job.Text)
		if err != nil {
			// log.Println("Error matching", job.Text, ":", err)
			continue
		}

		formatJob := t.Job{
			Company:     companyName,
			CompanyURL:  link,
			Position:    job.Text,
			PositionURL: job.HostedURL,
			Location:    job.Country,
		}

		allSlice = append(allSlice, formatJob)
		if matched == true && (location == formatJob.Location || location == "ALL") {
			// log.Printf("%v\t%v\t%v", formatJob.Company, formatJob.Position, formatJob.PositionURL)
			filSlice = append(filSlice, formatJob)
		}

	}

	return filSlice, allSlice, nil
}
