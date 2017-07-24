package session

import (
	"math/rand"
	"sync"
	"time"
)

type Session struct {
	ID      string
	UserID  uint64
	Expires time.Time
}

type SessionStore struct {
	m        sync.RWMutex
	sessions map[string]*Session
}

func NewSessionStore() *SessionStore {
	return &SessionStore{sessions: make(map[string]*Session)}
}

// func CreateSession(userID uint64) *Session {
// 	return &Session{}
// }

// func RequestSession(id string) (*Session, error) {
// 	return &Session{}, error
// }

func (s *SessionStore) RequestSession(id string) (*Session, error) {
	s.m.RLock()
	// ...
}

const randomChars = "abcdefghijklmnopqrstuvwsyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_"

func randomString(length int) string {
	buf := make([]byte, length)
	for i := range buf {
		buf[i] = randomChars[rand.Intn(len(randomChars))]

	}
	return string(buf)
}

// func (s *SessionStore) CreateSession(userID uint64) *Session {
// 	id := randomString(32)
// 	return &Session{
// 		ID: id,
// 	}
// }

func (s *SessionStore) uniqueID() string {
	//code here
}

func (s *SessionStore) CreateSession(userID uint64) *Session {
	s.m.Lock()
	defer s.m.Unlock()

	id := s.uniqueID()
	if id == "" { // No free ID found
		return nil
	}

	Session := &Session{
		ID:     id,
		userID: userID,
		//assignment 1...
		Expires: time.Now().Add(time.Hour),
	}
	s.sessions[id] = session

	return session

}

// func (s *SessionStore) CreateSession(userID uint64) *Session {
// 	id := s.uniqueID()
// 	Session := &Session{
// 		ID: id,
// 		userID: userID,
// 		//assignment 1...
// 		Expires: time.Now().Add(time.Hour)
// 	}
// 	s.sessions[id] = session
//
// 	return session
//
// }
