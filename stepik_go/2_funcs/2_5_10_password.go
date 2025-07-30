/**
Ваша задача сделать проверку подходит ли пароль вводимый пользователем под заданные требования.
Длина пароля должна быть не менее 5 символов, он может содержать только арабские цифры
и буквы латинского алфавита.

На вход подается строка-пароль.
Если пароль соответствует требованиям - вывести "Ok", иначе вывести "Wrong password"

Sample Input:
fdsghdfgjsdDD1

Sample Output:
Ok
*/

package main

import (
	"fmt"
	"unicode"
)

func main() {
	var password string
	fmt.Scan(&password)

	if isValidPassword(password) {
		fmt.Println("Ok")
	} else {
		fmt.Println("Wrong password")
	}
}

func isValidPassword(password string) bool {
	// Check length
	if len(password) < 5 {
		return false
	}

	// Check each character
	for _, char := range password {
		// Check if character is a letter or digit
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
			return false
		}
		// Check if letter is Latin
		if unicode.IsLetter(char) && !isLatinLetter(char) {
			return false
		}
	}

	return true
}

func isLatinLetter(char rune) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}
