package session

type SessionManager interface {
	GetSession(sessionId string) (session *Session)
	Create() (session *Session)
	Add(session *Session)
	RemoveSession(sessionId string)
	Exists(sessionId string)
	GC()
	Destory()
}
type SessionManager struct {
	expireSeconds   int64
	sessionPool     map[string]*session
	sessionListener *sessionListener
}

func New() {

}
