package main

func SumOf1Darr(nums []int) []int {
	sums := []int{}
	prev := 0

	for _, val := range nums {
		sums = append(sums, val+prev)
		prev = val + prev
	}
	return sums
}
