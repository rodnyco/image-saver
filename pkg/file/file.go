package file

import (
	"os"
	"time"
)

func CreateFileName() string {
	return time.Now().Format("20060102150405")
}

func SaveFile(directory string, fileName string, fileData []byte) (string, error) {
	path := directory + string('/') + fileName
	file, err := os.Create(path)
	defer file.Close()
	_, err = file.Write(fileData)
	if err != nil {
		return "", err
	}

	return path, nil
}
