package env

// Environment Variables Repository
type Repo interface {
	GetEnv() Env
}