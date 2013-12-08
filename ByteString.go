package utils

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
)

type ByteString struct {
	buff *bytes.Buffer
}

func NewByteString() *ByteString {
	bs := new(ByteString)
	bs.buff = bytes.NewBuffer(make([]byte, 0))

	return bs
}

func (this *ByteString) ClearBuff() {
	this.buff.Reset()
}

/**
  将字节流追加到当前处理之后
*/
func (this *ByteString) SetBuff(buff []byte) {
	this.buff.Write(buff)
}

func (this *ByteString) ReadFromReader(reader io.Reader) (int64, error) {
	return this.buff.ReadFrom(reader)
}

/*
	获取剩余未读的字节流
*/
func (this *ByteString) GetBuff() []byte {
	return this.buff.Bytes()
}

func (this *ByteString) WriteToWriter(w io.Writer) (int64, error) {
	return this.buff.WriteTo(w)
}

func (this *ByteString) Read(args ...interface{}) error {
	var err error
	for _, a := range args {
		switch v := a.(type) {
		case *int:
			var i int32
			i, err = this.ReadInt32()
			if err == nil {
				*v = int(i)
			}
		case *uint:
			var i int32
			i, err = this.ReadInt32()
			if err == nil {
				*v = uint(i)
			}
		case *int64:
			*v, err = this.ReadInt64()
		case *int32:
			*v, err = this.ReadInt32()
		case *int16:
			*v, err = this.ReadInt16()
		case *int8:
			*v, err = this.ReadInt8()
		case *uint64:
			*v, err = this.ReadUInt64()
		case *uint32:
			*v, err = this.ReadUInt32()
		case *uint16:
			*v, err = this.ReadUInt16()
		case *uint8:
			*v, err = this.ReadUInt8()

		case *string:
			*v, err = this.ReadVarString()
		case []byte:
			v, err = this.ReadBytes(len(v))
		default:
			err = errors.New("Not Support Data Type")
		}
		if err != nil {
			return err
		}
	}

	return nil

}

func (this *ByteString) Write(args ...interface{}) error {
	var err error
	for _, a := range args {
		switch v := a.(type) {
		case int64:
			err = this.WriteInt64(v)
		case int32:
			err = this.WriteInt32(v)
		case int:
			var i int = v
			err = this.WriteInt32(int32(i))
		case int16:
			err = this.WriteInt16(v)
		case int8:
			err = this.WriteInt8(v)
		case uint64:
			err = this.WriteUInt64(v)
		case uint32:
			err = this.WriteUInt32(v)
		case uint:
			var i uint = v
			err = this.WriteUInt32(uint32(i))
		case uint16:
			err = this.WriteUInt16(v)
		case uint8:
			err = this.WriteUInt8(v)
		case string:
			err = this.WriteVarString(v)
		case []byte:
			err = this.WriteBytes(v, len(v))
		default:
			err = errors.New("Not Support Data Type")
		}
		if err != nil {
			return err
		}
	}

	return nil

}

///读的方法完成

func (this *ByteString) ReadInt8() (int8, error) {
	var n int8
	if err := binary.Read(this.buff, binary.BigEndian, &n); err != nil {
		return 0, err
	}
	return n, nil
}

func (this *ByteString) ReadUInt8() (uint8, error) {
	var n uint8
	if err := binary.Read(this.buff, binary.BigEndian, &n); err != nil {
		return 0, err
	}
	return n, nil
}

func (this *ByteString) ReadInt16() (int16, error) {
	var n int16
	if err := binary.Read(this.buff, binary.BigEndian, &n); err != nil {
		return 0, err
	}
	return n, nil
}

func (this *ByteString) ReadUInt16() (uint16, error) {

	var n uint16
	if err := binary.Read(this.buff, binary.BigEndian, &n); err != nil {
		return 0, err
	}
	return n, nil
}

func (this *ByteString) ReadInt32() (int32, error) {

	var n int32
	if err := binary.Read(this.buff, binary.BigEndian, &n); err != nil {
		return 0, err
	}
	return n, nil
}

func (this *ByteString) ReadUInt32() (uint32, error) {
	var n uint32
	if err := binary.Read(this.buff, binary.BigEndian, &n); err != nil {
		return 0, err
	}
	return n, nil
}
func (this *ByteString) ReadInt64() (int64, error) {
	var n int64
	if err := binary.Read(this.buff, binary.BigEndian, &n); err != nil {
		return 0, err
	}
	return n, nil
}

func (this *ByteString) ReadUInt64() (uint64, error) {
	var n uint64
	if err := binary.Read(this.buff, binary.BigEndian, &n); err != nil {
		return 0, err
	}
	return n, nil
}

func (this *ByteString) ReadString(size int) (string, error) {
	bs, err := this.ReadBytes(size)
	return string(bs), err
}

func (this *ByteString) ReadBytes(size int) ([]byte, error) {
	bs := make([]byte, size)
	if _, err := this.buff.Read(bs); err != nil {
		return bs, err
	} else {
		return bs, nil
	}
}

func (this *ByteString) ReadVarString() (string, error) {
	size, err := this.ReadUInt16()
	if err != nil {
		return "", err
	}
	return this.ReadString(int(size))
}

///写处理
func (this *ByteString) WriteInt8(n int8) error {
	return binary.Write(this.buff, binary.BigEndian, n)
}

func (this *ByteString) WriteUInt8(n uint8) error {
	return binary.Write(this.buff, binary.BigEndian, n)
}

func (this *ByteString) WriteInt16(n int16) error {
	return binary.Write(this.buff, binary.BigEndian, n)
}

func (this *ByteString) WriteUInt16(n uint16) error {
	return binary.Write(this.buff, binary.BigEndian, n)
}

func (this *ByteString) WriteInt32(n int32) error {
	return binary.Write(this.buff, binary.BigEndian, n)
}

func (this *ByteString) WriteUInt32(n uint32) error {
	return binary.Write(this.buff, binary.BigEndian, n)
}

func (this *ByteString) WriteInt64(n int64) error {
	return binary.Write(this.buff, binary.BigEndian, n)
}

func (this *ByteString) WriteUInt64(n uint64) error {
	return binary.Write(this.buff, binary.BigEndian, n)
}

func (this *ByteString) WriteString(str string) error {
	return binary.Write(this.buff, binary.BigEndian, str)
}

func (this *ByteString) WriteVarString(str string) error {
	size := len(str)
	if err := this.WriteUInt16(uint16(size)); err != nil {
		return err
	}
	return this.WriteString(str)
}

func (this *ByteString) WriteBytes(bs []byte, size int) error {
	_bs := bs[:size]
	return binary.Write(this.buff, binary.BigEndian, _bs)
}
