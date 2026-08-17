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
