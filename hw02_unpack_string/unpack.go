package hw02unpackstring

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(_ string) (string, error) {
	return "", nil
}

func main(mystring string) (string, error) {
	s := make([]string, len(mystring))
	var prev string = ""

	for _, v := range mystring {

		if amount, err := strconv.Atoi(string(v)); err == nil {
			//fmt.Printf("%q looks like a number.\n", v)
			if prev != "" {
				s = s[:len(s)-1]
				for amount > 0 {
					s = append(s, prev)
					amount--
				}
				prev = ""
			} else {
				err := errors.New("amount is not appropriate to text")
				fmt.Printf("%q message incorrect\n", v) //убрать
				return "", err
			}

		} else {
			prev = string(v)
			s = append(s, prev)
			//fmt.Printf("%q буква \n", v)
		}
	}
	result := strings.Join(s, "")
	return result, nil
}
