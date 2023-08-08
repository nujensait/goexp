// https://stepik.org/lesson/228838/step/5?unit=201372
// Возвращение нескольких значений
// В Go функция может возвращать сразу несколько значений.
// В этом случае после списка параметров указывается в скобках список типов возвращаемых значений.
// А после оператора return располагаются через запятую все возвращаемые значения:
// Функция add принимает четыре параметра: два числа и две строки. Возвращает число (значение типа int) и строку.
// Возвращаемые значения указаны после оператора return.

package main

import "fmt"

func main() {
	var age, name = CustomAdd(4, 5, "Tom", "Simpson")
	fmt.Println(age)  // 9
	fmt.Println(name) // Tom Simpson
}

func CustomAdd(x, y int, firstName, lastName string) (int, string) {
	var z int = x + y
	var fullName = firstName + " " + lastName
	return z, fullName
}
