package t

type GoogleResponse struct {
	Kind              string            `json:"kind"`
	URL               URL               `json:"url"`
	Queries           Queries           `json:"queries"`
	Context           Context           `json:"context"`
	SearchInformation SearchInformation `json:"searchInformation"`
	Items             []Items           `json:"items"`
}
type URL struct {
	Type     string `json:"type"`
	Template string `json:"template"`
}
type Request struct {
	Title          string `json:"title"`
	TotalResults   string `json:"totalResults"`
	SearchTerms    string `json:"searchTerms"`
	Count          int    `json:"count"`
	StartIndex     int    `json:"startIndex"`
	InputEncoding  string `json:"inputEncoding"`
	OutputEncoding string `json:"outputEncoding"`
	Safe           string `json:"safe"`
	Cx             string `json:"cx"`
}
type NextPage struct {
	Title          string `json:"title"`
	TotalResults   string `json:"totalResults"`
	SearchTerms    string `json:"searchTerms"`
	Count          int    `json:"count"`
	StartIndex     int    `json:"startIndex"`
	InputEncoding  string `json:"inputEncoding"`
	OutputEncoding string `json:"outputEncoding"`
	Safe           string `json:"safe"`
	Cx             string `json:"cx"`
}
type Queries struct {
	Request  []Request  `json:"request"`
	NextPage []NextPage `json:"nextPage"`
}
type Context struct {
	Title string `json:"title"`
}
type SearchInformation struct {
	SearchTime            float64 `json:"searchTime"`
	FormattedSearchTime   string  `json:"formattedSearchTime"`
	TotalResults          string  `json:"totalResults"`
	FormattedTotalResults string  `json:"formattedTotalResults"`
}
type CseThumbnail struct {
	Src    string `json:"src"`
	Width  string `json:"width"`
	Height string `json:"height"`
}
type CseImage struct {
	Src string `json:"src"`
}
type Pagemap struct {
	CseThumbnail []CseThumbnail `json:"cse_thumbnail"`
	CseImage     []CseImage     `json:"cse_image"`
}
type Metatags struct {
	OgImage            string `json:"og:image"`
	TwitterTitle       string `json:"twitter:title"`
	OgImageWidth       string `json:"og:image:width"`
	Viewport           string `json:"viewport"`
	TwitterDescription string `json:"twitter:description"`
	OgTitle            string `json:"og:title"`
	OgImageHeight      string `json:"og:image:height"`
	OgURL              string `json:"og:url"`
	OgDescription      string `json:"og:description"`
	TwitterImage       string `json:"twitter:image"`
}
type Pagemap0 struct {
	CseThumbnail []CseThumbnail `json:"cse_thumbnail"`
	Metatags     []Metatags     `json:"metatags"`
	CseImage     []CseImage     `json:"cse_image"`
}

type Items struct {
	Kind             string  `json:"kind"`
	Title            string  `json:"title"`
	HTMLTitle        string  `json:"htmlTitle"`
	Link             string  `json:"link"`
	DisplayLink      string  `json:"displayLink"`
	Snippet          string  `json:"snippet"`
	HTMLSnippet      string  `json:"htmlSnippet"`
	CacheID          string  `json:"cacheId"`
	FormattedURL     string  `json:"formattedUrl"`
	HTMLFormattedURL string  `json:"htmlFormattedUrl"`
	Pagemap          Pagemap `json:"pagemap,omitempty"`
	Pagemap0         Pagemap `json:"pagemap,omitempty"`
	Pagemap1         Pagemap `json:"pagemap,omitempty"`
	Pagemap2         Pagemap `json:"pagemap,omitempty"`
	Pagemap3         Pagemap `json:"pagemap,omitempty"`
	Pagemap4         Pagemap `json:"pagemap,omitempty"`
	Pagemap5         Pagemap `json:"pagemap,omitempty"`
	Pagemap6         Pagemap `json:"pagemap,omitempty"`
	Pagemap7         Pagemap `json:"pagemap,omitempty"`
	Pagemap8         Pagemap `json:"pagemap,omitempty"`
}

type GoogleMatches struct {
	Name string
	Link string
}
