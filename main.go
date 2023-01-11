package main

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	http "net/http"
	"os"
	"regexp"
	"strings"

	slices "golang.org/x/exp/slices"

	"github.com/joho/godotenv"

	t "github.com/moreSalt/lever-scrape/types"
	"github.com/rodaine/table"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

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

	matches, err := ScrapeGoogle(os.Getenv("API"), os.Getenv("CX"), config.StartIndex, config.EndIndex)
	if err != nil {
		log.Fatal(err)
	}

	// debug
	// example := `[{"Name":"Josh's Toys \u0026 Games - Careers","Link":"https://jobs.lever.co/joshstoysandgames"},{"Name":"Coupa Software, Inc.","Link":"https://jobs.lever.co/coupa/"},{"Name":"Qashier","Link":"https://jobs.lever.co/Qashier"},{"Name":"Nord Security","Link":"https://jobs.lever.co/nordsec"},{"Name":"Bounteous","Link":"https://jobs.lever.co/bounteous"},{"Name":"Parallel Wireless","Link":"https://jobs.lever.co/parallelwireless"},{"Name":"Qonto","Link":"https://jobs.lever.co/qonto"},{"Name":"Carbon Health","Link":"https://jobs.lever.co/carbonhealth"},{"Name":"Scaleway","Link":"https://jobs.lever.co/scaleway"},{"Name":"Token Metrics","Link":"https://jobs.lever.co/tokenmetrics"},{"Name":"Gowan Company","Link":"https://jobs.lever.co/gowanco"},{"Name":"Extreme Networks","Link":"https://jobs.lever.co/extremenetworks"},{"Name":"Atlassian","Link":"https://jobs.lever.co/atlassian"},{"Name":"Palantir Technologies","Link":"https://jobs.lever.co/palantir"},{"Name":"QANDA","Link":"https://jobs.lever.co/mathpresso"},{"Name":"MetLife","Link":"https://jobs.lever.co/metlife"},{"Name":"AppLovin","Link":"https://jobs.lever.co/applovin"},{"Name":"Alan","Link":"https://jobs.lever.co/alan"},{"Name":"Bayzat","Link":"https://jobs.lever.co/bayzat"},{"Name":"Pythian","Link":"https://jobs.lever.co/pythian"},{"Name":"Couchbase","Link":"https://jobs.lever.co/couchbase"},{"Name":"FloQast","Link":"https://jobs.lever.co/floqast"},{"Name":"Arrive Logistics","Link":"https://jobs.lever.co/arrivelogistics"},{"Name":"Xometry Inc.","Link":"https://jobs.lever.co/xometry"},{"Name":"Klarna","Link":"https://jobs.lever.co/klarna/"},{"Name":"Quotient Technology","Link":"https://jobs.lever.co/quotient"},{"Name":"DREAM","Link":"https://jobs.lever.co/wearedream/"},{"Name":"Marcus \u0026 Millichap","Link":"https://jobs.lever.co/marcusmillichap"},{"Name":"Reorg","Link":"https://jobs.lever.co/reorgresearch"},{"Name":"Klick","Link":"https://jobs.lever.co/klick"},{"Name":"Brainnest","Link":"https://jobs.lever.co/brainnest"},{"Name":"Imperfect Foods","Link":"https://jobs.lever.co/imperfectfoods"},{"Name":"Betstamp","Link":"https://jobs.lever.co/Betstamp"},{"Name":"Fi","Link":"https://jobs.lever.co/epifi"},{"Name":"Life.Church logo","Link":"https://jobs.lever.co/life"},{"Name":"Foodstuffs North Island","Link":"https://jobs.lever.co/foodstuffs"},{"Name":"Highspot","Link":"https://jobs.lever.co/highspot"},{"Name":"ZeroFox","Link":"https://jobs.lever.co/zerofox"},{"Name":"Dun \u0026 Bradstreet","Link":"https://jobs.lever.co/dnb"},{"Name":"Revolut","Link":"https://jobs.lever.co/revolut"},{"Name":"Kpler","Link":"https://jobs.lever.co/kpler"},{"Name":"Pigment","Link":"https://jobs.lever.co/pigment"},{"Name":"Taptap Send","Link":"https://jobs.lever.co/taptapsend"},{"Name":"Cirque du Soleil Entertainment Group","Link":"https://jobs.lever.co/cirquedusoleil"},{"Name":"Blue Bottle Coffee","Link":"https://jobs.lever.co/bluebottlecoffee"},{"Name":"Black Rifle Coffee","Link":"https://jobs.lever.co/blackriflecoffee"},{"Name":"AMARO","Link":"https://jobs.lever.co/amaro/"},{"Name":"Saviynt","Link":"https://jobs.lever.co/saviynt"},{"Name":"KPA","Link":"https://jobs.lever.co/kpaonline"},{"Name":"Another","Link":"https://jobs.lever.co/anotherco"},{"Name":"Sensor Tower","Link":"https://jobs.lever.co/sensortower"},{"Name":"Alation, Inc.","Link":"https://jobs.lever.co/alation"},{"Name":"Lionheart Children's Academy","Link":"https://jobs.lever.co/lionheartkid"},{"Name":"Boston Red Sox","Link":"https://jobs.lever.co/redsox"},{"Name":"Planned Parenthood of Illinois","Link":"https://jobs.lever.co/ppil"},{"Name":"Guardian Security Services","Link":"https://jobs.lever.co/guardiansecurityinc"},{"Name":"Binance","Link":"https://jobs.lever.co/binance"},{"Name":"Planned Parenthood of Wisconsin, Inc.","Link":"https://jobs.lever.co/ppwi"},{"Name":"Getty Images","Link":"https://jobs.lever.co/gettyimages"},{"Name":"SupportNinja","Link":"https://jobs.lever.co/supportninja"},{"Name":"Jam City","Link":"https://jobs.lever.co/jamcity"},{"Name":"Apartment Life","Link":"https://jobs.lever.co/apartmentlife"},{"Name":"DAZN","Link":"https://jobs.lever.co/dazn"},{"Name":"Founders Factory","Link":"https://jobs.lever.co/foundersfactory"},{"Name":"Uniphore","Link":"https://jobs.lever.co/uniphore/"},{"Name":"Nortal","Link":"https://jobs.lever.co/nortal"},{"Name":"Coins.ph","Link":"https://jobs.lever.co/coins"},{"Name":"Bowery Farming","Link":"https://jobs.lever.co/boweryfarming"},{"Name":"Rivos","Link":"https://jobs.lever.co/rivosinc"},{"Name":"Kava Labs","Link":"https://jobs.lever.co/kava"},{"Name":"Planned Parenthood of Greater New York","Link":"https://jobs.lever.co/ppgny"},{"Name":"Animoca Brands Limited","Link":"https://jobs.lever.co/animocabrands"},{"Name":"Wattpad","Link":"https://jobs.lever.co/wattpad"},{"Name":"Pediatric Healthcare Connection","Link":"https://jobs.lever.co/phcpdn"},{"Name":"Binance.US","Link":"https://jobs.lever.co/BAMTradingServices"},{"Name":"Protegrity","Link":"https://jobs.lever.co/protegrity"},{"Name":"Hinge Health","Link":"https://jobs.lever.co/hingehealth"},{"Name":"Weights \u0026 Biases","Link":"https://jobs.lever.co/wandb"},{"Name":"Neowiz","Link":"https://jobs.lever.co/neowiz"},{"Name":"Plivo","Link":"https://jobs.lever.co/plivo"},{"Name":"SFEIR","Link":"https://jobs.lever.co/sfeir"},{"Name":"Cprime","Link":"https://jobs.lever.co/cprime"},{"Name":"Spotify","Link":"https://jobs.lever.co/spotify/"},{"Name":"Foxtrot","Link":"https://jobs.lever.co/foxtrotco"},{"Name":"Adaptavist","Link":"https://jobs.lever.co/adaptavist"},{"Name":"Bullhorn, Inc.","Link":"https://jobs.lever.co/bullhorn"},{"Name":"Geotab","Link":"https://jobs.lever.co/geotab"},{"Name":"Infrastructure Ontario","Link":"https://jobs.lever.co/infrastructureontario"},{"Name":"Skydance","Link":"https://jobs.lever.co/skydance"},{"Name":"Viseven","Link":"https://jobs.lever.co/viseven"},{"Name":"Mejuri","Link":"https://jobs.lever.co/mejuri"},{"Name":"Verkada","Link":"https://jobs.lever.co/verkada"},{"Name":"Match Group","Link":"https://jobs.lever.co/matchgroup"},{"Name":"Anduril Industries","Link":"https://jobs.lever.co/anduril/"},{"Name":"Waabi Innovation Inc.","Link":"https://jobs.lever.co/waabi"},{"Name":"Aera Technology","Link":"https://jobs.lever.co/aeratechnology"},{"Name":"Window Nation","Link":"https://jobs.lever.co/windownation"},{"Name":"Mendix","Link":"https://jobs.lever.co/mendix"},{"Name":"Shield AI","Link":"https://jobs.lever.co/shieldai"},{"Name":"Hypebeast","Link":"https://jobs.lever.co/hypebeast"}]`
	// var matches []t.GoogleMatches
	// err = json.Unmarshal([]byte(example), &matches)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }
	ScrapeLever(matches, config.Country, config.Keywords)
}

