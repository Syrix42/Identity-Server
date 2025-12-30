package domain 

import "errors"

type User struct {
	userID   string
	UserName string
	Password string
	Role     string
	sessions []*Session
}

func NewUser(UserID string, Name string, password string, Role string) User {
	return User{
		userID:   UserID,
		UserName: Name,
		Password: password,
		Role:     Role,
	}
}

func (u *User) ID() string {
	return u.userID
}

func (u *User) Username() string {
	return u.UserName
}

func (u *User) GetRole() string{
	return u.Role
}

func (u *User) Sessions() []*Session {
	copies := make([]*Session, len(u.sessions))
	copy(copies, u.sessions)
	return copies
}

func (u *User) AddSession(s *Session) {
	u.sessions = append(u.sessions, s)
}

func (u *User) GetSession(SessionId sessionID) (Session, error) {
	for _, s := range u.sessions {
		if s.id == SessionId {
			return *s, nil
		}
	}
	return Session{} , errors.New("session not found ")
}


func (u *User) DeactivateSession(sessionID sessionID) error {
    s, err := u.GetSession(sessionID)
    if err != nil {
        return err
    }
    s.Deactivate()
    return nil
}
func (u *User) ActivateSession(sessionID sessionID) error {
    s, err := u.GetSession(sessionID)
    if err != nil {
        return err
    }
    s.Activate()
    return nil
}