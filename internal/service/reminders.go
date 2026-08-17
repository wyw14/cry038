package service

import (
	"github.com/wyw14/cry038/internal/domain"
	"sort"
	"time"
)

type Reminder struct {
	SessionID string
	StartsAt  time.Time
	RiskCount int
}

func Upcoming(sessions []domain.Session, now time.Time, window time.Duration) []Reminder {
	out := []Reminder{}
	for _, s := range sessions {
		if s.StartsAt.After(now) && !s.StartsAt.After(now.Add(window)) {
			out = append(out, Reminder{s.ID, s.StartsAt, len(s.Risks())})
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].StartsAt.Before(out[j].StartsAt) })
	return out
}