func ScrapeGoogle(apiKey string, cx string, index int, endIndex int) ([]t.GoogleMatches, error) {
	var sites []t.GoogleMatches
	log.Println(apiKey, cx, index, endIndex)
	for i := index; i < endIndex; i += 10 {
		// !TODO change cx
		url := fmt.Sprintf("https://www.googleapis.com/customsearch/v1?cx=%v&q=site:jobs.lever.co+intern&key=%v&num=10&start=%v", cx, apiKey, i)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		defer res.Body.Close()

		// Invalid statuscode
		if res.StatusCode != 200 {
			log.Println(i)
			return nil, errors.New(res.Status)
		}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		var result t.GoogleResponse

		err = json.Unmarshal([]byte(body), &result)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		log.Printf("\tSEARCHING %v - %v", result.Queries.Request[0].StartIndex, result.Queries.Request[0].StartIndex+10)
		condition := len(result.Items)
		for i := 0; i < condition; i++ {
			item := result.Items[i]
			log.Printf("%v - %v", item.Title, item.Link)
			v := t.GoogleMatches{
				Name: item.Title,
				Link: item.Link,
			}
			sites = append(sites, v)
		}
	}

	return sites, nil

}

// Makes a req to the companies api endpoint and then looks for matches
func ScrapeLever(data []t.GoogleMatches, country string, keywords []string) ([]string, error) {
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

	for i := 0; i < len(data); i++ {
		company := data[i]
		log.Println("Searching:", company.Name)
		url := fmt.Sprintf("https://api.lever.co/v0/postings/%v", strings.Split(company.Link, "/")[3])
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}

		defer res.Body.Close()

		// Invalid statuscode
		if res.StatusCode != 200 {
			return nil, errors.New(res.Status)
		}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		var result t.LeverRes

		err = json.Unmarshal([]byte(body), &result)
		if err != nil {
			return nil, err
		}

		for i := 0; i < len(result); i++ {
			job := result[i]
			allJobs = append(allJobs, []string{company.Name, job.Text, job.HostedURL})
			matched, err := KeywordsSearch(requ, pos, neg, job.Text)
			if err != nil {
				continue
			}

			if matched == true && job.Country == country {
				// log.Printf("%v \t %v \t %v", name, job.Text, job.ApplyURL)
				tbl.AddRow(company.Name, job.Text, job.HostedURL)
				filteredJobs = append(filteredJobs, []string{company.Name, job.Text, job.HostedURL})

			}
		}

	}
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
