package functions

import (
	"regexp"
	"strings"

	slices "golang.org/x/exp/slices"
)

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

		isNeg := slices.Index(negative, word)
		if isNeg != -1 {
			return false, nil
		}

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
