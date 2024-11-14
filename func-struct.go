package main

import "fmt"

// 1. структуры
// 2. функции
// 3. интерфейсы

// создание типа
// инициализация экземпляра типа
// передачу в функцию
// сравниваемость
// содержимое типа, значения и указателя

// содзали тип
type car struct {
	color string
}

// инициализировали тип на уровне пакета
var a = car{
	color: "yellow",
}

func main() {
	// инициализировали тип на уровне функции main
	b := &car{
		color: "yellow",
	}

	// инициализировали тип на уровне функции main
	c := new(car)
	c.color = "red"

	compare(&a, b)
	compare(&a, c)
	compare(b, c)
	compare(b, nil)
}

func compare(a, b *car) {
	fmt.Printf("pointer 1 == pointer 2 %v\n", a == b)
	fmt.Printf("1 == 2 %v\n", *a == *b)
	fmt.Printf("Type 1 %T, 1=%v, pointer 1 %p\n", a, a, a)
	fmt.Printf("Type 2 %T, 2=%v, pointer 2 %p\n", b, b, b)
	fmt.Println()
	// вывести тут содержимое типа, значения и указателя
}
