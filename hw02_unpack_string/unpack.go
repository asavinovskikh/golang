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
	var prev rune = 0
	var specs = false

	for ind, v := range mystring {
		fmt.Printf("%q ну че погнали %q .\n", ind, v) //убрать1
		if v == 92 {
			specs = !specs
			if specs == false {
				prev = v
				s = append(s, string(prev))
			}
			continue
		}

		if amount, err := strconv.Atoi(string(v)); err == nil {
			//fmt.Printf("%q looks like a number.\n", v) //убрать1
			if specs == true {
				//	fmt.Printf("%q но мы будем думать что это символ.\n", v) //убрать1
				prev = v
				specs = false
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
				if specs == true {
					fmt.Printf("%q specs\n", string(v)) //kill
					specs = false
					continue
				}
				err := errors.New("amount is not appropriate to text")
				fmt.Printf("%q message incorrect\n", v) //убрать2
				return "", err
			}

		} else {
			//s := "hello " + string('0' + 3)
			//s := "hello " + string('A' + 1)
			prev = v
			if specs == false {
				s = append(s, string(prev))
			}
			specs = false
			//fmt.Printf("%q буква \n", v)
		}
	}
	result := strings.Join(s, "")
	return result, nil
}
