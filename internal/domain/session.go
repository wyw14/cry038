package domain

import (
	"errors"
	"sort"
	"time"
)

type ItemState string

const (
	Todo         ItemState = "todo"
	Done         ItemState = "done"
	Blocked      ItemState = "blocked"
	Skipped      ItemState = "skipped"
	Supplemented ItemState = "supplemented"
)

var (
	ErrSessionStarted = errors.New("critical preparation cannot be deleted after class starts")
	ErrUnresolvedRisk = errors.New("critical preparation risk remains unresolved")
	ErrSessionLocked  = errors.New("session checklist is locked")
)

type TemplateItem struct {
	Key, Name, Kind string
	Critical        bool
	DefaultOwner    string
	Quantity        int
}
type ChecklistItem struct {
	ID, TemplateKey, Name, Kind, Owner, Note string
	Critical                                 bool
	Priority                                 int
	Quantity, Consumed                       int
	State                                    ItemState
	Timeline                                 []Event
}
type Event struct {
	At                  time.Time
	Actor, Action, Note string
}
type Session struct {
	ID, Course, Classroom string
	StartsAt              time.Time
	LockedAt              *time.Time
	Items                 []ChecklistItem
	Review                []string
	Version               uint64
}

func GenerateSession(id, course, room string, start time.Time, template []TemplateItem) Session {
	items := make([]ChecklistItem, len(template))
	for i, t := range template {
		items[i] = ChecklistItem{ID: id + "-" + t.Key, TemplateKey: t.Key, Name: t.Name, Kind: t.Kind, Owner: t.DefaultOwner, Critical: t.Critical, Quantity: t.Quantity, State: Todo}
	}
	return Session{ID: id, Course: course, Classroom: room, StartsAt: start, Items: items}
}
func (s *Session) DeleteItem(id string, now time.Time) error {
	for i, item := range s.Items {
		if item.ID != id {
			continue
		}
		if s.LockedAt != nil {
			return ErrSessionLocked
		}
		if !now.Before(s.StartsAt) && item.Critical {
			return ErrSessionStarted
		}
		s.Items = append(s.Items[:i], s.Items[i+1:]...)
		s.Version++
		return nil
	}
	return nil
}
func (s *Session) Lock(actor string, now time.Time) error {
	for _, item := range s.Items {
		if item.Critical && (item.State == Todo || item.State == Blocked) {
			return ErrUnresolvedRisk
		}
	}
	s.LockedAt = &now
	s.Version++
	for i := range s.Items {
		s.Items[i].Timeline = append(s.Items[i].Timeline, Event{now, actor, "session_locked", ""})
	}
	return nil
}
func (s *Session) SetState(id string, state ItemState, actor, note string, now time.Time) error {
	for i := range s.Items {
		if s.Items[i].ID == id {
			s.Items[i].State = state
			s.Items[i].Note = note
			s.Items[i].Timeline = append(s.Items[i].Timeline, Event{now, actor, "state:" + string(state), note})
			s.Version++
			return nil
		}
	}
	return nil
}
func (s *Session) RecordReplacement(id string, replacementUsed int) error {
	for i := range s.Items {
		if s.Items[i].ID == id {
			s.Items[i].Consumed += replacementUsed
			s.Items[i].State = Supplemented
			s.Version++
			return nil
		}
	}
	return nil
}
func (s Session) Risks() []string {
	out := []string{}
	for _, i := range s.Items {
		if i.Critical && (i.State == Todo || i.State == Blocked) {
			out = append(out, i.ID)
		}
	}
	sort.Strings(out)
	return out
}
