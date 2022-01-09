package user

// User Repository
type Repo interface {
	GetUser(name string) User
}