package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64 // задание одного числа в формате дробного
	var b float64 // задание второго числа в формате дробного

	fmt.Println("Введите числа a и b: (для факториала одно число)")
	fmt.Scanln(&a, &b)

	var oper string
	fmt.Println("Введите операцию: (возведение в степень: n, факториал: !)")
	fmt.Scanln(&oper)
	calc(a, b, oper)
}

func calc(a, b float64, oper string) {
	switch oper {
	case "+": // сложение
		fmt.Printf("Результат сложения %g и %g это: %.2f\n", a, b, a+b)
		// fmt.Printf("%.2f", a+b)
	case "-": // вычитание
		fmt.Printf("Результат вычитания %g и %g это: %.2f\n", a, b, a-b)
		// fmt.Printf("%.2f\n", a-b)
	case "*": // умножение
		fmt.Printf("Результат умножения %g и %g это: %.2f\n", a, b, a*b)
		// fmt.Printf("%.2f\n", a*b)
	case "/": // деление с проверкой деления на ноль с пом отд ф-ции
		division(a, b, oper)
	case "n": // возвед в степень с пом встроен ф-ии math
		fmt.Printf("Результат возведения %g в степень %g это: %.2f\n", a, b, math.Pow(a, b))
		// fmt.Printf("%.2f\n", math.Pow(a, b))
	case "!": // взятие факториала с пом отд ф-ции
		fmt.Printf("Факториал %.2f = %.2f\n", a, factorial(a))
	}
}

/* ф-я с проверкой дел-я на 0 */
func division(a, b float64, oper string) {
	if b == 0 {
		fmt.Print("Делить на ноль нельзя.\n")
	} else {
		fmt.Printf("Результат деления %g на %g это: %.2f\n", a, b, a/b)
		// fmt.Printf("%.2f\n", a/b)
	}
}

/*
		ф-я факториала,
	  написала не сама, взяла код из инета
*/
func factorial(a float64) float64 {
	var result float64 = 1 // созд перем-ой д/хран-я рез-та, если а = 1
	if a < 0 {             // если а < 0, то
		// ничего не делать
	} else { // иначе запустить цикл
		var i float64
		for i = 1; i <= a; i++ {
			result *= float64(i)
		}
	}
	return result
}

