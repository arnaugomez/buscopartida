package jwt

type Repo interface {
	CreateJWT(userId uint) (string, error)
}
