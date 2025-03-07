package main

import (
	"bufio"
	"finder/finders"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	arr := make([]float64, 0, 100)
	arr = get_arr(arr, "so")
	// count_comp := 0
	// fmt.Println(finders.Sequential_search(arr, 12, &count_comp))
	// fmt.Println(finders.Sequential_search(arr, 44, &count_comp))
	// fmt.Println(finders.Sequential_search(arr, 18, &count_comp))
	// fmt.Println(finders.Sequential_search(arr, 786, &count_comp))
	// fmt.Println(finders.Sequential_search(arr, 1000, &count_comp))

	// fmt.Printf("count compare: %d\n", (count_comp+5)/5)
	// count_comp = 0

	// fmt.Println(finders.JumpSearch(arr, 12, int(math.Sqrt(float64(len(arr)))), &count_comp))
	// fmt.Println(finders.JumpSearch(arr, 44, int(math.Sqrt(float64(len(arr)))), &count_comp))
	// fmt.Println(finders.JumpSearch(arr, 18, int(math.Sqrt(float64(len(arr)))), &count_comp))
	// fmt.Println(finders.JumpSearch(arr, 786, int(math.Sqrt(float64(len(arr)))), &count_comp))
	// fmt.Println(finders.JumpSearch(arr, 1000, int(math.Sqrt(float64(len(arr)))), &count_comp))

	// fmt.Printf("count compare: %d\n", count_comp/5)
	// count_comp = 0

	// fmt.Println(finders.TwoJumpSearch(arr, 12, int(math.Pow(float64(len(arr)), 2.0/3.0)), int(math.Pow(float64(len(arr)), 1.0/3.0)), &count_comp))
	// fmt.Println(finders.TwoJumpSearch(arr, 44, int(math.Pow(float64(len(arr)), 2.0/3.0)), int(math.Pow(float64(len(arr)), 1.0/3.0)), &count_comp))
	// fmt.Println(finders.TwoJumpSearch(arr, 18, int(math.Pow(float64(len(arr)), 2.0/3.0)), int(math.Pow(float64(len(arr)), 1.0/3.0)), &count_comp))
	// fmt.Println(finders.TwoJumpSearch(arr, 786, int(math.Pow(float64(len(arr)), 2.0/3.0)), int(math.Pow(float64(len(arr)), 1.0/3.0)), &count_comp))
	// fmt.Println(finders.TwoJumpSearch(arr, 1000, int(math.Pow(float64(len(arr)), 2.0/3.0)), int(math.Pow(float64(len(arr)), 1.0/3.0)), &count_comp))

	// fmt.Printf("count compare: %d\n", count_comp/5)
	// count_comp = 0

	// fmt.Println(finders.BinarySearch(arr, 12, &count_comp))
	// fmt.Println(finders.BinarySearch(arr, 44, &count_comp))
	// fmt.Println(finders.BinarySearch(arr, 18, &count_comp))
	// fmt.Println(finders.BinarySearch(arr, 786, &count_comp))
	// fmt.Println(finders.BinarySearch(arr, 1000, &count_comp))

	// fmt.Printf("count compare: %d\n", count_comp/5)
	// count_comp = 0

	//////////////////////////////////////////////////////////////

	// count_comp := 0
	// fmt.Println(finders.Sequential_search(arr, 11, &count_comp))
	// fmt.Println(finders.Sequential_search(arr, 43, &count_comp))
	// fmt.Println(finders.Sequential_search(arr, 17, &count_comp))
	// fmt.Println(finders.Sequential_search(arr, 785, &count_comp))
	// fmt.Println(finders.Sequential_search(arr, 999, &count_comp))

	// fmt.Printf("count compare: %d\n", (count_comp+5)/5)
	// count_comp = 0

	// fmt.Println(finders.JumpSearch(arr, 11, int(math.Sqrt(float64(len(arr)))), &count_comp))
	// fmt.Println(finders.JumpSearch(arr, 43, int(math.Sqrt(float64(len(arr)))), &count_comp))
	// fmt.Println(finders.JumpSearch(arr, 17, int(math.Sqrt(float64(len(arr)))), &count_comp))
	// fmt.Println(finders.JumpSearch(arr, 785, int(math.Sqrt(float64(len(arr)))), &count_comp))
	// fmt.Println(finders.JumpSearch(arr, 999, int(math.Sqrt(float64(len(arr)))), &count_comp))

	// fmt.Printf("count compare: %d\n", count_comp/5)
	// count_comp = 0

	// fmt.Println(finders.TwoJumpSearch(arr, 11, int(math.Pow(float64(len(arr)), 2.0/3.0)), int(math.Pow(float64(len(arr)), 1.0/3.0)), &count_comp))
	// fmt.Println(finders.TwoJumpSearch(arr, 43, int(math.Pow(float64(len(arr)), 2.0/3.0)), int(math.Pow(float64(len(arr)), 1.0/3.0)), &count_comp))
	// fmt.Println(finders.TwoJumpSearch(arr, 17, int(math.Pow(float64(len(arr)), 2.0/3.0)), int(math.Pow(float64(len(arr)), 1.0/3.0)), &count_comp))
	// fmt.Println(finders.TwoJumpSearch(arr, 785, int(math.Pow(float64(len(arr)), 2.0/3.0)), int(math.Pow(float64(len(arr)), 1.0/3.0)), &count_comp))
	// fmt.Println(finders.TwoJumpSearch(arr, 999, int(math.Pow(float64(len(arr)), 2.0/3.0)), int(math.Pow(float64(len(arr)), 1.0/3.0)), &count_comp))

	// fmt.Printf("count compare: %d\n", count_comp/5)
	// count_comp = 0

	// fmt.Println(finders.BinarySearch(arr, 11, &count_comp))
	// fmt.Println(finders.BinarySearch(arr, 43, &count_comp))
	// fmt.Println(finders.BinarySearch(arr, 17, &count_comp))
	// fmt.Println(finders.BinarySearch(arr, 785, &count_comp))
	// fmt.Println(finders.BinarySearch(arr, 999, &count_comp))

	// fmt.Printf("count compare: %d\n", count_comp/5)
	// count_comp = 0

	////////////////////////

	count_comp := 0
	for i := 0; i <= 11; i++ {

		fmt.Println(finders.JumpSearch(arr, 12, int(math.Pow(2, float64(i))), &count_comp))
		fmt.Println(finders.JumpSearch(arr, 424, int(math.Pow(2, float64(i))), &count_comp))
		fmt.Println(finders.JumpSearch(arr, 952, int(math.Pow(2, float64(i))), &count_comp))
		fmt.Println(finders.JumpSearch(arr, 1086, int(math.Pow(2, float64(i))), &count_comp))
		fmt.Println(finders.JumpSearch(arr, 2000, int(math.Pow(2, float64(i))), &count_comp))

		fmt.Printf("count compare: %d %d\n", count_comp/5, i)
		count_comp = 0

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
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error scanning file: %s", err)
	}
	return s
}
