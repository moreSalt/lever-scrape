package main

import (
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"

	// "github.com/rodaine/table"

	functions "github.com/moreSalt/lever-scrape/functions"
	t "github.com/moreSalt/lever-scrape/types"
	// "golang.org/x/exp/slices"
)

func main() {

	// Create CSVS
	content, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	// Read config.json
	var config t.Config
	err = json.Unmarshal(content, &config)
	if err != nil {
		log.Fatal("Error reading config file: ", err)
	}

	// Read companies.json
	var matches []string
	content, err = ioutil.ReadFile("./companies.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}
	err = json.Unmarshal(content, &matches)
	if err != nil {
		log.Fatal("Error reading companies file: ", err)
	}

	// Create CSVs
	allF, err := os.Create("./output/all.csv")
	if err != nil {
		log.Fatal("Error creating all jobs csv")
	}

	filteredF, err := os.Create("./output/filtered.csv")
	if err != nil {
		log.Fatal("Error creating filtered jobs csv")
	}

	allWriter := csv.NewWriter(allF)
	filteredWriter := csv.NewWriter(filteredF)

	// slices containing all of the positions, csv format
	allJobs := [][]string{
		{"Company", "Location", "Position", "URL"},
	}
	filteredJobs := [][]string{
		{"Company", "Location", "Position", "URL"},
	}

	// get required, positive, and negative keywords from config
	requ, pos, neg, err := functions.Keywords(config.Keywords)
	if err != nil {
		log.Println("Error parsing keyword types", err)
		return
	}

	// slices of all the positions
	var filtered [][]t.Job
	var all [][]t.Job

	// Create a wait group
	var wg sync.WaitGroup

	// Loop through all of the companies in companies.json
	for i := 0; i < len(matches); i++ {
		// Add one to wait group
		wg.Add(1)
		go func(index int, r []string, p []string, n []string, f *[][]t.Job, a *[][]t.Job) {
			var filteredCom []t.Job
			var allCom []t.Job

			// Greenhouse
			if strings.Contains(matches[index], "https://boards.greenhouse.io/") {
				filteredCom, allCom, err = functions.ScrapeGreenhouse(matches[index], r, p, n)

				// Lever
			} else if strings.Contains(matches[index], "https://jobs.lever.co/") {
				filteredCom, allCom, err = functions.ScrapeLever(matches[index], config.Country, r, p, n)

				// Workable
			} else if strings.Contains(matches[index], "https://apply.workable.com/") {
				filteredCom, allCom, err = functions.ScrapeWorkable(matches[index], config.Country, r, p, n)
			} else {
				log.Println("Unknown platform", matches[index])
				wg.Done()
				return
			}

			if err != nil {
				log.Println(matches[index], "ERROR:", err)
				wg.Done()
				return
			}
			*f = append(*f, filteredCom)
			*a = append(*a, allCom)
			wg.Done()
		}(i, requ, pos, neg, &filtered, &all)
	}

	// Wait for all go-routines to finish
	wg.Wait()

	log.Println("FINISHED")

	// Add all and filtered positions to the csv slices
	for i := 0; i < len(all); i++ {
		for k := 0; k < len(all[i]); k++ {
			job := all[i][k]

			allJobs = append(allJobs, []string{job.Company, job.Location, job.Position, job.PositionURL})
		}
	}

	for i := 0; i < len(filtered); i++ {
		for k := 0; k < len(filtered[i]); k++ {
			job := filtered[i][k]
			filteredJobs = append(filteredJobs, []string{job.Company, job.Location, job.Position, job.PositionURL})
		}
	}

	// Write to CSVs
	err = allWriter.WriteAll(allJobs)
	if err != nil {
		log.Println("Error writing all jobs to csv")
		return
	}

	err = filteredWriter.WriteAll(filteredJobs)
	if err != nil {
		log.Println("Error writing filtered jobs to csv")
		return
	}

}
