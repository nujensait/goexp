/**
1.5 Переменные и ариф. операции, ввод/вывод данных
https://stepik.org/lesson/228261/step/16?unit=200794

Часовая стрелка повернулась с начала суток на d градусов. Определите, сколько сейчас целых часов h и целых минут m.

Входные данные
На вход программе подается целое число d (0 < d < 360).

Выходные данные
Выведите на экран фразу:
It is ... hours ... minutes.
Вместо многоточий программа должна выводить значения h и m, отделяя их от слов ровно одним пробелом.

Sample Input:
90

Sample Output:
It is 3 hours 0 minutes.
*/

package main

import "fmt"

func main() {
	var g int
	fmt.Scan(&g)

	var h, m int

	// 360 degrees = 12 hours
	h = g * 12 / 360

	// 360 degrees = 60 minutes
	m = 2 * (g % 30)

	fmt.Print("It is ")
	fmt.Print(h)
	fmt.Print(" hours ")
	fmt.Print(m)
	fmt.Print(" minutes.")
}
