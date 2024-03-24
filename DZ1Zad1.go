package main

import "fmt"

//Задание 1. Слияние отсортированных массивов
// Напишите функцию, которая производит слияние двух отсортированных массивов
//длиной четыре и пять в один массив длиной девять.

const size1 = 4
const size2 = 5
const size3 = size1 + size2

func merge(input1 [size1]int, input2 [size2]int) (output [size3]int) {
	for i := 0; i < len(input1); i++ {
		output[i] = input1[i]
	}
	j := size3 - size2
	for i := 0; i < len(input2); i++ {
		output[j] = input2[i]
		j++
	}
	return output
}

func main() {
	massiv1 := [size1]int{1, 2, 3, 4}
	massiv2 := [size2]int{5, 6, 7, 8, 9}
	fmt.Println(merge(massiv1, massiv2))
}
