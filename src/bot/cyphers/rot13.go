package cyphers

import "strings"

const ALPHABET_LENGTH = 26

var (
	upperAlphabet = [ALPHABET_LENGTH]rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	lowerAlphabet = [ALPHABET_LENGTH]rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
)

type CaesarString struct {
	originalString *string
	key            rune
	encodedString  string
}

func (s CaesarString) Encoded() (cypheredString string) {
	if len(s.encodedString) < len(*s.originalString) {
		s.encodedString = rot13(*s.originalString, s.key)
	}
	return s.encodedString
}

func rot13(message string, key rune) string {
	cyphered := strings.Builder{}
	for _, char := range message {
		if char >= 'a' && char <= 'z' {
			char = lowerAlphabet[(char+key-'a')%ALPHABET_LENGTH]
		} else if char >= 'A' && char <= 'Z' {
			char = upperAlphabet[(char+key-'A')%ALPHABET_LENGTH]
		}
		cyphered.WriteRune(char)
	}
	return cyphered.String()
}

func NewCaesarString(str string) CaesarString {
	cs := CaesarString{
		originalString: &str,
		key:            rune(13),
	}
	return cs
}
