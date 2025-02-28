/**
В рамках этого урока мы постарались представить себе уже привычные нам переменные и функции, как объекты из реальной жизни.
Чтобы закрепить результат мы предлагаем вам небольшую творческую задачу.

Вам необходимо реализовать структуру со свойствами-полями On, Ammo и Power, с типами bool, int, int соответственно.
У этой структуры должны быть методы: Shoot и RideBike, которые не принимают аргументов, но возвращают значение bool.

Если значение On == false, то оба метода вернут false.

Делать Shoot можно только при наличии Ammo (тогда Ammo уменьшается на единицу, а метод возвращает true),
если его нет, то метод вернет false. Метод RideBike работает также, но только зависит от свойства Power.

Чтобы проверить, что вы все сделали правильно, вы должны создать указатель на экземпляр этой структуры с именем testStruct в функции main,
в дальнейшем программа проверит результат.
*/

package main

import "fmt"

type TestStruct struct {
	On    bool
	Ammo  int
	Power int
}

func (t *TestStruct) Shoot() bool {
	if !t.On || t.Ammo == 0 {
		return false
	}
	t.Ammo--
	return true
}

func (t *TestStruct) RideBike() bool {
	if !t.On || t.Power == 0 {
		return false
	}
	t.Power--
	return true
}

func main() {
	testStruct := &TestStruct{
		On:    true,
		Ammo:  5,
		Power: 10,
	}

	// Вывод значений testStruct
	fmt.Println("On:", testStruct.On)
	fmt.Println("Ammo:", testStruct.Ammo)
	fmt.Println("Power:", testStruct.Power)
}
