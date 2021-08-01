package utils

import (
	"github.com/ozoncp/ocp-instruction-api/internal/models"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func generateInstructionSlice(count int) []models.Instruction {
	ret := make([]models.Instruction, 0, count)
	for i := 0; i < count; i++ {
		ret = append(ret, models.Instruction{Id: rand.Uint64(), ClassroomId: rand.Uint64(), PrevId:rand.Uint64(), Text: randString(rand.Intn(64))})
	}

	return ret
}



func TestBatchInstructionSlice(t *testing.T) {
	_, err := BatchInstructionSlice(make([]models.Instruction, 10), 0)
	assert.NotNil(t, err)

	ret, err := BatchInstructionSlice(make([]models.Instruction, 0), 10)
	assert.Nil(t, err)
	assert.Equal(t, ret, make([][]models.Instruction, 0))


	sl_1 := generateInstructionSlice(1)
	ret, err = BatchInstructionSlice(sl_1, 10)
	assert.Nil(t, err)
	assert.Equal(t, ret, [][]models.Instruction{sl_1})

	sl_2 := generateInstructionSlice(2)
	ret, err = BatchInstructionSlice(sl_2, 3)
	assert.Nil(t, err)
	assert.Equal(t, ret, [][]models.Instruction{sl_2})

	sl_3 := generateInstructionSlice(3)
	ret, err = BatchInstructionSlice(sl_3, 3)
	assert.Nil(t, err)
	assert.Equal(t, ret, [][]models.Instruction{sl_3})

	sl1 := generateInstructionSlice(3)
	sl2 := generateInstructionSlice(1)
	sl := append(sl1, sl2...)
	exptd := [][]models.Instruction{sl1, sl2}

	ret, err = BatchInstructionSlice(sl, 3)
	assert.Nil(t, err)
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
