package main

import (
	"fmt"
	"sort"
)

func AnagramSorter(words []string) [][]string {
	var result [][]string

	// return empty result if sliceOfStr is empty
	if len(words) <= 0 {
		return result
	}

	tempHashTable := make(map[int][]string)
	for _, str := range words {
		sumRunes := sumOfRune(str)

		temp, ok := tempHashTable[sumRunes]

		if ok {
			tempHashTable[sumRunes] = append(temp, str)
		} else {
			tempHashTable[sumRunes] = []string{str}
		}
	}

	result = make([][]string, 0)
	for _, val := range tempHashTable {
		result = append(result, val)
	}

	sort.Slice(result, func(i, j int) bool { return len(result[i]) > len(result[j]) })

	return result
}

func sumOfRune(str string) int {
	var sum int

	if len(str) <= 0 {
		return 0
	}

	for _, rune := range []rune(str) {
		sum += int(rune)
	}

	return sum
}

var words = []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}

func main() {
	result := AnagramSorter(words)
	fmt.Println(result)
}
