package repository

import (
	"context"
	"testing"

	"github.com/wyw14/cry038/internal/domain"
)

func TestReviewIsolationRepository(t *testing.T) {
	repo := NewSessionMemory(
		domain.Session{ID: "a", Review: []string{"绝缘手套不足"}},
		domain.Session{ID: "b", Review: []string{"陶土余量不足"}},
	)
	a, err := repo.Get(context.Background(), "a")
	if err != nil {
		t.Fatal(err)
	}
	b, err := repo.Get(context.Background(), "b")
	if err != nil {
		t.Fatal(err)
	}
	if len(a.Review) != 1 || a.Review[0] != "绝缘手套不足" || len(b.Review) != 1 || b.Review[0] != "陶土余量不足" {
		t.Fatalf("reviews crossed session boundary: a=%v b=%v", a.Review, b.Review)
	}
	a.Review[0] = "tampered"
	again, err := repo.Get(context.Background(), "a")
	if err != nil {
		t.Fatal(err)
	}
	if again.Review[0] != "绝缘手套不足" {
		t.Fatalf("read result mutated stored review: %v", again.Review)
	}
}
