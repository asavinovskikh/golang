package hw03frequencyanalysis

import (
	"fmt"
	"strings"
)

func Top10(text string) []string {
	var stats map[string]int

	for i, v := range strings.Split(text, " ") {
		k := stats[v]
		stats[v] = k + 1
		fmt.Println(i, v)
	}

	a := make([]string, 10, 10)
	var frequencies map[string]int

	for i := 0; i < 10; i++ {
		frequencies[stats[i]]
	}

	return nil
}
