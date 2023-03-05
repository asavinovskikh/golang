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
	s := make([]int, 0, 0)
	//freqResult := make([]string, 0, 0)

	if text == "" {
		return nil
	}

	spliterate := strings.FieldsFunc(text, Split)

	for _, v := range spliterate { //strings.Split(text, " ") {
		//fmt.Println("что-то нет то " + v)

		k := stats[v]
		stats[v] = k + 1

		//fmt.Println(v, k+1)
	}

	max := 0
	word := ""
	//fmt.Println(stats)
	for jj, m := range stats {
		if m > max {
			word = jj
			max = m
		}
		//fmt.Println("stats", jj, m)
		s = append(s, m)
	}

	//sort.Sort(find10max)//чет не работает!
	slicesort(s)

	s = s[:10]
	fmt.Println("MAXIMUM", max, word)
	fmt.Println("s", s)

	//var reverse map[int][]string //частота слова
	reverse := make(map[int][]string)

	for i, v := range stats {
		k := reverse[v]
		k = append(k, i)
		reverse[v] = k
	}

	freeee := make([]string, 0, 0)

	jjkj := 0

	for i, v := range reverse {
		if s[jjkj] == i {
			for kkk := range v {
				freeee = append(freeee, v[kkk])
			}
			jjkj += 1
		}
	}

	sort.Reverse(sort.IntSlice(varietyamount))

	fmt.Println("freeee")
	for q := range freeee {
		fmt.Println(" ", freeee[q])
	}
	//fmt.Println(freeee)
	return freeee

	//	ii := 10
	//for ind := range reverse {
	//	if ii > 0 {
	//		k := reverse[ind]
	//		for kk := range k {
	//			freqResult = append(freqResult, k[kk])
	//			ii--
	//		}
	//	} else {
	//		return freqResult[:10]
	//	}
	//}

	//return nil
}

func slice2sort(arr [][]int) {

	sort.Slice(arr, func(i, j int) bool {
		if arr[i][1] > arr[j][1] {
			return true
		}
		if arr[i][1] == arr[j][1] && arr[i][0] < arr[j][0] {
			return true
		}
		return false
	})
}

func slicesort(arr []int) {

	sort.Slice(arr, func(i, j int) bool {
		if arr[i] > arr[j] {
			return true
		}
		if arr[i] == arr[j] && arr[i] < arr[j] {
			return true
		}
		return false
	})
}
