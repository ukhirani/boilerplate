package types

import "time"

// HubTemplate represents a template from the bp-hub API
type HubTemplate struct {
	Clones             int       `json:"Clones"`
	Code               string    `json:"Code"`
	CreatedAt          time.Time `json:"CreatedAt"`
	Description        string    `json:"Description"`
	Documentation      string    `json:"Documentation"`
	ForkOf             int       `json:"ForkOf"`
	ForkedBoilerplates []int     `json:"ForkedBoilerplates"`
	GithubRepoLink     string    `json:"GithubRepoLink"`
	PostCmds           []string  `json:"PostCmds"`
	PreCmds            []string  `json:"PreCmds"`
	Stars              int       `json:"Stars"`
	Tags               []string  `json:"Tags"`
	TemplateID         int       `json:"TemplateID"`
	TemplateName       string    `json:"TemplateName"`
	Type               string    `json:"Type"`
	UpdatedAt          time.Time `json:"UpdatedAt"`
	Usage              string    `json:"Usage"`
	Username           string    `json:"Username"`
}
