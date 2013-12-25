package utils

import (
	"crypto"
	"github.com/joy999/mahonia"
	"io"
	"log"
)

func GBKToUTF8(inputStr string) string {
	return EncodingConvert(inputStr, "gbk", "utf-8")
}

func UTF8ToGBK(inputStr string) string {
	return EncodingConvert(inputStr, "utf-8", "gbk")
}

func trace(v ...interface{}) {
	log.Print(v...)
}

func EncodingConvert(inputStr string, fromCode string, toCode string) string {
	outStr := String("")

	outPointer := &outStr
	reader := String(inputStr).NewReader()
	var r io.Reader = reader
	var w io.Writer = outPointer

	if fromCode != "utf-8" {
		decode := mahonia.NewDecoder(fromCode)
		if decode == nil {
			log.Fatalf("Could not create decoder for %s", fromCode)

		}
		r = decode.NewReader(r)
	}

	if toCode != "utf-8" {
		encode := mahonia.NewEncoder(toCode)
		if encode == nil {
			log.Fatalf("Could not create encoder for %s", toCode)
		}
		w = encode.NewWriter(w)
	}
	io.Copy(w, r)

	return string(outStr)
}

func Md5(s string) string {
	h := crypto.MD5.New()
	io.WriteString(h, s)
	return string(h.Sum(nil))
}
