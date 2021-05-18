package main

import "fmt"

/*
Напишите функцию f(), которая будет принимать строку text и выводить ее (печатать на экране).

От вас требуется дописать только эту функцию, считайте что функция main() уже объявлена, считывать с консоли ничего не нужно!

Пакет "fmt" уже импортирован!
 */

func f(text string){
	fmt.Println(text)
}

/*

Напишите функцию, находящую наименьшее из четырех введённых в этой же функции чисел.

Входные данные

Вводится четыре числа.

Выходные данные

Необходимо вернуть из функции наименьшее из 4-х данных чисел
 */

func min(numOne, numTwo, numThree, numFour int) int {
	var totalSlice []int

	totalSlice = append(totalSlice, numOne, numTwo, numThree, numFour)
	minNum := totalSlice[0]

	for _, elem := range totalSlice {

		if elem < minNum {
			minNum = elem
		}
	}
	return minNum
}

func minimumFromFour() int {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	minResult := min(a, b, c, d)

	return minResult
}

/*
Напишите "функцию голосования", возвращающую то значение (0 или 1),
которое среди значений ее аргументов x, y, z встречается чаще.

Входные данные

Вводится 3 числа - x, y и z (x, y и z равны 0 или 1).

Выходные данные

Необходимо вернуть значение функции от x, y и z.
 */

func vote(x int, y int, z int) int {

	if x == y || x == z {
		return x
	}
	return y
}

/*
Последовательность Фибоначчи определена следующим образом: φ1=1, φ2=1, φn=φn-1+φn-2 при n>1.
Начало ряда Фибоначчи выглядит следующим образом: 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, ...
Напишите функцию, которая по указанному натуральному n возвращает φn.

Входные данные

Вводится одно число n.

Выходные данные

Необходимо вывести  значение φn.
 */

func fibonacci(n int) int {
	var a, b = 0, 1

	if n == 0 {
		return n
	} else {
		for i := 2; i < n + 1; i++ {
			a, b = b, a + b
		}
	}
	return b
}

/*
Напишите функцию sumInt, принимающую переменное количество аргументов типа int,
и возвращающую количество полученных функцией аргументов и их сумму.
Пакет "fmt" уже импортирован, функция и пакет main объявлены.

Пример вызова вашей функции:

a, b := sumInt(1, 0)
fmt.Println(a, b)

Результат: 2, 1
 */

func sumInt(num ...int) (int, int) {

	var result int

	for _, elem := range num {
		result += elem
	}
	countNum := len(num)

	return countNum, result
}
