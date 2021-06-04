package main

import (
	"fmt"
	"os"
	"sleepSort/sleep_sort"
	"strconv"
)

func main() {
	args := os.Args[1:]
	var numbers = []int{}
	for _, val := range args {
		valInt, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, valInt)
	}
	sleep_sort.SleepSort(numbers)
	for _, val := range numbers {
		fmt.Printf("%d ", val)
	}
	fmt.Println()

}
