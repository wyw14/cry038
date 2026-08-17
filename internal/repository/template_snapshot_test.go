package repository

import (
	"context"
	"testing"
	"time"

	"github.com/wyw14/cry038/internal/domain"
)

func TestTemplateSnapshotPipelineRepository(t *testing.T) {
	s := domain.Session{ID: "s1", StartsAt: time.Now(), Items: []domain.ChecklistItem{{
		ID: "s1-kiln", TemplateKey: "kiln", Name: "窑炉", Kind: "equipment",
		Owner: "teacher", Critical: true, Quantity: 1, State: domain.Todo,
	}}}
	repo := NewSessionMemory(s)
	got, err := repo.Get(context.Background(), "s1")
	if err != nil {
		t.Fatal(err)
	}
	item := got.Items[0]
	if item.TemplateKey != "kiln" || item.Owner != "teacher" || !item.Critical || item.Quantity != 1 {
		t.Fatalf("stored snapshot metadata was lost: %+v", item)
	}
}
