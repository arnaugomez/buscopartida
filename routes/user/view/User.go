package userView

import "github.com/arnaugomez/buscopartida/core/user"

type User struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func ToUserView(u *user.User) *User {
	return &User{u.ID, u.Name, u.Email}
}
