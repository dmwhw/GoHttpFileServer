package session

type Session interface {
	GetSessionId() string
	GetAttr(sessionId string) interface{}
	SetAttr(key string, obj interface{})
	SetAttEx(key string, obj interface{})
	SetAttrIfAbsence(key string, obj interface{})
	SetAttrIfAbsenceEx(key string, obj interface{}, seconds int64)
	PersistAttr(key string)
	TTLAttr(key string)
	RemoveAttr(key string)
	RenewalAttr(key string, seconds int64)
	ExistsAttr(key string)
	SetExpireTime(seconds int64)
	inAvaliable()
}
