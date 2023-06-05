package model

type User struct {
	ID       uint
	Username string
	Password string
}
type Income struct {
	ID          uint
	IDUser      string
	Description string
	Money       uint
}
