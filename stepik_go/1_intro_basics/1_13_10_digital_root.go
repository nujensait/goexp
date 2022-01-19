/**
Цифровой корень
https://stepik.org/lesson/229320/step/10?unit=201906

Цифровой корень натурального числа — это цифра, полученная в результате итеративного процесса суммирования цифр,
на каждой итерации которого для подсчета суммы цифр берут результат, полученный на предыдущей итерации.
Этот процесс повторяется до тех пор, пока не будет получена одна цифра.

Например цифровой корень 65536 это 7 , потому что 6+5+5+3+6=25 и 2+5=7 .
По данному числу определите его цифровой корень.

Входные данные
Вводится одно натуральное число n, не превышающее 10^710
7

Выходные данные
Вывести цифровой корень числа n.

Sample Input:
3456

Sample Output:
9
*/

package main

import (
	"fmt"
	"strconv"
)

/**
Return digital root of number
*/
func digital_root(n int) int {
	var res, i int
	res = 0
	var a string

	a = strconv.Itoa(n)

	for i = 0; i < len(a); i++ {
		val, err := strconv.Atoi(string(a[i]))
		if err == nil {
			res += val
		}
	}

	if res > 9 {
		return digital_root(res)
	}

	return res
}

func main() {
	var a int
	var res int

	fmt.Scan(&a)

	res = digital_root(a)

	fmt.Print(res)
}
