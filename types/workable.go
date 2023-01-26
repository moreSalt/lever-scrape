package t

import (
	"time"
)

type WorkableRes struct {
	Total    int       `json:"total,omitempty"`
	Results  []Results `json:"results,omitempty"`
	NextPage string    `json:"nextPage,omitempty"`
}

type Location struct {
	Country     string `json:"country,omitempty"`
	CountryCode string `json:"countryCode,omitempty"`
	City        string `json:"city,omitempty"`
	Region      string `json:"region,omitempty"`
}

// Ind job listing
type Results struct {
	ID             int       `json:"id,omitempty"`
	Shortcode      string    `json:"shortcode,omitempty"`
	Title          string    `json:"title,omitempty"`
	Remote         bool      `json:"remote,omitempty"`
	Location       Location  `json:"location,omitempty"`
	State          string    `json:"state,omitempty"`
	IsInternal     bool      `json:"isInternal,omitempty"`
	Code           string    `json:"code,omitempty"`
	Published      time.Time `json:"published,omitempty"`
	Type           string    `json:"type,omitempty"`
	Language       string    `json:"language,omitempty"`
	Department     []string  `json:"department,omitempty"`
	AccountUID     string    `json:"accountUid,omitempty"`
	ApprovalStatus string    `json:"approvalStatus,omitempty"`
}
