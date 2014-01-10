package utils

import (
	"bytes"
	"compress/zlib"
	"errors"
	"io"
)

//数据以Zlib方式解压，参数为输入，返回为输出
func ZlibDecompress(data []byte) ([]byte, error) {

	r, err := zlib.NewReader(String(data).NewReader())
	if err != nil {
		return nil, err
	}
	defer r.Close()

	bs := NewByteString()
	buff := make([]byte, 4096)
	for {
		if n, err := r.Read(buff); err != nil {
			if err == io.EOF {
				return bs.GetBuff(), nil
			} else {
				return bs.GetBuff(), err
			}
		} else {
			bs.WriteBytes(buff, n)
		}
	}
}

//数据以zlib压缩，参数为输入，返回为输出
func ZlibCompress(data []byte) ([]byte, error) {
	var b bytes.Buffer

	w := zlib.NewWriter(&b)
	if w == nil {
		err := errors.New("zlib.NewWriter failed!")
		return b.Bytes(), err
	}
	defer w.Close()

	_, err := w.Write([]byte(data))
	return b.Bytes(), err

}
