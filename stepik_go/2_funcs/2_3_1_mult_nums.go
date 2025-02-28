/**
Напишите функцию, которая умножает значения на которые ссылаются два указателя
и выводит получившееся произведение в консоль.

Sample Input:
2 2

Sample Output:
4

Exec:
go run 2_3_1_mult_nums.go
*/

package main

import "fmt"

// считайте что fmt уже импортирован и main объявлен
func mult(x1 *int, x2 *int) {
	fmt.Println(*x1 * *x2)
}

func main() {
	a, b := 5, 7
	mult(&a, &b) // Передаем адреса переменных a и b; Программа выведет 35 - результат умножения 5 и 7.
}
