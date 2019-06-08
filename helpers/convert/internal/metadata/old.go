package metadata

import (
	"time"
)

type Old struct {
	KepNumber         string    `json:"kep_number,omitempty"`
	Title             string    `json:"title,omitempty"`
	Status            string    `json:"status,omitempty"`
	Authors           []string  `json:"authors"`
	OwningSIG         string    `json:"owning-sig,omitempty"`
	ParticipatingSIGs []string  `json:"participating-sigs,omitempty"`
	Reviewers         []string  `json:"reviewers,omitempty"`
	Approvers         []string  `json:"approvers,omitempty"`
	Editor            string    `json:"editor,omitempty"`
	CreationDate      time.Time `json:"creation-date,omitempty"`
	LastUpdated       time.Time `json:"last-updated,omitempty"`
	SeeAlso           []string  `json:"see-also,omitempty"`
	Replaces          []string  `json:"replaces,omitempty"`
	SupersededBy      []string  `json:"superseded-by,omitempty"`
}
