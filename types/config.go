package t

type Config struct {
	StartIndex int      `json:"startIndex"`
	EndIndex   int      `json:"endIndex"`
	Keywords   []string `json:"keywords"`
	Country    string   `json:"country"`
}
