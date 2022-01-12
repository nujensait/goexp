/**
https://stepik.org/lesson/228261/step/14?unit=200794

Дано натуральное число, выведите его последнюю цифру.

Формат входных данных
На вход дается натуральное число N, не превосходящее 10000.

Формат выходных данных
Выведите одно целое число - ответ на задачу.

Sample Input:

123
Sample Output:

3
*/

package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	var a, b, c string
	a = string(s[0])
	b = string(s[1])
	c = string(s[2])

	if a != b && a != c && b != c {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
