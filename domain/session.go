package domain 




type sessionID string
type SessionState string

const (
	Active SessionState = "active"
	Deactive SessionState = "deactive"

)

func NewSession(id , tokenID string) *Session{
	return &Session{
		id:sessionID(id),
		state: Active,
		tokenID: tokenID,
	}
}


type Session struct{
	id sessionID
	state SessionState
	tokenID string
}


func (s *Session) ID() sessionID {
    return s.id
}

func (s *Session) State() SessionState {
    return s.state
}

func (s *Session) TokenID() string {
    return s.tokenID
}
func (s *Session) Activate(){
	s.state = Active
}

func (s *Session) Deactivate() {
    s.state = Deactive
}