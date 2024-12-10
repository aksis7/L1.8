package main

import (
	"fmt"
)

// SetBit изменяет i-й бит числа num на значение value (1 или 0).
func SetBit(num int64, i uint, value bool) int64 {
	if value {
		// Установить i-й бит в 1
		return num | (1 << i)
	} else {
		// Установить i-й бит в 0
		return num &^ (1 << i)
	}
}

func main() {
	var num int64
	var i uint
	var value int

	// Ввод числа
	fmt.Print("Введите число: ")
	fmt.Scan(&num)

	// Ввод индекса бита
	fmt.Print("Введите индекс бита (начиная с 0): ")
	fmt.Scan(&i)

	// Ввод значения для бита (0 или 1)
	fmt.Print("Введите значение бита (0 или 1): ")
	fmt.Scan(&value)

	if value != 0 && value != 1 {
		fmt.Println("Ошибка: значение бита должно быть 0 или 1")
		return
	}

	// Установить значение бита
	originalNum := num
	num = SetBit(num, i, value == 1)

	// Вывод результатов
	fmt.Printf("Исходное число: %b (%d)\n", originalNum, originalNum)
	fmt.Printf("Измененное число: %b (%d)\n", num, num)
}
