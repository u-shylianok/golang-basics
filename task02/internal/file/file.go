package file

import (
	"os"

	"github.com/sirupsen/logrus"
)

// WriteBytes func use os.Create and truncate existing file.
func WriteBytes(path string, bytes []byte) (int, error) {
	file, err := os.Create(path)
	if err != nil {
		logrus.Error(err)
		return 0, err
	}
	defer file.Close()

	n, err := file.Write(bytes)
	if err != nil {
		logrus.Error(err)
		return 0, err
	}

	return n, nil
}

// ReadBytes func use os.ReadFile.
func ReadBytes(path string) ([]byte, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return bytes, nil
}
