package main

import (
	"fmt"
	"math"
)

func main() {
	//простейший вывод на консоль. PrintLn - вывод аргумента + '\n'
	fmt.Println("Hello, World!")
	fmt.Println("Second Line")
	//функция Print
	fmt.Print("First")
	fmt.Print("Second")
	fmt.Print("Third")
	//форматированный вывод: Printf - стандартный вывод os.Stdout с флагами форматирования
	fmt.Printf("\nHello, my name is %s\nMy age is %d\n", "Oleg", 35)

	////////////////////////////////////////////////////////////////////////////////////////////
	/////////////Декларирование переменных/////////////////////////////////////////////////////
	//////////////////////////////////////////////////////////////////////////////////////////

	var age int
	fmt.Println("My age is:", age)
	age = 35
	fmt.Println("Age after assignment:", age)

	//////////////////////////////////////////////////////////////////////////////////////////
	/////////////Декларирование и инициализация пользовательским значением////////////////////

	var height int = 180
	fmt.Println("My height is:", height)

	//В Чем "Получтрогость" типизации

	var uid = 12345
	fmt.Println("My uid is:", uid)

	//Декларирование и инициализация переменных дного типа (множественный случай)
	var firstvar, secondvar int = 20, 30
	fmt.Printf("FirstVar:%d SecondVar:%d\n", firstvar, secondvar)

	//Декларирование блока переменных

	var (
		personName string = "Bob"
		personAge  int    = 42
		peronUID   int
	)

	fmt.Printf("Name: %s\nAge: %d\nUID: %d\n", personName, personAge, peronUID)

	//немного странного
	var a, b = 30, "Vova"
	fmt.Println(a, b)

	//немного хорошего. Повторное декларирование переменной приводит к ошибке компиляции
	//var a = 200

	//Короткая деклрация (короткое объявление)
	count := 10
	fmt.Println("Count", count)
	count = count + 1
	fmt.Println("Count", count)

	//Множественное присваивание через :=
	aArg, bArg := 10, 30
	fmt.Println(aArg, bArg)
	aArg, bArg = 30, 40
	fmt.Println(aArg, bArg)
	//aArg, bArg := 10, 30
	//fmt.Println(aArg, bArg)

	//Исключение из этого правила
	bArg, cArg := 300, 400
	fmt.Println(aArg, bArg, cArg)

	//Пример
	width, length := 20.5, 30.2
	fmt.Printf("Min dementional of rectange is: %.3f\n", math.Min(width, length))

}
