package damm

import (
	"errors"
	"strconv"
	"strings"
)

func encode(s string) (int, error) {
	// 10 x 10 matrix
	matrix := [][]int{
		[]int{0, 3, 1, 7, 5, 9, 8, 6, 4, 2},
		[]int{7, 0, 9, 2, 1, 5, 4, 8, 6, 3},
		[]int{4, 2, 0, 6, 8, 7, 1, 3, 5, 9},
		[]int{1, 7, 5, 0, 9, 8, 3, 4, 2, 6},
		[]int{6, 1, 2, 3, 0, 4, 5, 9, 7, 8},
		[]int{3, 6, 7, 4, 2, 0, 9, 5, 8, 1},
		[]int{5, 8, 6, 9, 7, 2, 0, 1, 3, 4},
		[]int{8, 9, 4, 5, 3, 6, 2, 0, 1, 7},
		[]int{9, 4, 3, 8, 6, 1, 7, 2, 0, 5},
		[]int{2, 5, 8, 1, 4, 3, 6, 7, 9, 0},
	}

	interim := 0

	for _, v := range strings.Split(s, "") {
		row, err := strconv.Atoi(v)
		if err != nil {
			return 0, err
		}
		interim = matrix[interim][row]
	}

	return interim, nil
}

func Append(s string) (string, error) {
	if s == "" {
		return s, errors.New("Empty string not allowed.")
	}

	_, err := strconv.Atoi(s[len(s)-1:])
	if err != nil {
		return s, errors.New("Input must be numbers [0-9].")
	}

	n, err := encode(s[0:len(s)])
	checkdigit := strconv.Itoa(n)

	return strings.Join([]string{s, checkdigit}, ""), err // faster than `+`
}

func Verify(s string) (bool, error) {
	if s == "" {
		return false, errors.New("Empty string not allowed.")
	}

	last_word, err := strconv.Atoi(s[len(s)-1:])
	if err != nil {
		return false, errors.New("Input must be numbers [0-9].")
	}

	validator, err := encode(s[0 : len(s)-1])
	if err != nil {
		return false, errors.New("Input must be numbers [0-9].")
	}

	if validator != last_word {
		return false, nil
	}

	return true, nil
}
