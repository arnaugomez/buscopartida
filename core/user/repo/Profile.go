package userRepo

import (
	"github.com/arnaugomez/buscopartida/core/user"
	userDbModels "github.com/arnaugomez/buscopartida/core/user/repo/models"
	userDataTransformers "github.com/arnaugomez/buscopartida/core/user/repo/transformers"
)

func (r repo) CreateProfile(p *user.Profile) (*user.Profile, error) {
	newSlug, err := r.GetUniqueSlug(p.Slug)
	if err != nil {
		return p, err
	}
	p.Slug = newSlug

	profileDb := userDataTransformers.ToProfileDb(p)
	err = r.db.Create(profileDb).Error

	newProfile := userDataTransformers.ToProfileDomain(profileDb)
	return newProfile, err
}

func (r repo) GetProfileByUserId(userId uint) (*user.Profile, error) {
	profileDb := &userDbModels.Profile{}
	err := r.db.Where("user_id = ?", userId).Take(profileDb).Error
	return userDataTransformers.ToProfileDomain(profileDb), err
}