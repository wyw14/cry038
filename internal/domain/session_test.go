package domain

import (
	"errors"
	"testing"
	"time"
)

func baseSession() Session {
	start := time.Date(2026, 8, 17, 9, 0, 0, 0, time.UTC)
	return GenerateSession("s1", "陶艺", "C-12", start, []TemplateItem{{Key: "kiln", Name: "窑炉断电检查", Kind: "safety", Critical: true, DefaultOwner: "teacher", Quantity: 1}, {Key: "clay", Name: "陶土", Kind: "material", Critical: false, DefaultOwner: "assistant", Quantity: 20}})
}
func TestCriticalItemCannotBeDeletedAfterStart(t *testing.T) {
	s := baseSession()
	if !errors.Is(s.DeleteItem("s1-kiln", s.StartsAt.Add(time.Minute)), ErrSessionStarted) {
		t.Fatal("expected protection")
	}
}
func TestTemplateChangesDoNotRewriteGeneratedSession(t *testing.T) {
	tpl := []TemplateItem{{Key: "g", Name: "护目镜", Critical: true, Quantity: 10}}
	s := GenerateSession("s", "化学", "Lab", time.Now(), tpl)
	tpl[0].Name = "防护服"
	tpl[0].Quantity = 99
	if s.Items[0].Name != "护目镜" || s.Items[0].Quantity != 10 {
		t.Fatalf("history changed: %+v", s.Items[0])
	}
}
func TestLockRejectsBlockedCriticalWork(t *testing.T) {
	s := baseSession()
	s.Items[0].State = Blocked
	if !errors.Is(s.Lock("teacher", time.Now()), ErrUnresolvedRisk) {
		t.Fatal("lock should fail")
	}
}
func TestReplacementDoesNotDoubleCountOriginalQuantity(t *testing.T) {
	s := baseSession()
	if err := s.RecordReplacement("s1-clay", 8); err != nil {
		t.Fatal(err)
	}
	if s.Items[1].Consumed != 8 {
		t.Fatalf("consumed=%d", s.Items[1].Consumed)
	}
}
