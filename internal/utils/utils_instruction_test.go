package utils

import (
	"github.com/ozoncp/ocp-instruction-api/internal/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBatchInstructionSlice(t *testing.T) {
	_, err := BatchInstructionSlice(make([]models.Instruction, 0), 10)
	assert.NotNil(t, err)

	_, err = BatchInstructionSlice(make([]models.Instruction, 10), 0)
	assert.NotNil(t, err)

	sl := []models.Instruction{
		{Id: 1, ClassroomId: 15, Text: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.", PrevId: 0},
		{Id: 2, ClassroomId: 14, Text: "Maecenas imperdiet risus lectus, id ultrices odio gravida vitae.", PrevId: 0},
		{Id: 3, ClassroomId: 15, Text: "Aenean non orci eget lectus placerat porta non eu ligula.", PrevId: 0},
		{Id: 4, ClassroomId: 13, Text: "Ut sollicitudin malesuada mauris non pretium.", PrevId: 0},
		{Id: 5, ClassroomId: 1, Text: "Ut varius ligula metus, a volutpat leo porttitor et.", PrevId: 0},
		{Id: 6, ClassroomId: 5, Text: "Quisque ut porta libero.", PrevId: 0},
		{Id: 7, ClassroomId: 15, Text: "Curabitur sodales, nunc bibendum maximus faucibus, lectus erat fringilla nulla, quis tempor arcu magna ", PrevId: 0},
		{Id: 8, ClassroomId: 65, Text: "vel diam.", PrevId: 0},
		{Id: 9, ClassroomId: 2, Text: "Vestibulum posuere elit turpis, nec blandit ipsum sollicitudin imperdiet.", PrevId: 0},
		{Id: 10, ClassroomId: 6, Text: "Nam et consectetur enim.", PrevId: 0},
		{Id: 11, ClassroomId: 9, Text: "Mauris at rutrum enim.", PrevId: 0},
	}

	ret, err := BatchInstructionSlice(sl, 7)
	assert.Nil(t, err)

	exptd := [][]models.Instruction{
		{
			{Id: 1, ClassroomId: 15, Text: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.", PrevId: 0},
			{Id: 2, ClassroomId: 14, Text: "Maecenas imperdiet risus lectus, id ultrices odio gravida vitae.", PrevId: 0},
			{Id: 3, ClassroomId: 15, Text: "Aenean non orci eget lectus placerat porta non eu ligula.", PrevId: 0},
			{Id: 4, ClassroomId: 13, Text: "Ut sollicitudin malesuada mauris non pretium.", PrevId: 0},
			{Id: 5, ClassroomId: 1, Text: "Ut varius ligula metus, a volutpat leo porttitor et.", PrevId: 0},
			{Id: 6, ClassroomId: 5, Text: "Quisque ut porta libero.", PrevId: 0},
			{Id: 7, ClassroomId: 15, Text: "Curabitur sodales, nunc bibendum maximus faucibus, lectus erat fringilla nulla, quis tempor arcu magna ", PrevId: 0},
		}, {
			{Id: 8, ClassroomId: 65, Text: "vel diam.", PrevId: 0},
			{Id: 9, ClassroomId: 2, Text: "Vestibulum posuere elit turpis, nec blandit ipsum sollicitudin imperdiet.", PrevId: 0},
			{Id: 10, ClassroomId: 6, Text: "Nam et consectetur enim.", PrevId: 0},
			{Id: 11, ClassroomId: 9, Text: "Mauris at rutrum enim.", PrevId: 0},
		},
	}

	assert.Equal(t, exptd, ret)
}

func TestSlice2Map(t *testing.T) {
	sl := []models.Instruction{
		{Id: 1, ClassroomId: 15, Text: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.", PrevId: 0},
		{Id: 1, ClassroomId: 14, Text: "Maecenas imperdiet risus lectus, id ultrices odio gravida vitae.", PrevId: 0},
	}

	_, err := Slice2Map(sl)
	assert.NotNil(t, err)



	sl = []models.Instruction{
		{Id: 1, ClassroomId: 15, Text: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.", PrevId: 0},
		{Id: 2, ClassroomId: 14, Text: "Maecenas imperdiet risus lectus, id ultrices odio gravida vitae.", PrevId: 0},
		{Id: 3, ClassroomId: 15, Text: "Aenean non orci eget lectus placerat porta non eu ligula.", PrevId: 0},
	}

	ret, err := Slice2Map(sl)
	assert.Nil(t, err)

	assert.Equal(t, ret, map[uint64]models.Instruction{
		1: {Id: 1, ClassroomId: 15, Text: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.", PrevId: 0},
		2: {Id: 2, ClassroomId: 14, Text: "Maecenas imperdiet risus lectus, id ultrices odio gravida vitae.", PrevId: 0},
		3: {Id: 3, ClassroomId: 15, Text: "Aenean non orci eget lectus placerat porta non eu ligula.", PrevId: 0},
	})
}