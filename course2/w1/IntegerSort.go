package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

var (
	sortingAlgs = map[string]func([]int){
		"BubbleSort":    BubbleSort,
		"InsertionSort": InsertionSort,
		"QuickSort":     QuickSort,
	}
)

func Swap(a []int, j int) {
	a[j], a[j+1] = a[j+1], a[j]
}

func BubbleSort(a []int) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j+1] < a[j] {
				Swap(a, j)
			}
		}
	}
}

func InsertionSort(a []int) {
	for i := 0; i < len(a); i++ {
		j := i
		for j > 0 {
			if a[j-1] > a[j] {
				Swap(a, j-1)
			}
			j--
		}
	}
}

func QuickSort(a []int) {
	if len(a) < 2 {
		return
	}
	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	QuickSort(a[:left])
	QuickSort(a[left+1:])
}

func main() {
	values, _ := ReadValues()
	algorithm := GetAlgorithm()

	algorithm(values)
	fmt.Println("Sorted list of numbers:")
	fmt.Println(values)
}

func ReadValues() (values []int, err error) {
	fmt.Println("Please input numbers(separate with space):")
	br := bufio.NewReader(os.Stdin)
	numbers, _, err := br.ReadLine()

	nslice := strings.Split(string(numbers), " ")
	for _, digit := range nslice {
		n, _ := strconv.Atoi(digit)
		values = append(values, n)
	}

	return values, err
}

func InputAlgorithm() (string, error) {
	fmt.Println("Please input which sorting algorithm you want to use:")
	br := bufio.NewReader(os.Stdin)
	algorithm, _, err := br.ReadLine()

	return string(algorithm), err
}

func ChooseAlgorithm(algorithm string) func([]int) {
	var result func([]int)
	if f, ok := sortingAlgs[algorithm]; ok != true {
		fmt.Printf(`Sorry, Algorithm not implemented!
        Choose an algorithm from the list.`)
	} else {
		result = f
	}
	return result
}

func GetAlgorithm() func([]int) {
	var function func([]int)

	fmt.Println("List of implemented algorithms: ")
	PrintKeys(sortingAlgs)

	if algorithm, err := InputAlgorithm(); err != nil {
		fmt.Println("\nError: ", err)
	} else {
		function = ChooseAlgorithm(algorithm)
	}
	return function
}

func PrintKeys(m map[string]func([]int)) []string {
	keys := []string{}
	for k := range m {
		keys = append(keys, k)
	}
	fmt.Println(keys)
	return keys
}
