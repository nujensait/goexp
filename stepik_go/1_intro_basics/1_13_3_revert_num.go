/**
Дано трехзначное число. Переверните его, а затем выведите.

Формат входных данных
На вход дается трехзначное число, не оканчивающееся на ноль.

Формат выходных данных
Выведите перевернутое число.

Sample Input:
843

Sample Output:
348
*/

package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var a string
	var res, i int

	fmt.Scan(&a)

	res = 0

	for i = 0; i < len(a); i++ {
		val, err := strconv.Atoi(string(a[i]))
		if err == nil {
			res += val * int(math.Pow(10, float64(i)))
		}
	}

	fmt.Print(res)
}
