package data

import "time"

type User struct {
	Id       int
	Name     string
	Password string
	Email    string
}

type Session struct {
	Id        int
	Uuid      string
	Email     string
	UserId    int
	CreatedAt time.Time
}

func (s Session) Check() (ok bool, err error) {
	return true, nil
}

func (u User) CreateSession() Session {
	return Session{
		Id:     1,
		Uuid:   "1",
		Email:  u.Email,
		UserId: u.Id,
	}
}

func UserByEmail(email string) (User, error) {
	return User{
		Id:       1,
		Name:     "Haha",
		Password: "haha",
		Email:    email,
	}, nil
}

func Encrypt(p string) string {
	return p
}
