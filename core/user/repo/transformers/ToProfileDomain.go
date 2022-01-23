package userDataTransformers

import (
	"github.com/arnaugomez/buscopartida/core/user"
	"github.com/arnaugomez/buscopartida/core/user/repo/models"
)

func ToProfileDomain(p *userDbModels.Profile) *user.Profile {
	return &user.Profile{
		ID:            p.ID,
		UserID:        p.UserID,
		Slug: p.Slug,
		FirstName:     p.FirstName,
		LastName:      p.LastName,
		Description:   p.Description,
		ProfileImg:    p.ProfileImg,
		BackgroundImg: p.BackgroundImg,
	}
}

func ToProfileDb(p *user.Profile) *userDbModels.Profile {
	return &userDbModels.Profile{
		UserID:        p.UserID,
		Slug: p.Slug,
		FirstName:     p.FirstName,
		LastName:      p.LastName,
		Description:   p.Description,
		ProfileImg:    p.ProfileImg,
		BackgroundImg: p.BackgroundImg,
	}
}
