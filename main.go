package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const loop int = 10
	var chars = []string{"C", "S", "A", "P", "Ш", "Z", "A", "X", "F", "H", "B", "L", "0"}

outerloop:
	for i := 0; i <= loop; i++ {
		fmt.Printf("Nilai i = %d\n", i)

		if i == 4 {
			for j := 0; j <= loop; j++ {
				if j == 5 {
					continue
				}

				fmt.Printf("Nilai j = %d\n", j)

				if j == 4 {
					for k, v := range chars {
						if (k % 2) == 0 {
							unicode, _ := utf8.DecodeRuneInString(v)
							fmt.Printf("character %#U starts at byte position %d\n", unicode, k)
						}
					}

				}
			}
			break outerloop
		}

	}
}
