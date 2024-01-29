package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var n, m, temp int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		limitsTemp := []int{15, 30}
		fmt.Scan(&m)
		for j := 0; j < m; j++ {
			scanner.Scan()
			line := scanner.Text()
			temp, limitsTemp = comfortTemp(line, limitsTemp)
			fmt.Println(temp)

		}
		fmt.Println()
	}
}

func comfortTemp(str string, limitsTemp []int) (int, []int) {
	if limitsTemp[0] > limitsTemp[1] {
		return -1, limitsTemp
	}

	comfort, err := strconv.Atoi(string([]byte(str[3:])))
	if err != nil {
		fmt.Println(err)
		return -1, limitsTemp
	}

	if str[0] == '>' {
		limitsTemp[0] = comfort
	}
	if str[0] == '<' {
		limitsTemp[1] = comfort
	}
	if limitsTemp[0] == limitsTemp[1] {
		return limitsTemp[0], limitsTemp
	}
	if limitsTemp[0] > limitsTemp[1] {
		return -1, limitsTemp
	}
	return limitsTemp[1] - 1, limitsTemp

}

//     Пример теста 1
//     Входные данные

// 4
// 1
// >= 30
// 6
// >= 18
// <= 23
// >= 20
// <= 27
// <= 21
// >= 28
// 3
// <= 25
// >= 20
// >= 25
// 3
// <= 15
// >= 30
// <= 24
// Выходные данные
// 30

// 18
// 18
// 20
// 20
// 20
// -1

// 15
// 20
// 25

// 15
// -1
// -1
