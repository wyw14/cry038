package repository

import (
	"context"
	"testing"
	"time"

	"github.com/wyw14/cry038/internal/domain"
)

func TestTimelineHistoryPipelineRepository(t *testing.T) {
	now := time.Now()
	s := domain.Session{ID: "s1", Items: []domain.ChecklistItem{{ID: "i1", Timeline: []domain.Event{
		{At: now, Actor: "assistant", Action: "state:blocked", Note: "缺货"},
		{At: now.Add(time.Minute), Actor: "teacher", Action: "state:supplemented", Note: "已替换"},
	}}}}
	repo := NewSessionMemory(s)
	got, err := repo.Get(context.Background(), "s1")
	if err != nil {
		t.Fatal(err)
	}
	if len(got.Items[0].Timeline) != 2 {
		t.Fatalf("repository truncated audit history: %+v", got.Items[0].Timeline)
	}
	got.Items[0].Timeline[0].Note = "tampered"
	again, err := repo.Get(context.Background(), "s1")
	if err != nil {
		t.Fatal(err)
	}
	if again.Items[0].Timeline[0].Note != "缺货" {
		t.Fatalf("read result mutated stored history: %+v", again.Items[0].Timeline)
	}
}
