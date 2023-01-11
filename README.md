# Lever Scraper

A tool to scrape for Lever pages, their open positions, and keyword matching jobs.

## Usage

- Create a .env file and add your Google Search API key
```
GOOGLE=api_key_here
```
- Configure config.json (take a look at the file for an example)
    - startIndex: where in the google search result you would like to start, starting at 1
    - count: the number of results you would like to go though. Each request to google returns a maximum of 10 results, so if you put in 50, it would make 5 requests to google.
    - keywords: + and ~ means required, the job title must have at least one of the + and ~ keywords in it. - means negative, if the job name has a negative keyword it won't be matched.
    - Country: country you would like to match with `US`