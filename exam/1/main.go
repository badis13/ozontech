package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var mtx = map[int]int{1: 4, 2: 3, 3: 2, 4: 1}
		for j := 0; j < 10; j++ {
			fmt.Scan(&m)
			if mtx[m] > 1 {
				mtx[m]--
				continue
			}
			delete(mtx, m)
		}
		if len(mtx) == 0 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
