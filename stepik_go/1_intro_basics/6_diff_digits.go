/**
По данному трехзначному числу определите, все ли его цифры различны.

Формат входных данных
На вход подается одно натуральное трехзначное число.

Формат выходных данных
Выведите "YES", если все цифры числа различны, в противном случае - "NO".

Sample Input 1:
237

Sample Output 1:
YES

Sample Input 2:
117

Sample Output 2:
NO
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
