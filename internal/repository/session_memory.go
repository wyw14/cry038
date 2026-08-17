package repository

import (
	"context"
	"errors"
	"github.com/wyw14/cry038/internal/domain"
	"sync"
)

type SessionMemory struct {
	mu   sync.RWMutex
	data map[string]domain.Session
}

func NewSessionMemory(seed ...domain.Session) *SessionMemory {
	m := &SessionMemory{data: map[string]domain.Session{}}
	for _, s := range seed {
		m.data[s.ID] = s
	}
	return m
}
func (m *SessionMemory) Get(_ context.Context, id string) (domain.Session, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	s, ok := m.data[id]
	if !ok {
		return s, errors.New("session not found")
	}
	s.Items = append([]domain.ChecklistItem(nil), s.Items...)
	s.Review = append([]string(nil), s.Review...)
	return s, nil
}
func (m *SessionMemory) Save(_ context.Context, s domain.Session) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if old, ok := m.data[s.ID]; ok && old.Version+1 != s.Version {
		return errors.New("session version conflict")
	}
	m.data[s.ID] = s
	return nil
}
