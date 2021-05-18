package main

/*
Внутри функции main (объявлять функцию не нужно) необходимо написать программу:

На первом этапе на стандартный ввод подается 10 целых положительных чисел,
которые должны быть записаны в порядке ввода в массив из 10 элементов.
Тип чисел, входящих в массив, должен соответствовать минимально возможному целому беззнаковому числу.
Имя массива который вы должны сами создать workArray (условие обязательное).
Для чтения из стандартного ввода уже импортирован пакет fmt.

На втором этапе на стандартный ввод подаются еще 3 пары чисел - индексы элементов этого массива,
которые требуется поменять местами (если такая пара чисел 3 и 7,
значит в массиве элемент с 3 индексом нужно поменять местами с элементом, индекс которого 7).

Элементы полученного массива должны быть выведены через пробел на стандартный вывод.
Далее автоматически будет проведена проверка используемых типов, результат которой будет добавлен к вашему ответу.

Использование массива - обязательное условие!
 */

var num uint8
var idxOne uint8
var idxTwo uint8

workArray := [10]uint8{}

for i := 0; i < len(workArray); i++ {
fmt.Scan(&num)
workArray[i] = num
}
for u := 0; u <= 6; u++{
fmt.Scan(&idxOne)
fmt.Scan(&idxTwo)
workArray[idxOne], workArray[idxTwo] = workArray[idxTwo], workArray[idxOne]
}
for idx, _ := range workArray {
fmt.Print(workArray[idx], " ")
}

/*
Напишите программу, принимающая на вход число N (N \geq 4)N(N≥4), а затем NN целых чисел-элементов среза.
На вывод нужно подать 4-ый (3 по индексу) элемент данного среза.
 */

import "fmt"

func main() {
	var num int
	var sliceArray []int
	var numArray int

	fmt.Scan(&num)

	for i := 0; i <= num; i++ {
		fmt.Scan(&numArray)

		sliceArray = append(sliceArray, numArray)
	}
	fmt.Println(sliceArray[3])
}

/*
На ввод подаются пять чисел, которые записываются в массив. Однако эта часть программы уже написана.

Вам нужно написать фрагмент кода, с помощью которого можно найти и вывести максимальное число в этом массиве.
 */

import (
"fmt"
"sort"
)

func main()  {
	array := [5]int{}
	var a int
	var maxArray []int

	for i:=0; i < 5; i++{
		fmt.Scan(&a)
		array[i] = a
	}
	for _, elem := range array {
		maxArray = append(maxArray, elem)
	}
	sort.Ints(maxArray)
	fmt.Println(maxArray[4])
}

/*
Дан массив, состоящий из целых чисел. Нумерация элементов начинается с 0.
Напишите программу, которая выведет элементы массива, индексы которых четны (0, 2, 4...).

Входные данные

Сначала задано число NN — количество элементов в массиве (1 \leq N \leq 1001≤N≤100).
Далее через пробел записаны NN чисел — элементы массива. Массив состоит из целых чисел.

Выходные данные

Необходимо вывести все элементы массива с чётными номерами.
 */

func main() {

	var num int
	var numIN int

	fmt.Scan(&num)

	testArray := make([]int, num)
	var totalArray []int

	for idx, _ := range testArray {
		fmt.Scan(&numIN)
		testArray[idx] = numIN
	}
	for idx, elem :=range testArray {
		if idx % 2 == 0 {
			totalArray = append(totalArray, elem)
		}
	}
	for idx, _ := range totalArray {
		fmt.Print(totalArray[idx], " ")
	}
}

/*
Дана последовательность, состоящая из целых чисел.
Напишите программу, которая подсчитывает количество положительных чисел среди элементов последовательности.

Входные данные

Сначала задано число NN — количество элементов в последовательности (1\leq N\leq1001≤N≤100).
Далее через пробел записаны NN чисел — элементы последовательности. Последовательность состоит из целых чисел.

Выходные данные

Необходимо вывести единственное число - количество положительных элементов в последовательности.
 */

func main() {

	var count, num, totalCount int

	fmt.Scan(&count)

	numsArray := make([]int, count)

	for idx, _ := range numsArray {
		fmt.Scan(&num)
		numsArray[idx] = num
	}
	for _, elem := range numsArray {
		if elem > 0 {
			totalCount++
		}
	}
	fmt.Println(totalCount)
}
