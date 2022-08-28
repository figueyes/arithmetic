package check

import "sort"

func CheckArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	length := len(l)
	response := make([]bool, 0)

	for i := 0; i < length; i++ {
		sa := subarray(nums, l[i], r[i])
		cha := checkArithmetic(sa)
		response = append(response, cha)
	}

	return response
}

func subarray(nums []int, l, r int) (res []int) {
	for i := l; i <= r; i++ {
		res = append(res, nums[i])
	}
	sort.Ints(res)
	return res
}

func checkArithmetic(nums []int) (res bool) {
	previous := 0
	sequence := 0
	for k, v := range nums {
		if k == 0 {
			previous = v
			continue
		}
		if k == 1 {
			sequence = v - previous
			res = sequence == v-previous
			previous = v
			continue
		}
		res = sequence == v-previous
		previous = v
		if !res {
			return false
		}
	}
	return res
}
