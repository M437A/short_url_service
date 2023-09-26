package validators

import (
	"errors"
	"net/url"
	"short_url/iternal/entity"
)

func ValidateCreateShortUrl(link entity.Link) error {
	if link.UserId == 0 ||
		link.LongURL == "" {
		return errors.New("validate exception")
	}

	if !isURL(link.LongURL) {
		return errors.New("it's not a url")

	}
	return nil
}

func isURL(input string) bool {
	_, err := url.ParseRequestURI(input)
	if err != nil {
		return false
	}

	return true
}
