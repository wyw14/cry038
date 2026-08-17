package service

import (
	"testing"
	"time"

	"github.com/wyw14/cry038/internal/domain"
)

func TestBlockedCriticalRiskReminder(t *testing.T) {
	now := time.Now()
	s := domain.Session{ID: "s1", StartsAt: now.Add(time.Hour), Items: []domain.ChecklistItem{{ID: "power", Critical: true, State: domain.Blocked}}}
	reminders := Upcoming([]domain.Session{s}, now, 2*time.Hour)
	if len(reminders) != 1 || reminders[0].RiskCount != 1 {
		t.Fatalf("blocked risk missing from reminder: %+v", reminders)
	}
}
