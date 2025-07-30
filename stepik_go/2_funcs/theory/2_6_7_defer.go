/*
Оператор defer позволяет выполнить определенную операцию после 
выполнения других действий (даже если сработает panic), 
при этом не важно, где в реальности вызывается эта функция. 
Дополнение: Команда defer помещает вызов функции в стек. 
Поэтому вызовы функций выполняются в очередности LIFO (Last-In, First-Out), 
то есть последняя отложенная функция будет выполнена первой.
*/

package main
import "fmt"

func main() {
    defer finish() // Этот вызов отложен
    defer fmt.Println("Program has been started") // Этот вызов отложен
    fmt.Println("Program is working")
}

func finish() {
    fmt.Println("Program has been finished")
}