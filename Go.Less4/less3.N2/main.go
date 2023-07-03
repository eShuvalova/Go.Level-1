package main

import "fmt"

func main() {
	/*
	 * Необязательное задание из урока 3: Написать приложение,
	 * которое ищет все простые числа от 0 до N включительно.
	 * Число N должно быть задано из стандартного потока ввода.
	 */

	var number int // объявл-е перем-ой для введённого числа
	fmt.Println("Введите число: ")
	fmt.Scanln(&number)

	/* созд-е слайса интов с нуля до введенного числа */
	var sl []int = make([]int, 1, number)
	for i := 1; i <= number; i++ {
		sl = append(sl, i) // добавл-е эл-тов в слайс
	}

	// fmt.Println(sl) // печать слайса с нуля до введённого числа

	var flag int = 0 // созд-е перем-ой-флага
	fmt.Printf("Простые числа: \n")

	for i := 0; i <= len(sl)-1; i++ {
		flag = 0
		for j := 2; j < sl[i]/2; j++ {
			if sl[i]%j == 0 {
				flag = 1
				break
			}
		}
		if flag == 0 && sl[i] > 1 {
			fmt.Printf("%d ", sl[i])
		}
	}
	fmt.Println()
}

/* пример из интернета с заранее известным массивом*/
/*
	arr := [6]int{10, 11, 12, 13, 65, 73}
	fmt.Println("The array elements are: ", arr)

	var flag int = 0
	fmt.Printf("Prime Numbers: \n")

	for i := 0; i <= len(arr)-1; i++ {
		flag = 0
		for j := 2; j < arr[i]/2; j++ {
			if arr[i]%j == 0 {
				flag = 1
				break
			}
		}
		if flag == 0 && arr[i] > 1 {
			fmt.Printf("%d ", arr[i])
		}
	}
*/
