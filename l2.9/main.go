/*
Написать функцию Go, осуществляющую примитивную распаковку строки, содержащей повторяющиеся символы/руны.

Примеры работы функции:

Вход: "a4bc2d5e"
Выход: "aaaabccddddde"

Вход: "abcd"
Выход: "abcd" (нет цифр — ничего не меняется)

Вход: "45"
Выход: "" (некорректная строка, т.к. в строке только цифры — функция должна вернуть ошибку)

Вход: ""
Выход: "" (пустая строка -> пустая строка)

Дополнительное задание
Поддерживать escape-последовательности вида \:

Вход: "qwe\4\5"
Выход: "qwe45" (4 и 5 не трактуются как числа, т.к. экранированы)

Вход: "qwe\45"
Выход: "qwe44444" (\4 экранирует 4, поэтому распаковывается только 5)

Требования к реализации
Функция должна корректно обрабатывать ошибочные случаи (возвращать ошибку, например, через error), и проходить unit-тесты.

Код должен быть статически анализируем (vet, golint).
*/

package main

import (
	"errors"
	"strconv"
)

func unpackingString(str string) (string, error) {
	result := ""

	var (
		ch      string
		num     int64 = 1
		flag          = false
		flagNum       = false
	)
	for i := 0; i < len(str); i++ {
		if str[i] >= '0' && str[i] <= '9' && !flagNum {

			num, _ = strconv.ParseInt(string(str[i]), 10, 0)
			flag = false
		} else if str[i] == '\\' {
			if i == len(str)-1 {
				return "", errors.New("Unval. str")
			}
			flagNum = true
			ch = string(str[i+1])
			i++
			num = 1
		} else {

			num = 1
			ch = string(str[i])
			flagNum = false
		}
		if !flag || num == 1 {
			for range num {
				result += ch
			}
			flag = true
		}

	}
	if result != "" {
		return result, nil
	}
	return "", errors.New("res string is empty")

}
