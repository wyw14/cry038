package http

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/wyw14/cry038/internal/application"
	"github.com/wyw14/cry038/internal/domain"
	"github.com/wyw14/cry038/internal/repository"
	"go.uber.org/zap"
)

func TestBlockedCriticalRiskHTTP(t *testing.T) {
	s := domain.Session{ID: "s1", Items: []domain.ChecklistItem{{ID: "power", Critical: true, State: domain.Blocked}}}
	app := application.NewPreparation(repository.NewSessionMemory(s), time.Now)
	router := Routes(app, zap.NewNop())
	req := httptest.NewRequest(http.MethodPost, "/api/v1/sessions/s1/lock", bytes.NewBufferString(`{"actor":"teacher"}`))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)
	if res.Code != http.StatusConflict {
		t.Fatalf("lock status=%d body=%s", res.Code, res.Body.String())
	}
}
