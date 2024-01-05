package utils

import (
	"io"
	"os"
	"strconv"

	"github.com/fatih/color"
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

func StitchChunksIntoFile(folderName string, numberOfChunks int) {
	// create file
	// read chunks
	// write to file
	// delete chunks

	f, err := os.Create(folderName)

	if err != nil {
		color.Red("Error in creating file", folderName)
	}

	defer f.Close()

	for i := 0; i < numberOfChunks; i++ {
		chunkName := folderName + "/" + strconv.Itoa(i) + ".chunk"
		chunkFile, err := os.Open(chunkName)

		if err != nil {
			color.Red("Error in opening chunk", chunkName)
		}

		defer chunkFile.Close()

		_, err = io.Copy(f, chunkFile)

		if err != nil {
			color.Red("Error in writing chunk", chunkName)
		}
	}
}
