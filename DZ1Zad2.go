package main

import "fmt"

//Задание 2. Сортировка пузырьком
//Отсортируйте массив длиной шесть пузырьком.

const size4 = 6

func bubbleSort(input [size4]int) [size4]int {
	for i := 0; i < len(input)-1; i++ {
		for j := 0; j < len(input)-i-1; j++ {
			if input[j] > input[j+1] {
				input[j], input[j+1] = input[j+1], input[j]
			}
		}
	}
	return input
}

func main() {
	massiv3 := [size4]int{87, 67, 12, 47, 52, 9}
	fmt.Println(bubbleSort(massiv3))
}
