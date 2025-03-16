package finders

import (
	"math"
)

func Sequential_search(arr []float64, target float64, count_comp *int) int {
	result := -1
	for i := 0; i < len(arr); i++ {
		if math.Abs(arr[i]-target) < 0.000001 {
			result = i
			break
		}
		*count_comp++
	}
	return result
}

func JumpSearch(arr []float64, target float64, step1 int, count_comp *int) int {
	n := len(arr)
	if n == 0 {
		*count_comp++
		return -1
	}

	start := step1

	for arr[start] < target {
		start += step1

		if start >= n {
			*count_comp++
			return -1
		}
		*count_comp++
	}
	for i := start; i > 0; i-- {
		if arr[i] == target {
			*count_comp++
			return i
		}
		*count_comp++
	}
	return -1
}

func TwoJumpSearch(arr []float64, target float64, step1 int, step2 int, count_comp *int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}

	prev1 := 0
	for prev1 < n {
		nextJump := int(math.Min(float64(prev1+step1), float64(n)))
		if nextJump == n || arr[nextJump-1] >= target {
			*count_comp++
			break
		}
		*count_comp++
		prev1 += step1
	}

	if prev1 >= n {
		*count_comp++
		return -1
	}

	prev2 := prev1
	for prev2+step1 < n {
		nextJump := int(math.Min(float64(prev2+step2), float64(n)))
		if nextJump == n || arr[nextJump-1] >= target {
			*count_comp++
			break
		}
		*count_comp++
		prev2 += step2
	}

	for i := prev2; i < int(math.Min(float64(prev2+step2), float64(n))); i++ {
		if math.Abs(arr[i]-target) < 0.000001 {
			*count_comp++
			return i
		}
		*count_comp++
	}

	return -1
}

func BinarySearch(arr []float64, target float64, count_comp *int) int {
	low, high := 0, len(arr)-1
	result := -1
	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] == target {
			*count_comp++
			result = mid
			break
		} else if arr[mid] < target {
			*count_comp += 2
			low = mid + 1
		} else {
			*count_comp += 2
			high = mid - 1
		}
	}
	return result
}
