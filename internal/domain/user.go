package domain

type UserID int64

type Email string

type Name string

type User struct {
	ID    UserID
	Email Email
	Name  Name
}
