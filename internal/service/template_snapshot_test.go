package service

import (
	"testing"
	"time"

	"github.com/wyw14/cry038/internal/domain"
)

func TestTemplateSnapshotPipelineReminder(t *testing.T) {
	now := time.Date(2026, 8, 17, 8, 0, 0, 0, time.UTC)
	s := domain.GenerateSession("s1", "化学", "Lab", now.Add(time.Hour), []domain.TemplateItem{{
		Key: "power", Name: "总电源", Critical: true, DefaultOwner: "teacher", Quantity: 1,
	}})
	reminders := Upcoming([]domain.Session{s}, now, 2*time.Hour)
	if len(reminders) != 1 || reminders[0].RiskCount != 1 {
		t.Fatalf("critical template risk missing from reminder: %+v", reminders)
	}
}
