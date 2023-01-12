# Lever Scrape

A tool to scrape Lever pages for open positions based on your keywords. Originally it also used Google's search API with dorking to find a bunch of lever pages, but that only goes up to 100.

## Usage
1. create companies.json to include companies lever pages
```json
[
    "https://jobs.lever.co/joshstoysandgames",
    "https://jobs.lever.co/coupa/",
    "https://jobs.lever.co/verkada",
    "https://jobs.lever.co/Qashier",
]
```
2. Add config.json (example below)
    1. Keywords:
        1. `+` is a positive keyword, if keywords includes at least 1, then the position title needs to contain at least one of these.
        2. `~` Same as `+`, but considered it's own category, this way you can have + as the type (intern, junior, etc) and this one `~` as keyword about the position itself (software, sales, etc)
        3. `-` is a negative keyword if the position title contain it, it will not be considered.
    2. Change the country (US, UK, etc) or `ALL` to remove the country requirement
```json
{
    "keywords": [
        "+coop",
        "+co-op",
        "+intern",
        "+internship",
        "~summer",
        "-winter",
    ],
    "country": "ALL"
}
```
3. Run it `go run main.go`
4. All jobs will be in output/all.csv and filtered jobs will be in output/filtered.csv

## TODO
- [ ] Add Greenhouse
- [ ] Add Workday

