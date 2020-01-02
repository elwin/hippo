package hippo

import (
	"github.com/elwin/hippo/pkg/crypto"
	"github.com/pkg/errors"
)

const idLength = 32

type sessionHandler struct {
	sessions map[string]*memorySession
}

func NewSessionHandler() sessionHandler {
	return sessionHandler{sessions: map[string]*memorySession{}}
}

func (s sessionHandler) Get(id string) (*memorySession, bool) {
	session, ok := s.sessions[id]
	return session, ok
}

func (s sessionHandler) New() (*memorySession, error) {
	id, err := crypto.GenerateRandomString(idLength)
	if err != nil {
		return nil, errors.Wrap(err, "create crypto token")
	}

	s.sessions[id] = NewSession(id)

	return s.sessions[id], nil
}

type Session interface {
	ID() string
	Set(key, value string)
	Get(key string) (value string, ok bool)
}

type memorySession struct {
	id    string
	store map[string]string
}

func NewSession(id string) *memorySession {
	return &memorySession{
		id:    id,
		store: map[string]string{},
	}
}

func (m memorySession) ID() string {
	return m.id
}

func (m memorySession) Set(key, value string) {
	m.store[key] = value
}

func (m memorySession) Get(key string) (string, bool) {
	value, ok := m.store[key]
	return value, ok
}
