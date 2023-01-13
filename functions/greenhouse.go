package functions

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	http "net/http"
	"strings"
	"time"

	t "github.com/moreSalt/lever-scrape/types"
)

// Makes a req to the companies api endpoint and then looks for matches, returns to slices: all jobs, filtered jobs
func ScrapeGreenhouse(link string, requ []string, pos []string, neg []string) ([]t.Job, []t.Job, error) {
	companyName := strings.Split(link, "/")[3]
	url := fmt.Sprintf("https://boards-api.greenhouse.io/v1/boards/%v/jobs", companyName)
	log.Println(companyName, "- Searching")

	// Create request
	client := &http.Client{
		// set the time out
		Timeout: 30 * time.Second,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error creating req:", companyName)
		return nil, nil, err
	}

	// Send request
	res, err := client.Do(req)
	if err != nil {
		log.Println("Error sending req:", companyName)
		return nil, nil, err
	}
	defer res.Body.Close()

	// Invalid statuscode
	if res.StatusCode != 200 {
		log.Println("Error bad status:", companyName)
		return nil, nil, errors.New(res.Status)
	}

	// Read response body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("Error reading res body:", companyName)
		return nil, nil, err
	}

	// turn response body into a struct
	var result t.GreenHouseRes
	err = json.Unmarshal([]byte(body), &result)
	if err != nil {
		log.Println("Error converting:", companyName)
		return nil, nil, err
	}

	// var wg sync.WaitGroup
	var allSlice []t.Job
	var filSlice []t.Job

	for i := 0; i < len(result.Jobs); i++ {
		job := result.Jobs[i]
		matched, err := KeywordsSearch(requ, pos, neg, job.Title)
		if err != nil {
			log.Println(companyName, "- Error matching", job.Title, ":", err)
			continue
		}

		formatJob := t.Job{
			Company:     companyName,
			CompanyURL:  link,
			Position:    job.Title,
			PositionURL: job.AbsoluteURL,
			Location:    job.Location.Name,
		}

		allSlice = append(allSlice, formatJob)
		if matched == true {
			// log.Printf("%v\t%v\t%v", formatJob.Company, formatJob.Position, formatJob.PositionURL)
			filSlice = append(filSlice, formatJob)
		}

	}

	return filSlice, allSlice, nil
}
