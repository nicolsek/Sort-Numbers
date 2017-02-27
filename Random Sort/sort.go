package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

// main ... The main function?
func main() {
	nums := randomSort(getArgs())
	fmt.Printf("Sorted: %v\n", nums)
}

// randomSort ... Randomly sorts the list until it's done.
func randomSort(nums []int) []int {
	//Switches two positions.
	swap := func(array []int, pos0 int, pos1 int) {
		val0 := array[pos0]
		array[pos0] = array[pos1]
		array[pos1] = val0
	}

	for !checkSort(nums) {
		swap(nums, rand.Intn(len(nums)), rand.Intn(len(nums)))
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
