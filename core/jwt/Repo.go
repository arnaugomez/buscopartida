package jwt

type Repo interface {
	CreateJWT(userId uint) (string, error)
	GetUserIdByJWT(token string) (userId uint, err error)
}
