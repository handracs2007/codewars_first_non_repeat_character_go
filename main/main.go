package main

import (
	"fmt"
	"strings"
)

type StrMeta struct {
	ch    rune
	count uint
}

func FirstNonRepeating(str string) string {
	var metadata = make([]*StrMeta, 0)
	var charMap = make(map[rune]*StrMeta)

	for _, ch := range str {
		tempCh := rune(strings.ToLower(string(ch))[0])
		meta, ok := charMap[tempCh]

		if ok {
			meta.count += 1
			charMap[tempCh] = meta
		} else {
			var strMeta = StrMeta{
				ch:    ch,
				count: 1,
			}

			charMap[tempCh] = &strMeta
			metadata = append(metadata, &strMeta)
		}
	}

	var least = ""

	for _, meta := range metadata {
		if meta.count == 1 {
			least = string(meta.ch)
			break
		}
	}

	return least
}

func main() {
	fmt.Println(FirstNonRepeating("stress"))
	fmt.Println(FirstNonRepeating("sTreSS"))
}
