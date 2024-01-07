package files

import (
	"os"
	"strconv"
	"strings"
)

func GetFileNameFromUrl(fileUrl string) string {
	splits := strings.Split(fileUrl, "/")
	fileName := splits[len(splits)-1]
	return fileName
}

func FileAlreadExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func GenerateFileName(fileName string) string {

	tries := 0
	newFileName := fileName

	for FileAlreadExists(newFileName) {
		tries++
		newFileName = fileName + "(" + strconv.Itoa(tries) + ")"
	}

	return newFileName
}
