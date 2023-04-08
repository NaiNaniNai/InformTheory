package main

import (
	"fmt"
	"math"
)

func permute(array []string, prefix string, size int) {
	if size == 0 {
		fmt.Println(prefix)

		return
	}

	for _, i := range array {
		newPrefix := prefix + i
		permute(array, newPrefix, size-1)
	}
}

func main() {
	var size int
go
	fmt.Print("Введите длину алфавита: ")
	fmt.Scanln(&size)

	var array []string = make([]string, size)

	for i := 0; i < len(array); i++ {
		fmt.Printf("Введите %d 'ую букву массива: ", i+1)
		fmt.Scanln(&array[i])

	}

	fmt.Println("Задание а: ")
	fmt.Printf("Все комбинации перестановки: \n")
	permute(array, "", size)

	fmt.Println("Задание б: ")

	var I float64
	var massagesCount float64

	massagesCount = math.Pow(float64(size), float64(size))
	I = 1 * math.Log2(float64(massagesCount))

	fmt.Printf("Кол-во информации на одно сообщение: %f \n", I)

	fmt.Println("Задание в: ")

	var H float64

	H = math.Log2(float64(size))

	fmt.Printf("Кол-во информации на символ первичного алфавита: %f \n", H)

}
