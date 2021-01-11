package session

type SessionListener interface {
	onCreated(sessionId string)
	onDestroyed(sessionId string, bool isExpired)
}
