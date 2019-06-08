package metadata

import (
	"time"

	"github.com/calebamiles/keps/pkg/keps/events"
)

type kepEvent struct {
	PrincipalField string           `json:"principal,omitempty"`
	DateField      time.Time        `json:"date,omitempty"`
	TypeField      events.Lifecycle `json:"event_type,omitempty"`
	NotesField     string           `json:"notes,omitempty"`
}

func (e *kepEvent) Type() events.Lifecycle { return e.TypeField }
func (e *kepEvent) Principal() string      { return e.PrincipalField }
func (e *kepEvent) At() time.Time          { return e.DateField }
func (e *kepEvent) Notes() string          { return e.NotesField }

type byOldestFirst []*kepEvent

func (evts byOldestFirst) Len() int      { return len(evts) }
func (evts byOldestFirst) Swap(i, j int) { evts[i], evts[j] = evts[j], evts[i] }
func (evts byOldestFirst) Less(i, j int) bool {
	return evts[i].DateField.Before(evts[j].DateField)
}
