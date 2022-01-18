package tracker

import (
	"crypto/rand"
	"math/big"
	"strings"
)

type UID string

var (
	letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	chars   = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
)

func NewUID() (UID, error) {
	var sb strings.Builder
	r, err := randomLetter()
	if err != nil {
		return "", err
	}
	sb.WriteRune(r)
	for i := 0; i < 10; i++ {
		r, err := randomChar()
		if err != nil {
			return "", err
		}
		sb.WriteRune(r)
	}
	return UID(sb.String()), nil
}

func randomLetter() (rune, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
	if err != nil {
		return 0, err
	}
	return letters[n.Int64()], nil
}

func randomChar() (rune, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
	if err != nil {
		return 0, err
	}
	return chars[n.Int64()], nil
}
