package main

import "fmt"

const (
	minimum = "minimum"

	average = "average"

	maximum = "maximum"
)

func main() {
	fmt.Println("\nExercicio 4\n-----------------")

	minFunc, _ := operation(minimum)

	averageFunc, _ := operation(average)

	maxFunc, _ := operation(maximum)

	minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	fmt.Printf("Minimum value: %.2f\n", minValue)

	averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	fmt.Printf("Average value: %.2f\n", averageValue)

	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
	fmt.Printf("Maximum value: %.2f\n", maxValue)
}

func operation(op string) (func(...int) float64, error) {
	switch op {
	case "minimum":
		return func(nums ...int) float64 {
			if len(nums) == 0 {
				return 0
			}
			min := nums[0]
			for _, num := range nums {
				if num < min {
					min = num
				}
			}
			return float64(min)
		}, nil
	case "average":
		return func(nums ...int) float64 {
			if len(nums) == 0 {
				return 0
			}
			sum := 0
			for _, num := range nums {
				sum += num
			}
			return float64(sum) / float64(len(nums))
		}, nil
	case "maximum":
		return func(nums ...int) float64 {
			if len(nums) == 0 {
				return 0
			}
			max := 0
			for _, num := range nums {
				if num > max {
					max = num
				}
			}
			return float64(max)
		}, nil
	default:
		return nil, fmt.Errorf("operation %s not supported", op)
	}
}
