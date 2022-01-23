package user

// Repo is the user repository
type Repo interface {
	GetUserByEmail(email string) (*User, error)
	GetUserById(id uint) (*User, error)
	CreateUser(c *Credentials) (*User, error)
	DeleteUserById(id uint) error
	UserAlreadyExists(c *Credentials) bool
	UpdatePassword(userId uint, password string) error

	CreateProfile(u *Profile) (*Profile, error)
	GetProfileByUserId(userId uint) (*Profile, error)

	GetUniqueSlug(s string) (newSlug string, err error)
}
