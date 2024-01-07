package files

import (
	"io"
	"os"
	"strconv"

	"github.com/fatih/color"
)

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
	outputFile, err := os.Create(folderName + "/" + folderName)

	if err != nil {
		color.Red("error in open output file", err.Error())
		return
	}

	defer outputFile.Close()

	for i := 0; i < numberOfChunks; i++ {
		chunkName := folderName + "/" + strconv.Itoa(i) + ".chunk"
		chunkFile, err := os.Open(chunkName)

		if err != nil {
			color.Red("error in open chunk file", err.Error())
			return
		}

		defer chunkFile.Close()

		_, err = io.Copy(outputFile, chunkFile)

		if err != nil {
			color.Red(err.Error())
			return
		}
		os.Remove(chunkName)
	}

	// check if a chunk with name [numberOfChunks + 1].chunk exists
	finalChunkName := folderName + "/" + strconv.Itoa(numberOfChunks) + ".chunk"
	finalChunkFile, err := os.Open(finalChunkName)

	if err != nil {
		return
	}

	defer finalChunkFile.Close()

	_, err = io.Copy(outputFile, finalChunkFile)

	if err != nil {
		return
	}

	os.Remove(finalChunkName)
}
