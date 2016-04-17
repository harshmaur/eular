package problems

import (
	"fmt"
	"os"
	"time"
)

// EvenlyDivisible finds the smallest number evenly divisible from 1-20
func EvenlyDivisible() {
	start := time.Now()
	check := 20
	for {
		for i := 1; i <= 20; i++ {
			if check%i != 0 {
				break
			}
			if i == 20 {
				fmt.Println(check)
				elapsed := time.Since(start)
				fmt.Printf("Binomial took %s", elapsed)
				os.Exit(0)
			}
			continue
		}
		check = check + 20
	}

}
