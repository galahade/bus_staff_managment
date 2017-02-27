package session

import (
	"sync"
	. "fmt"
	"io"
	"crypto/rand"
	"encoding/base64"
	. "net/http"
	"net/url"
	"time"
)

var provides = make(map[string]Provider)

type Manager struct {
	cookieName  string
	lock        sync.Mutex
	provider    Provider
	maxlifetime int64
}

type Provider interface {
	SessionInit(sid string) (Session, error)
	SessionRead(sid string) (Session, error)
	SessionDestroy(sid string) error
	SessionGC(maxLifeTime int64)
}

type Session interface {
	Set(key, value interface{}) error
	Get(key interface{}) interface{}
	Delete(key interface{}) error
	SessionID() string
}

func NewManager(provideName, cookieName string, maxlifetime int64) (*Manager, error) {
	provider, ok := provides[provideName]

	if !ok {
		return nil, Errorf("session: unknown provide %q (forgotten import?)", provideName)
	}
	return &Manager{provider: provider, cookieName: cookieName, maxlifetime: maxlifetime}, nil
}

func Register(name string, provider Provider) {
	if provider == nil {
		panic("session: Register provide is nil")
	}
	if _, dup := provides[name]; dup {
		panic("session: Register called twice for provide" + name)
	}
	provides[name] = provider
}

func (manager *Manager) sessionID() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)

}
func (manager *Manager) SessionStart(w ResponseWriter, r *Request) (session Session) {
	manager.lock.Lock()
	defer manager.lock.Unlock()

	cookie, err := r.Cookie(manager.cookieName)

	if err != nil || cookie.Value == "" {
		sid := manager.sessionID()
		session, _  = manager.provider.SessionInit(sid)
		cookie := Cookie{Name: manager.cookieName, Value: url.QueryEscape(sid), Path: "/",
			HttpOnly: true, MaxAge: int(manager.maxlifetime)}
		SetCookie(w, &cookie)
	} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		session, _ = manager.provider.SessionRead(sid)
	}
	return
}

func (manager *Manager) SessionDestroy(w ResponseWriter, r *Request)  {
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		return
	} else {
		manager.lock.Lock()
		defer manager.lock.Unlock()
		manager.provider.SessionDestroy(cookie.Value)
		expiration := time.Now()
		cookie := Cookie{Name: manager.cookieName, Path: "/", HttpOnly: true, Expires: expiration, MaxAge: -1}
		SetCookie(w, &cookie)
	}

}

func (manager *Manager) GC() {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	manager.provider.SessionGC(manager.maxlifetime)
	time.AfterFunc(time.Duration(manager.maxlifetime), func() {manager.GC()})
}



