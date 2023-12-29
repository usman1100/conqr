package utils

import "strings"

func GetFileNameFromUrl(fileUrl string) string {
	splits := strings.Split(fileUrl, "/")
	fileName := splits[len(splits)-1]
	return fileName
}
