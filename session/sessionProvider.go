package session

type SessionProvider interface {
	Set(key string, value interface{})
	SetIfAbsence(key string, value interface{}) bool
	SetEx(key string, value interface{}, seconds int64)
	Put(key string, value interface{}, setIfAbsence bool, exSeconds int64)
	Get(key string) interface{}
	Exists(key string) bool
	Remove(key string)
	TTL(key string, seconds int64) int64
	Renewal(key string, seconds int64) int64
	Persist(key string) int64
}
