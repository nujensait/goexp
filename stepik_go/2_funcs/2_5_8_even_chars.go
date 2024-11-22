/**
На вход дается строка, из нее нужно сделать другую строку,
оставив только нечетные символы (считая с нуля)

Sample Input:
ihgewlqlkot

Sample Output:
hello
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

	// Поиск по буквам
	for i := 1; i < len(X); i++ {
		fmt.Printf("%c", X[i])
		i++
	}
}
