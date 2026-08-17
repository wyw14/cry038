package platform

import (
	"bytes"
	"encoding/csv"
	"strings"
	"testing"

	"github.com/wyw14/cry038/internal/domain"
)

func TestReplacementAccountingExport(t *testing.T) {
	s := domain.Session{Items: []domain.ChecklistItem{{Name: "陶土", Quantity: 20, Consumed: 8, State: domain.Supplemented}}}
	var out bytes.Buffer
	if err := ExportSession(&out, s); err != nil {
		t.Fatal(err)
	}
	records, err := csv.NewReader(strings.NewReader(out.String())).ReadAll()
	if err != nil {
		t.Fatal(err)
	}
	if records[1][3] != "8" {
		t.Fatalf("exported consumption=%q, want 8", records[1][3])
	}
}
