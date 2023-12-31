package utils

import (
	"regexp"
	"strings"
)

func GetFileNameFromUrl(fileUrl string) string {
	splits := strings.Split(fileUrl, "/")
	fileName := splits[len(splits)-1]
	return fileName
}

func IsStringUrl(str string) bool {
	urlRegex := regexp.MustCompile(`^(https?|ftp):\/\/[^\s\/$.?#].[^\s]*$|^(localhost|127\.0\.0\.1):[0-9]+[^\s]*$`)
	return urlRegex.MatchString(str)
}
