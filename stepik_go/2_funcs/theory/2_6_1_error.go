/*
Errors handling / Обработка ошибок
https://stepik.org/lesson/264473/step/1?auth=login&unit=245397

Run samples:
go run .\2_6_1_error.go
7
1

go run .\2_6_1_error.go
hgfdgf
Проверьте типы входных параметров.	
*/

package main

import "fmt"

func divide(a int, b int) int {
	return a / b
}

func main() {
	var input int
	_, err := fmt.Scan(&input) // fmt.Scan возвращает два параметра, первый — это количество считанных значений, второй — ошибка
	if err != nil {
		fmt.Println("Проверьте типы входных параметров.")
	} else {
		fmt.Println(divide(input, 5)) // Выводим результат, если ошибок нет
	}
}