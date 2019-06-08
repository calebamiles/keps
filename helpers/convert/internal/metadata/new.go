package metadata

import (
	"time"
)

type New struct {
	AuthorsField           []string  `json:"authors,omitempty"`
	TitleField             string    `json:"title,omitempty"`
	ShortIDField           *int      `json:"kep_number,omitempty"`
	ReviewersField         []string  `json:"reviewers,omitempty"`
	ApproversField         []string  `json:"approvers,omitempty"`
	EditorsField           []string  `json:"editors,omitempty"`
	StateField             string    `json:"state,omitempty"`
	ReplacesField          []string  `json:"replaces,omitempty"`
	SupersededByField      []string  `json:"superseded_by,omitempty"`
	DevelopmentThemesField []string  `json:"development_themes,omitempty"`
	LastUpdatedField       time.Time `json:"last_updated,omitempty"`
	CreatedField           time.Time `json:"created,omitempty"`
	UniqueIDField          string    `json:"uuid,omitempty"`
	SectionsField          []string  `json:"sections,omitempty"`

	OwningSIGField           string   `json:"owning_sig,omitempty"`
	AffectedSubprojectsField []string `json:"affected_subprojects,omitempty"`
	ParticipatingSIGsField   []string `json:"participating_sigs,omitempty"`
	KubernetesWideField      bool     `json:"kubernetes_wide,omitempty"`
	SIGWideField             bool     `json:"sig_wide,omitempty"`
}

func (n *New) Authors() []string      { return n.AuthorsField }
func (n *New) Title() string          { return n.TitleField }
func (n *New) State() string          { return n.StateField }
func (n *New) LastUpdated() time.Time { return n.LastUpdatedField }
func (n *New) Created() time.Time     { return n.CreatedField }
func (n *New) OwningSIG() string      { return n.OwningSIGField }
