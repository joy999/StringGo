package utils

import (
	"os"
)

func FilePutContents(filename string, c []byte) error {
	f, err := os.Create(filename)
	if os.IsExist(err) {
		os.Remove(filename)
		f, err = os.Create(filename)
	}

	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.Write(c)

	return err
}
