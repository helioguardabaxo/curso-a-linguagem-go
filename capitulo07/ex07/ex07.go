package main

import (
	"fmt"
)

func main() {
	for i := 10; i <= 100; i++ {
		if i%4 == 0 {
			fmt.Printf("%d / 4: %d\n", i, (i % 4))
		} else if i%3 == 0 {
			fmt.Printf("%d / 3: %d\n", i, (i % 3))
		} else {
			fmt.Printf("%d não é divisível por 4 nem por 3\n", i)
		}
	}
}
