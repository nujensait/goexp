/**
Дается строка. Нужно удалить все символы, которые встречаются более одного раза и вывести получившуюся строку

Sample Input:
zaabcbd

Sample Output:
zcd
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var cnt [1000]int

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	X := scanner.Text()

	// Поиск по буквам
	for _, ch := range X {
		//fmt.Println(ch, string(ch)) // выводим код символа и его строковое представление
		cnt[ch]++
	}

	for _, ch := range X {
		if cnt[ch] <= 1 {
			fmt.Printf("%s", string(ch))
		}
	}
}
