package platform

import (
	"bytes"
	"strings"
	"testing"

	"github.com/wyw14/cry038/internal/domain"
)

func TestTemplateSnapshotPipelineExport(t *testing.T) {
	s := domain.Session{Items: []domain.ChecklistItem{{Name: "陶土", Owner: "assistant", State: domain.Todo}}}
	var out bytes.Buffer
	if err := ExportSession(&out, s); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(out.String(), "assistant") {
		t.Fatalf("export lost snapshotted owner: %q", out.String())
	}
}
