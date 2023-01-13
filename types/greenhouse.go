package t

type GreenHouseRes struct {
	Jobs []jobs `json:"jobs,omitempty"`
	Meta meta   `json:"meta,omitempty"`
}
type dataCompliance struct {
	Type            string      `json:"type,omitempty"`
	RequiresConsent bool        `json:"requires_consent,omitempty"`
	RetentionPeriod interface{} `json:"retention_period,omitempty"`
}
type location struct {
	Name string `json:"name,omitempty"`
}
type metadata struct {
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Value     string `json:"value,omitempty"`
	ValueType string `json:"value_type,omitempty"`
}
type jobs struct {
	AbsoluteURL    string           `json:"absolute_url,omitempty"`
	DataCompliance []dataCompliance `json:"data_compliance,omitempty"`
	InternalJobID  int              `json:"internal_job_id,omitempty"`
	Location       location         `json:"location,omitempty"`
	Metadata       []interface{}    `json:"metadata,omitempty"`
	ID             int              `json:"id,omitempty"`
	UpdatedAt      string           `json:"updated_at,omitempty"`
	RequisitionID  interface{}      `json:"requisition_id,omitempty"`
	Title          string           `json:"title,omitempty"`
}
type meta struct {
	Total int `json:"total,omitempty"`
}
