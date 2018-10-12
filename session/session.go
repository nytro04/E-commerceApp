package session

import (
	"errors"
	"math/rand"
	"sync"
	"time"
)

// Session holds a single session. UserID may be 0 to indicate that the session does not have a user (yet).
type Session struct {
	ID      string
	UserID  uint64
	Expires time.Time
}

// SessionStore offers functionality to create and store sessions.
type SessionStore struct {
	m        sync.RWMutex
	sessions map[string]*Session
}

// NewSessionStore creates and initializes a new session store.
func NewSessionStore() *SessionStore {
	return &SessionStore{sessions: make(map[string]*Session)}
}

// RequestSession retrieves an existing session from the session store or creates one when none is found. It returns an error when no session could be created.
func (s *SessionStore) RequestSession(sessionID string) (*Session, bool, error) {
	s.m.RLock()
	sess, ok := s.sessions[sessionID]
	s.m.RUnlock()

	var isNew bool
	if !ok {
		sess = s.CreateSession(0)
		isNew = true
	}

	if sess == nil {
		return nil, false, errors.New("session: Could not find free session ID")
	}
	return sess, isNew, nil
}

// randomChars specifies the characters that will be used in the session IDs. Make sure
// this string consists of 64 ASCII characters. We need to use ASCII characters because
// the randomString function generates a random index and we don't want it to use only
// parts of a Unicode rune. The number of characters should always be 2^n so the random
// function can optimize the generation of the index.
const randomChars = "abcdefghijklmnopqrstuvwsyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_"

func randomString(length int) string {
	buf := make([]byte, length)
	for i := range buf {
		buf[i] = randomChars[rand.Intn(len(randomChars))]

	}
	return string(buf)
}

func (s *SessionStore) uniqueID() string {
	// Only try 1000 times to prevent locking up the session manager.
	// If we can't find a unique ID, there is something seriously wrong with
	// the randomness on this system.
	for i := 0; i < 1000; i++ {
		id := randomString(32)
		if _, ok := s.sessions[id]; !ok {
			return id
		}
	}
	return ""
}

// CreateSession creates a new session in the session store. userID should be 0 to indicate that the session has no user attached to it.
func (s *SessionStore) CreateSession(userID uint64) *Session {
	s.m.Lock()
	defer s.m.Unlock()

	id := s.uniqueID()
	if id == "" {
		// this is extremely unlikely!
		return nil
	}

	session := &Session{
		ID:      id,
		UserID:  userID,
		Expires: time.Now().Add(time.Hour),
	}
	s.sessions[id] = session

	return session
}
