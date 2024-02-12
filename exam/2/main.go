package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var t, n int
	var p, a, result float64
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n)
		fmt.Fscan(in, &p)
		curPerc := p / 100
		curRes := 0.00
		for j := 0; j < n; j++ {
			fmt.Scan(&a)
			curRes += curPerc * a
		}
		if i == t {
			break
		}
		fmt.Printf("%.2f", curRes)
		result += curRes
	}
	fmt.Printf("%.2f", result)

}
