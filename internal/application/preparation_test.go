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

func TestReplacementAccountingApplication(t *testing.T) {
	s := domain.Session{ID: "s1", Items: []domain.ChecklistItem{{ID: "clay", Name: "陶土", Quantity: 20, State: domain.Blocked}}}
	repo := repository.NewSessionMemory(s)
	p := NewPreparation(repo, time.Now)
	got, err := p.RecordReplacement(context.Background(), "s1", "clay", 8)
	if err != nil {
		t.Fatal(err)
	}
	if got.Items[0].Consumed != 8 || got.Items[0].State != domain.Supplemented {
		t.Fatalf("replacement was not recorded atomically: %+v", got.Items[0])
	}
}
