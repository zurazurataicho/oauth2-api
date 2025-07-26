package domain

type UserRepository interface {
	FindByEmail(email Email) (*User, error)
	Create(user *User) error
}
