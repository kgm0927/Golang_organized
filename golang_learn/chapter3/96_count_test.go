package chapter3

import (
	"fmt"
	"sort"
)

func Example_count() {
	codeCount := map[rune]int{}
	Count("가나다나", codeCount)
	var keys sort.IntSlice

	for key := range codeCount {
		keys = append(keys, int(key))
	}

	sort.Sort(keys)

	for _, key := range keys {
		fmt.Println(rune(key), codeCount[rune(key)])
	}
}
