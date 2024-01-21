package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isRoman(input string) bool {
	// Проверяем, содержит ли строка римские цифры.
	return strings.ContainsAny(input, "IVXLCDM")
}

func romanToArabic(roman string) (int, error) {
	romanNumerals := map[rune]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

	var result int
	prevValue := 0

	for i := len(roman) - 1; i >= 0; i-- {
		value := romanNumerals[rune(roman[i])]

		if value < prevValue {
			result -= value
		} else {
			result += value
		}

		prevValue = value
	}

	return result, nil
}

func arabicToRoman(arabic int) (string, error) {
	if arabic <= 0 || arabic > 3999 {
		return "", fmt.Errorf("неверное арабское число: %d", arabic)
	}

	romanNumerals := []struct {
		Value  int
		Symbol string
	}{
		{1000, "M"}, {900, "CM"},
		{500, "D"}, {400, "CD"},
		{100, "C"}, {90, "XC"},
		{50, "L"}, {40, "XL"},
		{10, "X"}, {9, "IX"},
		{5, "V"}, {4, "IV"},
		{1, "I"},
	}

	var result strings.Builder

	for _, numeral := range romanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String(), nil
}

func main() {
	var num1Str, num2Str string
	var operator string

	fmt.Print("Введите первое число (от 1 до 10): ")
	fmt.Scanln(&num1Str)

	fmt.Print("Введите оператор (+, -, *, /): ")
	fmt.Scanln(&operator)

	fmt.Print("Введите второе число (от 1 до 10): ")
	fmt.Scanln(&num2Str)

	// Проверяем, что оба числа имеют одинаковый формат (либо арабские, либо римские).
	if (isRoman(num1Str) && !isRoman(num2Str)) || (!isRoman(num1Str) && isRoman(num2Str)) {
		fmt.Println("Ошибка: оба числа должны быть в одном формате (либо арабские, либо римские)")
		return
	}

	num1, err := strconv.Atoi(num1Str)
	if err != nil || num1 < 1 || num1 > 10 {
		// Если не удалось преобразовать в арабское число, то, возможно, это римское число.
		num1, err = romanToArabic(num1Str)
		if err != nil || num1 < 1 || num1 > 10 {
			fmt.Println("Ошибка: неверное первое число")
			return
		}
	}

	num2, err := strconv.Atoi(num2Str)
	if err != nil || num2 < 1 || num2 > 10 {
		// Если не удалось преобразовать в арабское число, то, возможно, это римское число.
		num2, err = romanToArabic(num2Str)
		if err != nil || num2 < 1 || num2 > 10 {
			fmt.Println("Ошибка: неверное второе число")
			return
		}
	}

	var result int

	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Println("Ошибка: деление на ноль")
			return
		}
		result = num1 / num2
	default:
		fmt.Println("Неверный оператор")
		return
	}

	// Если оба числа ввода были арабскими, выведем результат как арабское число.
	if isRoman(num1Str) || isRoman(num2Str) {
		// Проверяем, что результат с римскими числами не меньше единицы.
		if result < 1 {
			fmt.Println("Ошибка: результат с римскими числами не может быть меньше единицы")
			return
		}

		romanResult, err := arabicToRoman(result)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Результат: ", romanResult)
	} else {
		fmt.Println("Результат: ", result)
	}
}
