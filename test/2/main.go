package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		date := make([]int, 3)
		for j := 0; j < 3; j++ {
			fmt.Scan(&m)
			date[j] = m
		}
		fmt.Println(isValidDate(date))
	}

}

func isValidDate(date []int) string {

	if date[2]%4 == 0 {
		if date[2]%100 != 0 {
			if date[1] == 2 {
				if date[0] > 29 {
					return "NO"
				}
				return "YES"
			}

		}
	}
	if date[2]%400 == 0 {
		if date[1] == 2 {
			if date[0] > 29 {
				return "NO"
			}
			return "YES"
		}
	}

	if date[1] == 2 {
		if date[0] > 28 {
			return "NO"
		}
	}

	var isMinDay = map[int]struct{}{2: struct{}{}, 4: struct{}{}, 6: struct{}{}, 9: struct{}{}, 11: struct{}{}}
	if date[0] > 30 {
		_, ok := isMinDay[date[1]]
		if ok {
			return "NO"
		}
	}

	return "YES"
}
