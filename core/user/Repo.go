package user

// User Repository
type Repo interface {
	GetUserByEmail(email string) (*User, error)
	CreateUser(c *Credentials) (*User, error)
	UserAlreadyExists(c *Credentials) bool
}
