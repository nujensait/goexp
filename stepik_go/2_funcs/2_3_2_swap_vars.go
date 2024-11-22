/**
Поменяйте местами значения переменных на которые ссылаются указатели. После этого переменные нужно вывести.
ВАЖНО: Считайте что пакет main уже объявлен, а также функция main() уже есть.

Sample Input:
2 4

Sample Output:
4 2
*/

package main

import "fmt"

func swap(x1 *int, x2 *int) {
	tmp := *x1
	*x1 = *x2
	*x2 = tmp
}

func main() {
	a, b := 5, 7
	swap(&a, &b) // меняем местами
	fmt.Println(a, b)
}
