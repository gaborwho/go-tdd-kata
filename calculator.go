package calculator

import (
	"strconv"
	"strings"
)

func Add(numbers string) (int, error) {
	replaced := strings.Replace(numbers, "\n", ",", -1)
	splitNumbers := strings.Split(replaced, ",")

	accumulator := 0
	for _, number := range splitNumbers {
		first, err := strconv.Atoi(number)
		if err != nil {
			return -1, err
		}
		accumulator += first
	}

	return accumulator, nil
}
