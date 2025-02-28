/*
@source https://stepik.org/lesson/228838/step/4?unit=201372

Функции могут возвращать результат.
Для этого нужно после списка параметров функции указать тип возвращаемого результата.
А в теле функции использовать оператор return, после которого указывается возвращаемое значение:
Например, мы хотим возвратить из функции сумму двух чисел:
*/

package main

import "fmt"

func main() {
	var a = add(4, 5)  // 9
	var b = add(20, 6) // 26
	fmt.Println(a)
	fmt.Println(b)
}

func add(x, y int) int {
	return x + y
}
