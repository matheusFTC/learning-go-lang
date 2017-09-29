package main

import (
	"fmt"
	"os"
	"strconv"
)

func sort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	} else {
		numscopy := make([]int, len(nums))
		copy(numscopy, nums)

		indexPivot := len(numscopy) / 2
		numPivot := numscopy[indexPivot]

		numscopy = append(numscopy[:indexPivot], numscopy[indexPivot + 1:]...)

		minors, larger := partition(numscopy, numPivot)

		return append(append(sort(minors), numPivot), sort(larger)...)
	}
}

func partition(nums []int, numPivot int) (minors []int, larger []int) {
	for _, n := range nums {
		if n <= numPivot {
			minors = append(minors, n);
		} else {
			larger = append(larger, n);
		}
	}

	return minors, larger
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Enter the values ​​for sorting.")
		
		os.Exit(1)
	} else {
		values := os.Args[1:]
		nums := make([] int, len(values))

		for i, n := range values {
			num, err := strconv.Atoi(n)

			if err != nil {
				fmt.Printf("%s is a invalid number.\n", n)
				
				os.Exit(1)
			} else {
				nums[i] = num
			}
		}

		fmt.Println(sort(nums))

		os.Exit(0)
	}
}