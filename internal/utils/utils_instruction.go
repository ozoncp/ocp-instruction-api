package utils

import (
	"errors"
	"github.com/ozoncp/ocp-instruction-api/internal/models"
)

func BatchInstructionSlice(input []models.Instruction, chankSize int) ([][]models.Instruction, error) {
	if chankSize < 1 {
		return nil, errors.New("incorrect chank size")
	}

	inputSize := len(input)
	if inputSize < 1 {
		return nil, errors.New("incorrect input slice size")
	}

	res := make([][]models.Instruction, int(inputSize/chankSize)+1)
	var i, j int
	for ; i < inputSize-chankSize; i += chankSize {
		res[j] = input[i : i+chankSize]
		j++
	}
	res[j] = input[i:inputSize]

	return res, nil
}

func Slice2Map (source []models.Instruction) (map[uint64]models.Instruction, error) {
	res := make(map[uint64]models.Instruction, len(source))
	for _, value := range source {
		if _, ok := res[value.Id]; ok {
			return nil, errors.New("duplicate id's")
		}
		res[value.Id] = value
	}

	return res, nil
}
