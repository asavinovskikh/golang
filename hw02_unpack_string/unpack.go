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
	var builder strings.Builder
	var prev rune
	var specs bool

	for _, v := range unpackingString {
		if v == specSymbol {
			specs = !specs
			if !specs {
				prev = v
				builder.WriteString(string(prev))
			}
			continue
		}

		// amount
		if amount, err := strconv.Atoi(string(v)); err == nil {
			if specs {
				prev = v
				builder.WriteString(string(prev))
				specs = false
				continue
			}

			if prev == 0 {
				return "", ErrInvalidString
			}

			for amount > 0 {
				builder.WriteString(string(prev))
				amount--
			}
			prev = 0
		} else {
			// letter
			if specs {
				return "", ErrInvalidString
			}
			prev = v
			builder.WriteString(string(prev))
		}
	}
	result := builder.String()
	return result, nil
}
