// Go содержит большое количество функций для работы со строками в пакете strings , кратко рассмотрим основные:
// Usage:
// go run ./theory/2_5_2_strings.go

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(
		// Содержится ли подстрока в строке
		strings.Contains("test", "es"), "\n\n",
		// результат: true

		// Кол-во подстрок в строке
		strings.Count("test", "t"), "\n\n",
		// результат: 2

		// Начинается ли строка с префикса
		strings.HasPrefix("test", "te"), "\n\n",
		// результат: true

		// Заканчивается ли строка суффиксом
		strings.HasSuffix("test", "st"), "\n\n",
		// результат: true

		// Возвращает начальный индекс подстроки в строке, а при отсутствии вхождения возвращает -1
		strings.Index("test", "e"), "\n\n",
		// результат: 1

		// объединяет массив строк через символ
		strings.Join([]string{"hello", "world"}, "-"), "\n\n",
		// результат: "hello-world"

		// Повторяет строку n раз подряд
		strings.Repeat("a", 5), "\n\n",
		// результат: "aaaaa"

		// Функция Replace заменяет любое вхождение old в вашей строке на new
		// Если значение n равно -1, то будут заменены все вхождения.
		// Общий вид: func Replace(s, old, new string, n int) string
		// Пример:
		strings.Replace("blanotblanot", "not", "***", -1), "\n\n",
		// результат: "bla***bla***"

		// Разбивает строку согласно разделителю
		strings.Split("a-b-c-d-e", "-"), "\n\n",
		// результат: []string{"a","b","c","d","e"}

		// Возвращает строку c нижним регистром
		strings.ToLower("TEST"), "\n\n",
		// результат: "test"

		// Возвращает строку c верхним регистром
		strings.ToUpper("test"), "\n\n",
		// результат: "TEST"

		// Возвращает строку c верхним регистром (кириллица)
		strings.ToUpper("тест"), "\n\n",
		// результат: "ТЕСТ"

		// Возвращает строку с вырезанным набором
		strings.Trim("tetstet", "te"), "\n\n",
		// результат: s
	)
}
