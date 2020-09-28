// Author : @belikesayantan

package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var strArray []string
	var intArray []int

	fmt.Println("Enter an array with spaces")
	reader := bufio.NewReader(os.Stdin)
	arr, err := reader.ReadString('\n')
	if err != nil {
		inputErr := errors.New("something is wrong with your array input")
		log.Fatal(inputErr)
	}
	arr = strings.TrimSpace(arr)
	strArray = strings.Split(arr, " ")

	for _, val := range strArray {
		ele, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		intArray = append(intArray, ele)
	}

	start := time.Now() // timer starts and records running time of BubbleSort Algorithm
	fmt.Println("Sorted Array is: ", BubbleSort(intArray))
	end := time.Now() // timer stops and recording running time of BubbleSort Algorithm is completed

	fmt.Println("\nRunning Time of BubbleSort Algorithm is: ", end.Sub(start))
}

// BubbleSort Algorithm Time Complexity : O(n^2)
func BubbleSort(arr []int) []int {

	if len(arr) == 0 {
		return nil
	}

	if len(arr) == 1 {
		return arr
	}

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {

			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = swap(arr[j], arr[j+1])
			}
		}
	}

	return arr
}

func swap(i, j int) (a, b int) {
	i = i + j
	j = i - j
	i = i - j
	return i, j
}

/*
	I/O Sample:

	Sample 1:
	❯ go run bubbleSort.go
	Enter an array with spaces
	5 3 2 4 5 84 6 1 54
	Sorted Array is:  [1 2 3 4 5 5 6 54 84]

	Running Time of BubbleSort Algorithm is: 90.154µs


	Sample 2:
	❯ go run bubbleSort.go
	Enter an array with spaces
	5 8 6 9 1 3 2 7 4
	Sorted Array is:  [1 2 3 4 5 6 7 8 9]

	Running Time of BubbleSort Algorithm is:  116.805µs



	Time Complexity:
		1. Ω(n) in best case
		2. O(n^2) in average case
		3. O(n^2) in worst case

	Space Complexity: O(1) in worst case

*/
