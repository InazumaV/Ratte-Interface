package errors

import "encoding/gob"

func init() {
	gob.RegisterName("*error.StringError", new(StringError))
}

type Error interface {
	Raw() any
	Error() string
}

type StringError string

func NewString(text string) Error {
	return (*StringError)(&text)
}

func NewStringFromErr(err error) Error {
	if err == nil {
		return nil
	}
	s := err.Error()
	return (*StringError)(&s)
}

func (e StringError) Error() string {
	return string(e)
}

func (e StringError) Raw() any {
	return string(e)
}

type BytesError []byte

func NewBytes(bytes []byte) Error {
	return (*BytesError)(&bytes)
}

func (b BytesError) Error() string {
	return string(b)
}

func (b BytesError) Raw() any {
	return []byte(b)
}
