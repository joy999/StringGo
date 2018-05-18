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

func FileAppendContents(filename string, c []byte) error {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644 )
	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.Write(c)

	return err
}

func FileGetContents(filename string) ([]byte, error) {

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	bs := NewByteString()
	buff := make([]byte, 4096)

	for {
		if n, err := f.Read(buff); err != nil && err != io.EOF {
			return bs.GetBuff(), err
		} else {
			if n > 0 {
				bs.WriteBytes(buff, n)
			}
			if n < 4096 {
				return bs.GetBuff(), nil
			}
		}
	}
}
