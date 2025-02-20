package main

import (
	"bufio"
	"fmt"
	sort_utils "lab2/sorting"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("QuickSortDuplicateKeys: \n")
	check_quicks_sort[float64](sort_utils.QuickSortDuplicateKeys, "sr")

	fmt.Println("QuickSortMedian: \n")
	check_quicks_sort[float64](sort_utils.QuickSortMedian, "sr")

	check_sort_select("sr")
	check_sort_merge("sr")

}

func check_quicks_sort[T ~int | ~float64](f func([]T) (uint32, []T), filename string) {
	sliece := make([]T, 0, 100)
	sliece = get_arr(sliece, filename)
	var count uint32 = 0
	var res_time time.Duration

	res_time, _ = sort_utils.TimeFunction(f, sliece, &count)
	fmt.Println("Time :", res_time)

	sliece = make([]T, 0, 100)
	sliece = get_arr(sliece, filename)
	res_time, _ = sort_utils.TimeFunction(f, sliece, &count)

	fmt.Println("Time :", res_time)
	sliece = make([]T, 0, 100)
	sliece = get_arr(sliece, filename)

	res_time, _ = sort_utils.TimeFunction(f, sliece, &count)

	fmt.Println("Time :", res_time)
	sliece = make([]T, 0, 100)
	sliece = get_arr(sliece, filename)
	res_time, _ = sort_utils.TimeFunction(f, sliece, &count)

	fmt.Println("Time :", res_time)
	fmt.Println("Count comparison: ", count)
	fmt.Println("check_seq:", check_seq(sliece))
	fmt.Println("check_seq_rev:", check_seq_rev(sliece))
	fmt.Printf("\n\n\n")
}

func check_sort_select(filename string) {
	var count uint32 = 0
	sliece := make([]float64, 0, 100)
	sliece = get_arr(sliece, filename)
	// print_arr(sliece)
	fmt.Println("SelectionSort:")

	fmt.Println("Time :", sort_utils.TimeFunction2(sort_utils.SelectionSort, sliece, &count))
	sliece = make([]float64, 0, 100)
	sliece = get_arr(sliece, filename)

	fmt.Println("Time :", sort_utils.TimeFunction2(sort_utils.SelectionSort, sliece, &count))
	sliece = make([]float64, 0, 100)
	sliece = get_arr(sliece, filename)
	fmt.Println("Time :", sort_utils.TimeFunction2(sort_utils.SelectionSort, sliece, &count))
	fmt.Println("Count comparison: ", count)
	fmt.Println("check_seq:", check_seq(sliece))
	fmt.Println("check_seq_rev:", check_seq_rev(sliece))
	fmt.Println("\n\n")
	// print_arr(sliece)

}

func check_sort_merge(filename string) {
	fmt.Println("Merge_sort:")
	var count uint32 = 0
	var res_time time.Duration
	sliece := make([]float64, 0, 100)
	sliece = get_arr(sliece, filename)

	_, _ = sort_utils.TimeFunction(sort_utils.Merge_sort, sliece, &count)

	sliece = make([]float64, 0, 100)
	sliece = get_arr(sliece, filename)

	res_time, sliece = sort_utils.TimeFunction(sort_utils.Merge_sort, sliece, &count)

	fmt.Println("Time :", res_time)
	// print_arr(sliece)
	sliece = make([]float64, 0, 100)
	sliece = get_arr(sliece, filename)
	// print_arr(sliece)
	res_time, sliece = sort_utils.TimeFunction(sort_utils.Merge_sort, sliece, &count)
	fmt.Println("Time :", res_time)

	sliece = make([]float64, 0, 100)
	sliece = get_arr(sliece, filename)
	res_time, sliece = sort_utils.TimeFunction(sort_utils.Merge_sort, sliece, &count)
	fmt.Println("Time :", res_time)

	fmt.Println("Count comparison: ", count)
	fmt.Println("check_seq:", check_seq(sliece))
	fmt.Println("check_seq_rev:", check_seq_rev(sliece))
	fmt.Println("\n\n")
}

func print_arr[T ~int | ~float64](s []T) {
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}
func get_arr[T ~int | ~float64](s []T, filename string) []T {

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, " ")
		for _, numberStr := range numbers {
			numberStr = strings.TrimSpace(numberStr)
			if numberStr == "" {
				continue
			}

			number, err := strconv.ParseFloat(numberStr, 64)
			if err != nil {
				fmt.Printf("Error parsing float '%s': %s\n", numberStr, err)
				continue
			}
			s = append(s, T(number))
			// fmt.Printf("%f ", number)
		}
		// fmt.Printf("\n")

	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error scanning file: %s", err)
	}
	return s
}

func check_seq[T ~int | ~float64](s []T) bool {
	result := true
	if len(s) <= 1 {
		return result
	}

	for i := 1; i < len(s); i++ {
		if s[i] < s[i-1] {
			// fmt.Printf("%g %.2f\n", s[i], s[i-1])
			result = false
			break
		}
	}
	return result
}
func check_seq_rev[T ~int | ~float64](s []T) bool {
	result := true
	if len(s) <= 1 {
		return result
	}
	for i := 1; i < len(s); i++ {
		if s[i] > s[i-1] {
			result = false
			break
		}
	}
	return result
}
