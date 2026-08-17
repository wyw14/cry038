package application

import (
	"context"
	"errors"
	"github.com/wyw14/cry038/internal/domain"
	"sync"
	"time"
)

var ErrSuggestionAlreadyReviewed = errors.New("review suggestion already handled")

type SessionRepository interface {
	Get(context.Context, string) (domain.Session, error)
	Save(context.Context, domain.Session) error
}
type Preparation struct {
	repo        SessionRepository
	clock       func() time.Time
	mu          sync.Mutex
	suggestions map[string]bool
}

func NewPreparation(r SessionRepository, clock func() time.Time) *Preparation {
	return &Preparation{repo: r, clock: clock, suggestions: map[string]bool{}}
}
func (p *Preparation) Lock(ctx context.Context, id, actor string) (domain.Session, error) {
	s, err := p.repo.Get(ctx, id)
	if err != nil {
		return s, err
	}
	if err = s.Lock(actor, p.clock()); err != nil {
		return s, err
	}
	return s, p.repo.Save(ctx, s)
}
func (p *Preparation) ProposeTemplateChange(sessionID, key, text string) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	id := sessionID + "\x00" + key
	if p.suggestions[id] {
		return ErrSuggestionAlreadyReviewed
	}
	p.suggestions[id] = true
	return nil
}
