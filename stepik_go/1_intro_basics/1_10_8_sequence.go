/**
10 Циклы
https://stepik.org/lesson/228263/step/8?unit=200796

Напишите программу, которая считывает целые числа с консоли по одному числу в строке.

Для каждого введённого числа проверить:
- если число меньше 10, то пропускаем это число;
- если число больше 100, то прекращаем считывать числа;
- в остальных случаях вывести это число обратно на консоль в отдельной строке.

Sample Input:
30
11
7
101

Sample Output:
30
11
*/

package main

import "fmt"

func main() {
	const lim1 int = 10
	const lim2 int = 100
	var a int

	for fmt.Scan(&a); a <= lim2; fmt.Scan(&a) {
		if a < lim1 {
			continue
		}
		fmt.Println(a)
	}
}
