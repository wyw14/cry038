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
			riskCount := 0
			for _, item := range s.Items {
				if item.Critical && item.State == domain.Todo {
					riskCount++
				}
			}
			out = append(out, Reminder{s.ID, s.StartsAt, riskCount})
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].StartsAt.Before(out[j].StartsAt) })
	return out
}
