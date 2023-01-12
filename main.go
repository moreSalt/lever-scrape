package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	http "net/http"
	"os"
	"regexp"
	"strings"
	"sync"

	"github.com/rodaine/table"
	slices "golang.org/x/exp/slices"

	t "github.com/moreSalt/lever-scrape/types"
)

func main() {

	content, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	// Now let's unmarshall the data into `payload`
	var config t.Config
	err = json.Unmarshal(content, &config)
	if err != nil {
		log.Fatal("Error reading config file: ", err)
	}

	var matches []string
	content, err = ioutil.ReadFile("./companies.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}
	err = json.Unmarshal(content, &matches)
	if err != nil {
		log.Fatal("Error reading companies file: ", err)
	}

	ScrapeLever(matches, config.Country, config.Keywords)
}

// Makes a req to the companies api endpoint and then looks for matches
func ScrapeLever(data []string, country string, keywords []string) ([]string, error) {
	// Create CSV
	allF, err := os.Create("./output/all.csv")
	if err != nil {
		return nil, err
	}

	filteredF, err := os.Create("./output/filtered.csv")
	if err != nil {
		return nil, err
	}
	allWriter := csv.NewWriter(allF)
	filteredWriter := csv.NewWriter(filteredF)

	allJobs := [][]string{
		{"Company", "Position", "URL"},
	}
	filteredJobs := [][]string{
		{"Company", "Position", "URL"},
	}

	// CREATE TABLE
	table.DefaultHeaderFormatter = func(format string, vals ...interface{}) string {
		return strings.ToUpper(fmt.Sprintf(format, vals...))
	}

	tbl := table.New("Company", "Position", "URL")

	// Get req, pos, and neg keywords

	requ, pos, neg, err := Keywords(keywords)
	if err != nil {
		return nil, err
	}
	var wg sync.WaitGroup
	for i := 0; i < len(data); i++ {
		wg.Add(1)
		go func(k int, all *[][]string, fil *[][]string) {
			companyName := strings.Split(data[k], "/")[3]
			log.Println(companyName, "- Searching")
			url := fmt.Sprintf("https://api.lever.co/v0/postings/%v", companyName)
			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				log.Println(companyName, "-  creating req:", err)
			}

			res, err := http.DefaultClient.Do(req)
			if err != nil {
				log.Println(companyName, "-  with res:", err)
			}

			defer res.Body.Close()

			// Invalid statuscode
			if res.StatusCode != 200 {
				log.Println(companyName, "-  wrong status code:", res.StatusCode)
			}

			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				log.Println(companyName, "-  reading body", err)
			}

			var result t.LeverRes

			err = json.Unmarshal([]byte(body), &result)
			if err != nil {
				log.Println(companyName, "- Error unmarshaling JSON res body", err)
			}

			for i := 0; i < len(result); i++ {
				job := result[i]
				// allJobs = append(allJobs, []string{companyName, job.Text, job.HostedURL})
				*all = append(*all, []string{companyName, job.Text, job.HostedURL})
				matched, err := KeywordsSearch(requ, pos, neg, job.Text)
				if err != nil {
					continue
				}

				if matched == true && (job.Country == country || job.Country == "ALL") {
					// log.Printf("%v \t %v \t %v", name, job.Text, job.ApplyURL)
					tbl.AddRow(companyName, job.Text, job.HostedURL)
					// filteredJobs = append(filteredJobs, []string{companyName, job.Text, job.HostedURL})
					*fil = append(*fil, []string{companyName, job.Text, job.HostedURL})
				}
			}
			wg.Done()
		}(i, &allJobs, &filteredJobs)

	}
	wg.Wait()
	tbl.Print()

	err = allWriter.WriteAll(allJobs)
	if err != nil {
		return nil, err
	}

	err = filteredWriter.WriteAll(filteredJobs)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// Go through the job name and see if it matches with the req, pos, and neg keywords
func KeywordsSearch(required []string, positive []string, negative []string, jobName string) (bool, error) {

	// Format the productName, remove all symbols and lowercase the string
	// var re = regexp.MustCompile(`[\-]`)
	// name := re.ReplaceAllString(jobName, " ")
	name := strings.ToLower(jobName)
	re := regexp.MustCompile(`[^\w\s]`)
	name = re.ReplaceAllString(name, "")
	nameSlice := strings.Split(name, " ")

	// Number of matching keywords
	reqCount := 0

	posCount := 0

	// Delim for the loop
	delim := len(nameSlice)

	// Loop through job name looking at each word to see if it matches any kw
	for i := 0; i < delim; i++ {
		word := nameSlice[i]

		isReq := slices.Index(required, word)
		if isReq != -1 {
			reqCount += 1
			continue
		}

		isPos := slices.Index(positive, word)
		if isPos != -1 {
			posCount++
			continue
		}

		isNeg := slices.Index(negative, word)
		if isNeg != -1 {
			return false, nil
		}

	}

	if (reqCount > 0 || len(required) == 0) && (posCount > 0 || len(positive) == 0) {
		return true, nil
	}

	return false, nil
}

// Split a slice into sorted slices neg, req, and pos
func Keywords(keywords []string) ([]string, []string, []string, error) {

	delim := len(keywords)
	var neg []string
	var req []string
	var pos []string

	for i := 0; i < delim; i++ {
		word := keywords[i]
		if string(word[0]) == "+" {
			req = append(req, strings.ToLower(word[1:]))
		} else if string(word[0]) == "-" {
			neg = append(neg, strings.ToLower(word[1:]))
		} else if string(word[0]) == "~" {
			pos = append(pos, strings.ToLower(word[1:]))
		}
	}

	return req, pos, neg, nil
}
