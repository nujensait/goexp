/*
На вход подаются a и b - катеты прямоугольного треугольника. 
Нужно найти длину гипотенузы.

Sample Input:
6 8

Sample Output:
10
*/

package main

import (
    "fmt"
    "math"
)

func main() {
    var a, b float64
    fmt.Scan(&a, &b)
    
    // Вычисляем гипотенузу по теореме Пифагора: c = √(a² + b²)
    c := math.Sqrt(a*a + b*b)
    
    fmt.Println(int(c)) // Приводим к int, так как в примере результат целочисленный
}