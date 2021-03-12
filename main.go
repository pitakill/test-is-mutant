package main

import (
	"fmt"
	"strings"
)

func main() {
	input := []string{
		"TCGTA",
		"AGCTG",
		"ATGAT",
		"ATCGT",
		"AAAGT",
	}

	fmt.Println(isMutant(strings.Join(input, ""), len(input)))
}

func isMutant(dna string, k int) bool {
	for l := 0; l < k; l++ {
		for i, j := l*k, (l+1)*(k)-1; i < j; i, j = i+1, j-1 {
			chunk := dna[i:j]
			if strings.Index(chunk, "AA") != strings.LastIndex(chunk, "AA") {
				return true
			}
		}
	}

	return false
}
