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
	slices "golang.org/x/exp/slices"
)

// Makes a req to the companies api endpoint and then looks for matches, returns to slices: all jobs, filtered jobs
func ScrapeWorkable(link string, location []string, requ []string, pos []string, neg []string) ([]t.Job, []t.Job, error) {
	companyName := strings.Split(link, "/")[3]
	url := fmt.Sprintf("https://apply.workable.com/api/v3/accounts/%v/jobs", companyName)
	log.Println(companyName, "- Searching")

	// Create request
	client := &http.Client{
		// set the time out
		Timeout: 15 * time.Second,
	}
	req, err := http.NewRequest("POST", url, nil)
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
	var result t.WorkableRes
	err = json.Unmarshal([]byte(body), &result)
	if err != nil {
		// log.Println("Error reading converting:", companyName)
		return nil, nil, err
	}

	// var wg sync.WaitGroup
	var allSlice []t.Job
	var filSlice []t.Job

	for i := 0; i < len(result.Results); i++ {
		job := result.Results[i]
		matched, err := KeywordsSearch(requ, pos, neg, job.Title)
		if err != nil {
			// log.Println("Error matching", job.Text, ":", err)
			continue
		}

		formatJob := t.Job{
			Company:     companyName,
			CompanyURL:  link,
			Position:    job.Title,
			PositionURL: fmt.Sprintf("https://apply.workable.com/%v/j/%v/", companyName, job.Shortcode),
			Location:    job.Location.CountryCode,
		}

		allSlice = append(allSlice, formatJob)
		if matched == true && (slices.Contains(location, formatJob.Location) || (slices.Contains(location, "ALL"))) {
			// log.Printf("%v\t%v\t%v", formatJob.Company, formatJob.Position, formatJob.PositionURL)
			filSlice = append(filSlice, formatJob)
		}

	}

	return filSlice, allSlice, nil
}
