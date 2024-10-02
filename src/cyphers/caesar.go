package cyphers

import "strings"

const ALPHABET_LENGTH = 26

var (
	upperAlphabet = [ALPHABET_LENGTH]rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	lowerAlphabet = [ALPHABET_LENGTH]rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
)

type caesarString struct {
	originalString *string
	key            rune
	encodedString  string
}

func (s caesarString) Encoded() string {
	if len(s.encodedString) < len(*s.originalString) {
		s.encodedString = caesarCypher(*s.originalString, s.key)
	}
	return s.encodedString
}

func caesarCypher(message string, key rune) string {
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

func NewCaesarString(str string, key int) caesarString {
	cs := caesarString{
		originalString: &str,
		key:            rune(key),
	}
	return cs
}
