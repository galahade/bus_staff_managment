package memory

import (
	"container/list"
	"time"
	."github.com/galahade/bus_staff_managment/session"

	"sync"
)

var pder = &Provider{list: list.New()}

func init() {
	pder.sessions = make(map[string]*list.Element, 0)
	Register("memory", pder)
}

type SessionStore struct {
	sid string
	timeAccessed time.Time
	value map[interface{}]interface{}
}

func (st *SessionStore) Set(key, value interface{}) error {
	st.value[key] = value
	
}


type Provider struct {
	lock sync.Mutex
	sessions map[string]*list.Element
	list *list.List
}

func (pder *Provider) SessionInit(sid string) (Session, error) {
	pder.lock.Lock()
	defer pder.lock.Unlock()
	v := make(map[interface{}]interface{}, 0)
	newsess := &SessionStore{sid: sid, timeAccessed: time.Now(), value: v}
	element := pder.list.PushBack(newsess)
	pder.sessions[sid] = element
	return newsess, nil
}

func (pder *Provider) SessionRead(sid string) (Session, error) {
	if element, ok := pder.sessions[sid]; ok {
		return element.Value.(*SessionStore), nil
	} else {
		sess, err := pder.SessionInit(sid)
		return sess, err
	}
	return nil, nil
}

func (pder *Provider) SessionDestory(sid string) error {
	pder.lock.Lock()
	defer pder.lock.Unlock()

	if element, ok := pder.sessions[sid]; ok {
		delete(pder.sessions, sid)
		pder.list.Remove(element)
		return nil
	}
	return nil
}

func (pder *Provider) SessionGC(maxlifetime int64)  {
	pder.lock.Lock()
	defer pder.lock.Unlock()

	for {
		element := pder.list.Back()
		if element == nil {
			break
		}
		if (element.Value.(*SessionStore).timeAccessed.Unix() + maxlifetime) < time.Now().Unix() {
			pder.list.Remove(element)
			delete(pder.sessions, element.Value.(*SessionStore).sid)
		} else {
			break
		}
	}
}

func (pder *Provider) SessionUpdate(sid string) error {
	pder.lock.Lock()
	defer pder.lock.Unlock()
	if element, ok := pder.sessions[sid]; ok {
		element.Value.(*SessionStore).timeAccessed = time.Now()
		pder.list.MoveToFront(element)
		return nil
	}
	return nil
}
