/**
Найдите первое число от 1 до n включительно, кратное c, но НЕ кратное d.

Входные данные

Вводится 3 натуральных числа n, c, d, каждое из которых не превышает 10000.

Выходные данные
Вывести первое число от 1 до n включительно, кратное c, но НЕ кратное d. Если такого числа нет - выводить ничего не нужно.

Sample Input:
20
3
5

Sample Output:
3

@hint Should use break/continue statements here
*/

package main

import "fmt"

func main() {
	var a, c, d int

	fmt.Scan(&a)
	fmt.Scan(&c)
	fmt.Scan(&d)

	for i := 1; i <= a; i++ {
		if i%c == 0 && i%d != 0 {
			fmt.Println(i)
			break
		}
	}
}
