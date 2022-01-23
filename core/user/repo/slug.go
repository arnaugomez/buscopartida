package userRepo

import (
	"errors"
	"fmt"
	"github.com/arnaugomez/buscopartida/core/user/repo/models"
	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

func (r repo) GetUniqueSlug(s string) (newSlug string, err error) {
	baseSlug := slug.Make(s)
	profile := &userDbModels.Profile{}

	newSlug = baseSlug
	i := 0
	for {
		err = r.db.Where("slug = ?", newSlug).Take(profile).Error
		if err != nil {
			break
		}
		i++
		newSlug = fmt.Sprintf("%s-%d", baseSlug, i)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	}

	return newSlug, err
}
