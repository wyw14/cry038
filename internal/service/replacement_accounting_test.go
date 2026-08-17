package service

import (
	"testing"

	"github.com/wyw14/cry038/internal/domain"
)

func TestReplacementAccountingService(t *testing.T) {
	s := domain.Session{Items: []domain.ChecklistItem{{ID: "clay", Quantity: 20, Consumed: 8, State: domain.Supplemented}}}
	if got := ReplacementUsage(s, "clay"); got != 8 {
		t.Fatalf("replacement usage=%d, want 8", got)
	}
}
