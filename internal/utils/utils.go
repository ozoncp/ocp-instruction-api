package utils

import (
	"errors"
)

func BatchIntSlice(input []int, chankSize int) ([][]int, error) {
	if chankSize < 1 {
		return nil, errors.New("incorrect chank size")
	}

	inputSize := len(input)
	if inputSize < 1 {
		return nil, errors.New("incorrect input slice size")
	}

	res := make([][]int, int(inputSize/chankSize)+1)
	var i, j int
	for ; i < inputSize-chankSize; i += chankSize {
		res[j] = input[i : i+chankSize]
		j++
	}
	res[j] = input[i:inputSize]

	return res, nil
}

func SwapIntMap(input map[int]int) (map[int]int, error) {
	res := make(map[int]int, len(input))
	for key, value := range input {
		if _, ok := res[value]; ok {
			return nil, errors.New("duplicate values")
		}
		res[value] = key
	}

	return res, nil
}

func FilterIntSlice(input []int, values ...int) []int {
	res := make([]int, 0, len(input))

InputLoop:
	for _, inputVal := range input {
		for _, v := range values {
			if inputVal == v {
				continue InputLoop
			}
		}
		res = append(res, inputVal)
	}

	return res
}
