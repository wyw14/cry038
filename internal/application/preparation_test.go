package application

import (
	"errors"
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

func TestReviewIsolationApplication(t *testing.T) {
	p := NewPreparation(repository.NewSessionMemory(), time.Now)
	if err := p.ProposeTemplateChange("class-a", "safety", "补充绝缘手套"); err != nil {
		t.Fatal(err)
	}
	if err := p.ProposeTemplateChange("class-b", "safety", "增加护目镜"); err != nil {
		t.Fatalf("another session was treated as duplicate: %v", err)
	}
	if !errors.Is(p.ProposeTemplateChange("class-a", "safety", "重复提交"), ErrSuggestionAlreadyReviewed) {
		t.Fatal("same session and template key should remain idempotent")
	}
}
