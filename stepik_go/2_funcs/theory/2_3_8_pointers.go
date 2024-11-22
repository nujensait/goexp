/**
Что выведет программа на экран?
- 103
*/

package main

import "fmt"

func main() {
	a := 100
	b := &a
	*b++
	c := &b
	**c++
	*b++

	fmt.Println(a)
}
