package user

// User Repository
type Repo interface {
	GetUser(name string) User
	CreateUser(c *Credentials) (*User, error)
	UserAlreadyExists(c *Credentials) bool
}
