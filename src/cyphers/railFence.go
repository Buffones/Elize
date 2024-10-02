package cyphers

import "strings"

type railFenceString struct {
	originalString *string
	rails          int
	encodedString  string
}

func (r railFenceString) Encoded() string {
	if r.encodedString == "" && *r.originalString != "" {
		r.encodedString = railFenceCypher(*r.originalString, r.rails)
	}
	return r.encodedString
}

func railFenceCypher(message string, rails int) string {
	builders := make([]strings.Builder, rails)
	direction := 1
	builders[0].WriteByte(message[0])
	currentRail := 1
	for i := 1; i < len(message); i++ {
		char := message[i]
		if char != ' ' && char != '.' && char != '\n' {
			builders[currentRail].WriteByte(char)
			if currentRail == rails-1 || currentRail == 0 {
				direction = -direction
			}
			currentRail += direction
		}
	}
	for i := 1; i < len(builders); i++ {
		builders[0].WriteRune(' ')
		builders[0].WriteString(builders[i].String())
	}
	return builders[0].String()
}

func NewRailFenceString(str string, rails int) railFenceString {
	return railFenceString{
		originalString: &str,
		rails:          rails,
	}
}
