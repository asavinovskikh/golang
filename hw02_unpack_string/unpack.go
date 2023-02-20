package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var (
	ErrInvalidString       = errors.New("invalid string")
	specSymbol       int32 = 92
)

func Unpack(unpackingString string) (string, error) {
	var s []string
	var prev rune
	var specs bool

	for _, v := range unpackingString {
		if v == specSymbol {
			specs = !specs
			if !specs {
				prev = v
				s = append(s, string(prev))
			}
			continue
		}

		// amount
		if amount, err := strconv.Atoi(string(v)); err == nil {
			if specs {
				prev = v
				s = append(s, string(prev))
				specs = false
				continue
			}

			if prev == 0 {
				return "", ErrInvalidString
			}
			s = s[:len(s)-1]

			for amount > 0 {
				s = append(s, string(prev))
				amount--
			}
			prev = 0
		} else {
			// letter
			if specs {
				return "", ErrInvalidString
			}
			prev = v
			s = append(s, string(prev))
		}
	}
	result := strings.Join(s, "")
	return result, nil
}
