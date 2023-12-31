package utils

import (
	"io"
	"os"
	"strconv"
)

func FileAlreadExists(fileName string) bool {
	_, err := os.Stat(fileName)

	return err != nil
}

func GenerateFileName(fileName string) string {

	tries := 0
	newFileName := fileName

	for FileAlreadExists(fileName) {
		tries++
		newFileName = fileName + "(" + strconv.Itoa(tries) + ")"
	}

	return newFileName
}

func WriteDataToFile(reader *io.ReadCloser, fileName string) error {
	file, err := os.Create(fileName)

	if err != nil {
		return err
	}

	defer file.Close()

	_, err = io.Copy(file, *reader)

	if err != nil {
		return err
	}

	return nil
}
