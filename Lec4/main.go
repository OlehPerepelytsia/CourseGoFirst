package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
	"unsafe"
)

func main() {
	//Boolean => default false
	var firstBoolean bool
	fmt.Println(firstBoolean)

	//Boolean operands
	aBoolean, bBoolean := true, true
	fmt.Println("AND:", aBoolean && bBoolean)
	fmt.Println("OR:", aBoolean || bBoolean)
	fmt.Println("NOT:", !aBoolean)

	//Numerics. Intergers
	//int8, int16, int32, int64, int
	//uint8, uint16, uint32, ui nt64, uint

	var a int = 32
	b := 92

	fmt.Println("Value of a:", a, "Value of b:", b, "Sum of a+b:", a+b)
	//Вывод типа через %T форматирование
	fmt.Printf("Type is %T\n", a)
	//Узнаем, сколько байт занимает переменная типа int
	fmt.Printf("Type %T size of % d bytes\n", a, unsafe.Sizeof(a))

	//Эксперимент. При использовании короткого объявления - тип для целого числа - int, платформозависимый
	fmt.Printf("Type %T size of % d bytes\n", b, unsafe.Sizeof(b))

	//Эксперимент №2. Используйте явное применение типов при необходимости если уверены что не произойдет коллизии.
	var first32 int32 = 12
	var second64 int64 = 13
	fmt.Println(int64(first32) + second64)

	//Эксперимент №3. Если проводятся арифметические опреации над int и над intX, то обязательно нужно использовать механизм приведения т.к. int!=int64
	var third64 int64 = 16123414
	var fourthint int = 156234
	fmt.Println(third64 + int64(fourthint))

	//+ - * / %
	//Аналогичным образом устроены uint8, uint16, uint32, uint64, uint

	//Numerics. Float
	//float32, float64

	floatfirst, floatsecond := 5.67, 12.54
	fmt.Printf("Type of a %T and Type of b %T\n", floatfirst, floatsecond)
	sum := floatfirst + floatsecond
	sub := floatfirst - floatsecond

	fmt.Println("Sum", sum, "Sub", sub)
	fmt.Printf("Sum:%.3f and Sub:%.3f\n", sum, sub)

	//Numerics. Complex
	c1 := complex(5, 7)
	c2 := 12 + 32i
	fmt.Println(c1 + c2)
	fmt.Println(c1 * c2)

	//Strings Строка  - это набор БАЙТ
	name := "Федя"
	lastName := "Pupkin"
	concat := name + " " + lastName
	fmt.Println("Full name:", concat)
	fmt.Println("Length of string:", name, len(name)) //Функция len() возвращает количество элементов в наборе
	fmt.Println("Amount of chars", name, utf8.RuneCountInString(name))
	//Rune  - это один utf-ный символ
	//Поиск подстроки в строке
	totalstring, substring := "ABCDEDFG", "CDE"
	fmt.Println(strings.Contains(totalstring, substring))
	//rune -> alias int32
	var sampleRune rune
	var anotherRune rune = 'Q'
	var thirdRune rune = 234
	fmt.Println(sampleRune)
	fmt.Printf("Runes as char %c \n", sampleRune)
	fmt.Printf("Runes as char %c \n", anotherRune)
	fmt.Printf("Runes as char %c \n", thirdRune)
	//"A" , "abcd"
	fmt.Println(strings.Compare("abcd", "a")) // -1 if firsrt < second, 0 if first == second, 1 first > second
	var aByte byte                            //alias int8
	fmt.Println("Byte:", aByte)
}
