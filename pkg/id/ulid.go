package id

import (
	"github.com/oklog/ulid/v2"
)

func NewULID() string {
	return ulid.Make().String()
}

func CheckIsULID(input string) error {
	_, err := ulid.Parse(input)
	return err
}
