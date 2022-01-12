/**
Напишите программу, которая в последовательности чисел находит сумму двузначных чисел, кратных 8.
Программа в первой строке получает на вход число n - количество чисел в последовательности,
во второй строке -- n чисел, входящих в данную последовательность.

Sample Input:
5
38 24 800 8 16

Sample Output:
40
*/

package main

import "fmt"

func main() {
	const krat int = 8
	var a, b int
	var sum int = 0

	fmt.Scan(&a)

	for i := 0; i < a; i++ {
		fmt.Scan(&b)
		if b%krat == 0 && b > 9 && b < 100 {
			sum += b
		}
	}

	fmt.Println(sum)
}
