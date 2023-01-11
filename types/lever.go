package t

type LeverRes []struct {
	AdditionalPlain  string        `json:"additionalPlain"`
	Additional       string        `json:"additional"`
	Categories       Categories    `json:"categories"`
	CreatedAt        int64         `json:"createdAt"`
	DescriptionPlain string        `json:"descriptionPlain"`
	Description      string        `json:"description"`
	ID               string        `json:"id"`
	Lists            []interface{} `json:"lists"`
	Text             string        `json:"text"`
	Country          string        `json:"country"`
	WorkplaceType    string        `json:"workplaceType"`
	HostedURL        string        `json:"hostedUrl"`
	ApplyURL         string        `json:"applyUrl"`
}
type Categories struct {
	Commitment string `json:"commitment"`
	Location   string `json:"location"`
	Team       string `json:"team"`
}
