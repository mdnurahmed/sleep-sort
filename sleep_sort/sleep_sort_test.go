package sleep_sort

import "testing"

func TestSleepSort(t *testing.T) {

	numbers := []int{-1, 3, 0, 1, 2}
	result := SleepSort(numbers)
	for i := 0; i < len(result); i++ {
		if result[i] != i-1 {
			t.Errorf("Expected %dth element to be %d, but got %d", i, i-1, result[i])
		}
	}
}
