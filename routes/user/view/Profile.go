package userView

import "github.com/arnaugomez/buscopartida/core/user"

type Profile struct {
	ID uint
	UserID uint
	Slug string
	FirstName string
	LastName string
	Description string
	ProfileImg string
	BackgroundImg string
}

func ToProfileDomain(p *Profile) *user.Profile {
	return &user.Profile{
		ID:            p.ID,
		UserID:        p.UserID,
		Slug:          p.Slug,
		FirstName:     p.FirstName,
		LastName:      p.LastName,
		Description:   p.Description,
		ProfileImg:    p.ProfileImg,
		BackgroundImg: p.BackgroundImg,
	}
}

func ToProfileView(p *user.Profile) *Profile {
	return &Profile{
		ID:            p.ID,
		UserID:        p.UserID,
		Slug:          p.Slug,
		FirstName:     p.FirstName,
		LastName:      p.LastName,
		Description:   p.Description,
		ProfileImg:    p.ProfileImg,
		BackgroundImg: p.BackgroundImg,
	}
}

