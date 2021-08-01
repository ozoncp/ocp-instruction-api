package utils

import (
	"github.com/ozoncp/ocp-instruction-api/internal/models"
	"math/rand"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func GenerateInstructionSlice(count int) []models.Instruction {
	ret := make([]models.Instruction, 0, count)
	for i := 0; i < count; i++ {
		ret = append(ret, models.Instruction{Id: rand.Uint64(), ClassroomId: rand.Uint64(), PrevId:rand.Uint64(), Text: randString(rand.Intn(64))})
	}

	return ret
}
