package sleep_sort

import (
	"time"
)

func sleepAndSort(val, minimum int, res chan int) {
	// Sleeping for time proportional to value
	time.Sleep(time.Duration(val-minimum) * time.Millisecond)

	// return the value
	res <- val + minimum
}

// sleepSort Sorts an integer slice using sleep sort algorithm
func SleepSort(numbers []int) []int {
	res := make(chan int, len(numbers))
	minimum := 0
	//finding minimum value of the slice
	//so we can shift the negative values and make it positive
	for _, val := range numbers {
		if val < minimum {
			minimum = val
		}
	}
	for _, val := range numbers {
		// Spinning a Go routine
		go sleepAndSort(val-minimum, minimum, res)
	}
	for i := 0; i < len(numbers); i++ {
		numbers[i] = <-res
	}
	return numbers
}
