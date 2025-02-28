package sort_utils

import "time"

func Merge_sort[T ~int | ~float64](s []T) (uint32, []T) {
	var count_op uint32 = 0
	if len(s) <= 1 {
		count_op++
		return count_op, s
	}
	middle := len(s) / 2
	left_s := s[:middle]
	right_s := s[middle:]
	countLeft, sortedLeft := Merge_sort(left_s)
	countRight, sortedRight := Merge_sort(right_s)
	mergedSlice, countMerge := merging(sortedLeft, sortedRight)
	count_op += countLeft + countRight + countMerge
	return count_op, mergedSlice
}
func merging[T ~int | ~float64](left_s []T, right_s []T) ([]T, uint32) {
	var count_op uint32 = 0
	result := make([]T, len(left_s)+len(right_s))
	i := 0
	j := 0
	k := 0
	for i < len(left_s) && j < len(right_s) {

		if left_s[i] < right_s[j] {
			result[k] = left_s[i]
			i++
		} else {
			result[k] = right_s[j]
			j++
		}
		count_op++

		k++

	}
	for i < len(left_s) {
		count_op++
		result[k] = left_s[i]
		i++
		k++
	}

	for j < len(right_s) {
		count_op++
		result[k] = right_s[j]
		j++
		k++
	}
	return result, count_op
}
func SelectionSort[T ~int | ~float64](s []T) uint32 {
	n := len(s)
	var count_op uint32 = 0
	for i := 0; i < n-1; i++ {
		min_index := i
		for j := i + 1; j < n; j++ {
			if s[j] < s[min_index] {
				count_op++
				min_index = j
			}
		}
		if i != min_index {
			count_op++
			s[i], s[min_index] = s[min_index], s[i]
		}
	}
	return count_op
}

func TimeFunction[T ~int | ~float64](f func([]T) (uint32, []T), slice []T, count_op *uint32) (res_time float64, res []T) {
	startTime := time.Now()
	*count_op, res = f(slice)
	endTime := time.Now()
	res_time = float64(endTime.Sub(startTime)) / float64(time.Millisecond)
	return res_time, res
}

func TimeFunction2[T ~int | ~float64](f func([]T) uint32, slice []T, count_op *uint32) float64 {
	startTime := time.Now()
	*count_op = f(slice)
	endTime := time.Now()
	return float64(endTime.Sub(startTime)) / float64(time.Millisecond)
}

func QuickSortMedian[T ~int | ~float64](arr []T) (uint32, []T) {
	n := len(arr)
	var count uint32 = 0
	left := 0
	right := n - 1
	quickSortMedian(arr, &left, &right, &count)
	return count, arr
}

func quickSortMedian[T ~int | ~float64](arr []T, left, right *int, count *uint32) {
	if *left < *right {
		*count++
		pivotIndex := partition(arr, left, right, count)

		l := *left
		r := pivotIndex - 1
		quickSortMedian(arr, &l, &r, count)
		l = pivotIndex + 1
		r = *right
		quickSortMedian(arr, &l, &r, count)
	}
}

func partition[T ~int | ~float64](arr []T, left, right *int, count *uint32) int {
	medianIndex := medianOfThree(arr, *left, *right, count)

	arr[*right], arr[medianIndex] = arr[medianIndex], arr[*right]
	pivot := arr[*right]
	i := *left - 1

	for j := *left; j < *right; j++ {
		if arr[j] < pivot {
			*count++
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[*right] = arr[*right], arr[i+1]
	return i + 1
}

func medianOfThree[T ~int | ~float64](arr []T, left, right int, count *uint32) int {
	mid := left + (right-left)/2

	if arr[left] > arr[mid] {
		arr[left], arr[mid] = arr[mid], arr[left]
	}
	if arr[left] > arr[right] {
		arr[left], arr[right] = arr[right], arr[left]
	}
	if arr[mid] > arr[right] {
		arr[mid], arr[right] = arr[right], arr[mid]
	}
	*count += 3

	return mid
}

func QuickSortDuplicateKeys[T ~int | ~float64](arr []T) (uint32, []T) {
	n := len(arr)
	var count uint32 = 0
	left := 0
	right := n - 1
	quickSortDuplicateKeys(arr, &left, &right, &count)
	return count, arr
}

func quickSortDuplicateKeys[T ~int | ~float64](arr []T, left, right *int, count *uint32) {
	if *left < *right {
		*count++
		lt, gt := partition3Way(arr, left, right, count)

		l := *left
		r := lt - 1
		quickSortDuplicateKeys(arr, &l, &r, count)
		l = gt + 1
		r = *right
		quickSortDuplicateKeys(arr, &l, &r, count)
	}
}

func partition3Way[T ~int | ~float64](arr []T, left, right *int, count *uint32) (int, int) {
	pivot := arr[*left]
	lt := *left  // arr[left...lt-1] < pivot
	gt := *right // arr[gt+1...right] > pivot
	i := *left   // arr[lt...i-1] == pivot

	for i <= gt {
		if arr[i] < pivot {
			arr[lt], arr[i] = arr[i], arr[lt]
			lt++
			i++
			*count += 1
		} else if arr[i] > pivot {
			arr[i], arr[gt] = arr[gt], arr[i]
			gt--
			*count += 2
		} else {
			i++
			*count += 3
		}
	}

	return lt, gt
}
