package validators

import (
	"errors"
	"short_url/iternal/entity"
)

func ValidateCreateShortUrl(link entity.Link) error {
	if link.UserId == 0 ||
		link.LongURL == "" {
		return errors.New("validate exception")
	}
	return nil
}
