package user

// User Repository
type Repo interface {
	GetUserByEmail(email string) (*User, error)
	GetUserById(id uint) (*User, error)
	CreateUser(c *Credentials) (*User, error)
	DeleteUserById(id uint) error
	UserAlreadyExists(c *Credentials) bool
}
