package main

import "fmt"

func main() {
	var a int // сторона прямоугольника
	var b int // сторона прямоугольника

	fmt.Print("Введите знач-я сторон прямоуг-ка: ") // запрос знач-й у польз-ля
	fmt.Scanln(&a, &b)

	var square int = (a * b) // расчёт площади
	fmt.Println("Площ прямоуг-ка:", square, "кв. см") // вывод ответа
}
