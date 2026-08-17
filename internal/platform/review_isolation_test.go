package platform

import (
	"bytes"
	"strings"
	"testing"

	"github.com/wyw14/cry038/internal/domain"
)

func TestReviewIsolationExport(t *testing.T) {
	s := domain.Session{Review: []string{"绝缘手套不足", "陶土余量不足"}}
	var out bytes.Buffer
	if err := ExportSession(&out, s); err != nil {
		t.Fatal(err)
	}
	text := out.String()
	if !strings.Contains(text, "绝缘手套不足") || !strings.Contains(text, "陶土余量不足") {
		t.Fatalf("export omitted review history: %q", text)
	}
}
