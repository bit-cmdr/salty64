package salty64

import (
	b64 "encoding/base64"
	"errors"
)

var (
	// ErrOffsetNegative denotes a negative offset was provided
	ErrOffsetNegative = errors.New("Offset is negative")

	// ErrOffsetInvalid denotes an offset was provided that was greater
	// than the length of the salt
	ErrOffsetInvalid = errors.New("Offset is negative")
)

// Shaker is used to salt the string and then only use a fragment of it
type Shaker struct {
	Salt   string
	Offset int
}

// NewShaker returns a fully populated Shaker or an error
func NewShaker(salt string, offset int) (Shaker, error) {
	s := Shaker{salt, offset}
	err := validate(s)
	if err != nil {
		return Shaker{}, err
	}

	return s, nil
}

// Encode encodes a string to a base64 string by double encoding it with
// a prefix of a double encoded salt fragment
func (salt Shaker) Encode(s string) (string, error) {
	err := validate(salt)
	if err != nil {
		return "", err
	}

	return string(b64.URLEncoding.EncodeToString([]byte(salt.salty() + b64.URLEncoding.EncodeToString([]byte(s))))), nil
}

// Decode decodes a string from a base64 string by double decoding it with
// a prefix of a double decoded salt fragment
func (salt Shaker) Decode(s string) (string, error) {
	err := validate(salt)
	if err != nil {
		return "", err
	}

	un, err := b64.URLEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}

	dec, err := b64.URLEncoding.DecodeString(string(un[len(salt.salty()):]))
	if err != nil {
		return "", err
	}

	return string(dec), nil
}

func validate(s Shaker) error {
	switch {
	case s.Offset > len(s.Salt):
		return ErrOffsetInvalid
	case s.Offset < 0:
		return ErrOffsetNegative
	}
	return nil
}

func (salt Shaker) salty() string {
	enc := b64.URLEncoding.EncodeToString([]byte(b64.URLEncoding.EncodeToString([]byte(salt.Salt))))
	l := len(enc)
	return enc[:l-salt.Offset]
}
