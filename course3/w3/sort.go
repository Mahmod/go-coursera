package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sort"
	"sync"
	"strings"
)

func main() {
	arr := readInput()
	fmt.Println("Unsorted input array length:", len(arr), " ,values:",arr)
	fmt.Println("Sorting array in 4 parts...")
	sortInParts(arr)
	fmt.Println("Sorted array:", arr)
}

// readInput handles the user input for the array of integers.
func readInput() []int {
	fmt.Println("Enter integers separated by spaces ,then press Enter when you finish:")
	scanner := bufio.NewScanner(os.Stdin)

	var arr []int
	if scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " ")
		for _, token := range tokens {
			num, err := strconv.Atoi(token)
			if err != nil {
				fmt.Printf("Invalid input '%s'. Skipping this value.\n", token)
				continue
			}
			arr = append(arr, num)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}

	return arr
}

// sortInParts divides the array into 4 parts and sorts each part in a separate goroutine.
func sortInParts(arr []int) {
	partSize := (len(arr) + 3) / 4 // Ensures even distribution of elements
	var wg sync.WaitGroup

	for i := 0; i < 4; i++ {
		start := i * partSize
		end := start + partSize
		if end > len(arr) {
			end = len(arr)
		}
		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			subArray := arr[start:end]
			fmt.Println("coroutine:",i, " ,Sorting subarray :", subArray)
			sort.Ints(subArray)
		}(start, end)
	}

	wg.Wait()

	// Merge the sorted subarrays
	sort.Ints(arr)
}
