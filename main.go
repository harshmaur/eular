package main

import (
	"fmt"
	"time"

	"github.com/harshmaur/eular/problems"
)

func main() {
	start := time.Now()
	problems.EvenlyDivisible()
	elapsed := time.Since(start)
	fmt.Printf("Binomial took %s", elapsed)
}
