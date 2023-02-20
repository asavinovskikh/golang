package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")
var specSymbol int32 = 92

func Unpack(unpackingString string) (string, error) {
	s := make([]string, len(unpackingString))
	var prev rune = 0
	var specs = false

	for _, v := range unpackingString {
		if v == specSymbol {
			specs = !specs
			if specs == false {
				prev = v
				s = append(s, string(prev))
			}
			continue
		}

		//amount
		if amount, err := strconv.Atoi(string(v)); err == nil {
			if specs == true {
				prev = v
				s = append(s, string(prev))
				specs = false
				continue
			}

			if prev != 0 {
				s = s[:len(s)-1]

				for amount > 0 {
					s = append(s, string(prev))
					amount--
				}
				prev = 0
				continue
			} else {
				return "", ErrInvalidString
			}

		} else {
			//letter
			if specs == true {
				return "", ErrInvalidString
			}
			prev = v
			s = append(s, string(prev))
		}
	}
	result := strings.Join(s, "")
	return result, nil
}
