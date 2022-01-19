/**
1.13 Решение задач
https://stepik.org/lesson/229320/step/12?unit=201906

По данному числу n закончите фразу "На лугу пасется..." одним из возможных продолжений:
"n коров", "n корова", "n коровы", правильно склоняя слово "корова".

Входные данные
Дано число n (0<n<100).

Выходные данные
Программа должна вывести введенное число n и одно из слов (на латинице): korov, korova или korovy, например, 1 korova, 2 korovy, 5 korov.

Между числом и словом должен стоять ровно один пробел.

Sample Input:
10

Sample Output:
10 korov
*/

package main

import (
	"fmt"
)

func main() {
	var n int
	var cow string

	fmt.Scan(&n)
	if n == 1 {
		cow = "korova"
	} else if n >= 2 && n <= 4 {
		cow = "korovy"
	} else {
		cow = "korov"
	}

	fmt.Print(n)
	fmt.Print(" " + cow)
}
