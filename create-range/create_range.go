package main

func createRange(start int, end int) []int {
	// Your code here
	var slice []int
	for i := start; i <= end; i++ {
		slice = append(slice, i)
	}
	return slice
}
