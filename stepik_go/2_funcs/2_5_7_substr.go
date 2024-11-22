/**
Даются две строки X и S. Нужно найти и вывести первое вхождение подстроки S в строке X.
Если подстроки S нет в строке X - вывести -1

Sample Input:
awesome
es

Sample Output:
2
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	X := scanner.Text()

	scanner.Scan()
	S := scanner.Text()

	// simple solution:
	//index := strings.Index(X, S)
	//fmt.Println(index)

	// Поиск по буквам
	for i := 0; i <= len(X)-len(S); i++ {
		found := true
		for j := 0; j < len(S); j++ {
			if X[i+j] != S[j] {
				found = false
				break
			}
		}
		if found {
			fmt.Println(i)
			return
		}
	}

	fmt.Println(-1)
}
