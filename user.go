package event_channel

import "fmt"

type User struct {
	SubscriberDefault
	Username string
}

func NewUser(username string) *User {
	return &User{
		Username: username,
	}
}

func (u *User) OnReceive(msg string) {
	fmt.Printf("MESSAGE GOT: %s: %s\n", u.Username, msg)
}

func (u *User) OnUnsubscribe(msg string) {
	fmt.Printf("UNSUBSCRIBE: %s: %s\n", u.Username, msg)
}

func (u *User) GetID() string {
	return u.Username
}
