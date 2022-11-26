package utils

import (
	"github.com/asaskevich/govalidator"
)

func IsValidURL(url string) bool {
	return govalidator.IsURL(url)
}
