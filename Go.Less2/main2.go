package main

import (
	"fmt"
	"math"
)

func main() {

	/* Задание 1: расчет площади прямоуг-ка */
	var a int // сторона прямоугольника
	var b int // сторона прямоугольника

	fmt.Print("Введите знач-я сторон прямоуг-ка: ") // запрос знач-й у польз-ля
	fmt.Scanln(&a, &b)

	var square int = (a * b)                          // расчёт площади
	fmt.Println("Площ прямоуг-ка:", square, "кв. см") // вывод ответа

	/* Задание 2: расчет диаметра и длины окруж по её площади */
	var squareСircle float64
	fmt.Println("Введите площадь круга: ")
	fmt.Scanln(&squareСircle)

	var diametrCircle float64 = math.Round(2 * math.Sqrt(squareСircle/math.Pi)) //
	fmt.Println("Диаметр круга: ", diametrCircle)

	var lengthCircle float64 = math.Round(diametrCircle * math.Pi) //
	fmt.Println("Длина окружности: ", lengthCircle)
	// fmt.Println(math.Pi) // проверка знач-я числа

	/* Задание3: Выделение из трёхзначного числа сотен, десятков и тысяч */
	var number int
	fmt.Println("Введите трехзначное число: ") // запрос числа у польз-ля
	fmt.Scanln(&number)                        // приём числа

	var units = number % 10
	var tens = (number / 10) % 10
	var hundreds = (number / 100) % 10

	fmt.Println(units)

	fmt.Println("В числе ", number,
		"сотен:", hundreds, ", десятков:", tens, ", единиц:", units)
}
