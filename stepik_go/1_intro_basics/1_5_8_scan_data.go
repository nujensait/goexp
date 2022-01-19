/**
https://stepik.org/lesson/228261/step/8?unit=200794
Чтение данных с консоли
Программа сначала прочтёт имя, а затем запишет его в переменную name.
Аналогично, введённый возраст запишется в переменную age.
В конце программа выведет эти переменные через пробел.
*/
package main

import "fmt"

func main() {
	var name string
	var age int
	fmt.Print("Введите имя: ")
	fmt.Scan(&name)
	fmt.Print("Введите возраст: ")
	fmt.Scan(&age)

	fmt.Println(name, age)
}
