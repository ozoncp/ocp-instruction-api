package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInstruction_String(t *testing.T) {
	obj := Instruction{
		Id:          854,
		ClassroomId: 43,
		Text:        "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
		PrevId:      0,
	}

	s := obj.String()

	assert.Equal(t, s, "854: \"Lorem ipsum dolor sit amet, consectetur adipiscing elit\"")
}
