package models

import (
	"fmt"
	"strconv"
	"strings"
)

type Instruction struct {
	Id uint64
	ClassroomId uint64
	Text string
	PrevId uint64
}

func (instruction Instruction) String() string {
	var b strings.Builder

	b.WriteString(strconv.FormatUint(instruction.Id, 10))
	b.WriteString(": ")
	fmt.Fprintf(&b, "\"%s\"", instruction.Text)

	return b.String()
}

