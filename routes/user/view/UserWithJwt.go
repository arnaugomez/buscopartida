package userView

import "github.com/arnaugomez/buscopartida/core/user"

type UserWithJwt struct {
	Jwt  string `json:"jwt"`
	User *User  `json:"user"`
}

func ToUserWithJwtView(jwt string, usr *user.User) *UserWithJwt {
	return &UserWithJwt{jwt, ToUserView(usr)}
}
