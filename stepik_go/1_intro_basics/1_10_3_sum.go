/**
1.10 Циклы
https://stepik.org/lesson/228263/step/3?unit=200796

Требуется написать программу, при выполнении которой с клавиатуры считываются два натуральных числа A и B (каждое не более 100, A < B).
Вывести сумму всех чисел от A до B  включительно.

Sample Input:
1 5

Sample Output:
15
*/

package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)

	sum := 0
	for i := a; i <= b; i++ {
		sum += i
	}

	fmt.Println(sum)
}
