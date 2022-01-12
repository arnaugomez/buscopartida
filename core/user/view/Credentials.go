package userView

import "github.com/arnaugomez/buscopartida/core/user"

type Credentials struct {
	Name     string
	Email    string
	Password string
}

func ToCredentialsDomain(c *Credentials) *user.Credentials {
	return &user.Credentials{
		Name:     c.Name,
		Email:    c.Email,
		Password: c.Password,
	}
}