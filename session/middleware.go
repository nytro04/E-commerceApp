package session

import (
	"net/http"
	"time"
	"context"
)

type contextKey struct{}

type contextWrapper interface {
	Context() context.Context
	WithContext(context.Context) *http.Request
}

func (s *SessionStore) Wrap(h func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			sess  *Session
			isNew bool
			err   error
		)
		
		c, err := r.Cookie("nytroshop_session")
		if err != nil {
			// There is no session cookie, create a new one
			sess = s.CreateSession(0)
			if sess == nil {
				http.Error(w, "Unable to create new session", http.StatusInternalServerError)
				return
			}
			
			isNew = true
		} else {
			// Try load the session
			sess, isNew, err = s.RequestSession(c.Value)
			// Handle the error
				if err != nil {
				http.Error(w, "Unable to request session", http.StatusInternalServerError)
				return 
			}
		}
		
		// Write the cookie
		if isNew {
			http.SetCookie(w, &http.Cookie{
				Name: "nytroshop_session",
				Value: sess.ID,
				Expires: time.Now().Add(30 * time.Minute),
			})
		}

		var i interface{} = r
		if cw, ok := i.(contextWrapper); ok {
			r = cw.WithContext(context.WithValue(cw.Context(), contextKey{}, sess))
		}
		
		h(w, r)
	}
}

func (s *SessionStore) SessionForRequest(r *http.Request) *Session {
		var i interface{} = r
		if cw, ok := i.(contextWrapper); ok {
			return cw.Context().Value(contextKey{}).(*Session)
		} else {
			panic("Old go version")
		}
}
