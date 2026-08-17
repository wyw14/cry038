package application

import (
	"context"
	"errors"
	"github.com/wyw14/cry038/internal/domain"
	"github.com/wyw14/cry038/internal/repository"
	"testing"
	"time"
)

func TestReviewSuggestionDoesNotMutateTemplateAndIsUnique(t *testing.T) {
	p := NewPreparation(repository.NewSessionMemory(), time.Now)
	if err := p.ProposeTemplateChange("s", "safety", "补充绝缘手套"); err != nil {
		t.Fatal(err)
	}
	if !errors.Is(p.ProposeTemplateChange("s", "safety", "再次提交"), ErrSuggestionAlreadyReviewed) {
		t.Fatal("expected duplicate review guard")
	}
}

func TestBlockedCriticalRiskApplication(t *testing.T) {
	s := domain.Session{ID: "s1", Items: []domain.ChecklistItem{{ID: "power", Critical: true, State: domain.Blocked}}}
	repo := repository.NewSessionMemory(s)
	p := NewPreparation(repo, time.Now)
	if _, err := p.Lock(context.Background(), "s1", "teacher"); !errors.Is(err, domain.ErrUnresolvedRisk) {
		t.Fatalf("lock error=%v, want unresolved risk", err)
	}
}
