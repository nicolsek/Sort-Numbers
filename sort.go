package main

import (
	"fmt"
	"os"
	"strconv"
)

/*

   --  Sort.go --

   The goal of this is to take a list of numbers in the terminal and sorts them from low to high. It does this using Selection Sort.
   In essence what it's going to do is take an array of x numbers and with those x numbers it recursively checks the positions of the numbers at
   a position and switches if it should be higher or lower. It does this until the list is completely sorted.

*/

// main ... The main function?
func main() {
	nums := orderNumbers(getArgs())
	fmt.Printf("Ordered Numbers: %v\n", nums)
}

// orderNumbers ... Uses the Selection Sort algorithm to sort an array of numbers from lowest to highest.
func orderNumbers(nums []int) []int {
	//Creating an anonymous function to swap two position of an int array.
	swap := func(swapNums []int, pos0 int, pos1 int) {
		//Store the value to access later.
		value0 := swapNums[pos0]
		swapNums[pos0] = swapNums[pos1]
		swapNums[pos1] = value0
	}

	//While it isn't sorted, sort.
	for !checkSort(nums) {
		for i := 1; i < len(nums) && i+1 < len(nums)+1; i++ {
			if nums[i-1] > nums[i] {
				swap(nums, i-1, i)
			}
		}
	}

	return nums
}

// checkSort ... Runs a for loop through the size of an array and checks to see if it's sorted.
func checkSort(nums []int) bool {
	isSorted := true

	//Using a for loop it checks the array to make sure that [i + 1] < [i] of the list.
	for i := 1; i < len(nums) && i+1 < len(nums); i++ {
		if isSorted {
			if nums[i-1] > nums[i] {
				isSorted = false
			}
		}
	}

	return isSorted
}

// getArgs ... Used the os.Args[] to get the numbers from the command line.
func getArgs() []int {
	var args []int

	args = checkArgs(os.Args[1:])

	return args
}

// checkArgs ... Using strconv.Atoi it checks to see if the list of numbers is a number if it returns a letter or something similar then panic and write the usage.
func checkArgs([]string) []int {

	// var args []int
	args := make([]int, len(os.Args[1:]))

	for i, s := range os.Args[1:] {
		v, err := strconv.Atoi(s)

		if err != nil {
			fmt.Printf("err: %v\n", err.Error())
			panic(err)
		} else {
			args[i] = v
		}
	}

	return args
}
