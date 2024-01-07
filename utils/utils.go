package utils

import (
	"regexp"
)

func IsStringUrl(str string) bool {
	urlRegex := regexp.MustCompile(`^(https?|ftp):\/\/[^\s\/$.?#].[^\s]*$|^(localhost|127\.0\.0\.1):[0-9]+[^\s]*$`)
	return urlRegex.MatchString(str)
}
