package utils

import (
	"io"
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

func FileGetContents(filename string) ([]byte, error) {
	content := make([]byte, 0)
	f, err := os.Open(filename)
	if err != nil {
		return content, err
	}

	defer f.Close()
	buff := make([]byte, 4096)

	for {
		if n, err := f.Read(buff); err != nil && err != io.EOF {
			return content, err
		} else {
			if n > 0 {
				content = append(content, buff[:n]...)
			}
			if n < 4096 {
				return content, nil
			}
		}
	}
}
