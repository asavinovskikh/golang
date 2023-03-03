package hw03frequencyanalysis

import (
	"fmt"
	"sort"
	"strings"
)

func Split(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n' || r == ',' || r == '.' || r == '!'
}

func reverseInts(input []int) []int {
	if len(input) == 0 {
		return input
	}
	return append(reverseInts(input[1:]), input[0])
}

func Top10(text string) []string {
	//var stats map[string]int //частота слова
	stats := make(map[string]int)
	varietyamount := make([]int, 10, 10)
	freqResult := make([]string, 0, 0)

	if text == "" {
		return nil
	}

	for _, v := range strings.FieldsFunc(text, Split) { //strings.Split(text, " ") {
		//fmt.Println(i, v)
		k := stats[v]
		stats[v] = k + 1

		//fmt.Println(i, v)
	}

	//fmt.Println(stats)

	//var reverse map[int][]string //частота слова
	reverse := make(map[int][]string)

	for i, v := range stats {
		varietyamount = append(varietyamount, v)

		k := reverse[v]
		k = append(k, i)
		reverse[v] = k
	}

	sort.Sort(sort.Reverse(sort.IntSlice(varietyamount)))
	//sort.Ints(varietyamount)

	ii := 10

	//fmt.Println(varietyamount)

	for ind := range varietyamount {
		//fmt.Println(ii)
		if ii > 0 {
			//fmt.Println(ind)
			k := reverse[ind]
			//fmt.Println(k)
			fmt.Println(k)
			for kk := range k {
				freqResult = append(freqResult, k[kk])
				ii--
			}

			//freqResult = append(k)
			//ii--
		} else {
			return freqResult
		}

	}

	return nil
}
