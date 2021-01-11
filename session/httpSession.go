package session

type HttpSession struct {
	sessionId       string
	sessionProvider *SessionProvider
	expireTime      int64
}

func (httpSession HttpSession) GetSessionId() string {

}

func (httpSession HttpSession) GetAttr(sessionId string) interface{} {

}

func (httpSession HttpSession) SetAttr(key string, obj interface{}) {

}
func (httpSession HttpSession) SetAttEx(key string, obj interface{}) {

}
func (httpSession HttpSession) SetAttrIfAbsence(key string, obj interface{}) {

}

func (httpSession HttpSession) SetAttrIfAbsenceEx(key string, obj interface{}, seconds int64) {

}
func (httpSession HttpSession) PersistAttr(key string) {

}
func (httpSession HttpSession) TTLAttr(key string) {

}
func (httpSession HttpSession) RemoveAttr(key string) {

}
func (httpSession HttpSession) RenewalAttr(key string, seconds int64) {

}
func (httpSession HttpSession) ExistsAttr(key string) {

}
