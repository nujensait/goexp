/**
На вход подается строка. Нужно определить, является ли она палиндромом.
Если строка является палиндромом - вывести Палиндром иначе - вывести Нет.
Все входные строчки в нижнем регистре.

Палиндром — буквосочетание, слово или текст,
одинаково читающееся в обоих направлениях (например, "топот", "око", "заказ").

Sample Input:
топот

Sample Output:
Палиндром
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {

	//text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	//text = strings.TrimSpace(text) // Удаляем пробелы в начале и конце

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	//text = text[:len(text)-1] // Удаляем перевод строки
	text = strings.TrimSpace(text) // Удаляем пробелы в начале и конце

	// Проверка палиндрома
	len := utf8.RuneCountInString(text)

	for i := 0; i < len/2; i++ {
		if text[i] != text[len-1-i] {
			fmt.Printf("Буквы не совпали: %c и %c\n", text[i], text[len-1-i])
			fmt.Println("Нет")
			return
		}
	}
	fmt.Println("Палиндром")
}
