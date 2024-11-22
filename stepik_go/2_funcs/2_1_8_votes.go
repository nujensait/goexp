/**
Напишите "функцию голосования", возвращающую то значение (0 или 1),
которое среди значений ее аргументов x, y, z встречается чаще.

Входные данные
Вводится 3 числа - x, y, z (x, y, z равны 0 или 1).

Выходные данные
Необходимо вернуть значение функции от x, y, z.

Sample Input:
0 0 1

Sample Output:
0
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var vote = vote(0, 1, 1)
	fmt.Println(vote)
}

func vote(x, y, z int) int {
	sum := x + y + z
	vote := int(math.Round(float64(sum) / 3))
	return vote
}
