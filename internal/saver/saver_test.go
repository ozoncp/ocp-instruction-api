package saver

import (
	"bytes"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/ozoncp/ocp-instruction-api/internal/mocks"
	"github.com/ozoncp/ocp-instruction-api/internal/models"
	"github.com/ozoncp/ocp-instruction-api/internal/utils"
	"log"
	"strings"
	"testing"
	"time"
)

func TestSaver_Save(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFlusher := mocks.NewMockFlusher(ctrl)
	s := NewSaver(10, mockFlusher, (500 * time.Millisecond))

	mockFlusher.EXPECT().
		Flush(gomock.Any()).
		Times(2).
		Return(make([]models.Instruction, 0), nil)

	data := utils.GenerateInstructionSlice(1)
	s.Save(data[0])
	time.Sleep(750 * time.Millisecond)

	data = utils.GenerateInstructionSlice(1)
	s.Save(data[0])
	s.Close()
	time.Sleep(200 * time.Millisecond)
}

func TestSaver_SaveError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var logBuff bytes.Buffer
	log.SetOutput(&logBuff)

	mockFlusher := mocks.NewMockFlusher(ctrl)
	s := NewSaver(10, mockFlusher, (500 * time.Millisecond))

	mockFlusher.EXPECT().
		Flush(gomock.Any()).
		Times(1).
		Return(make([]models.Instruction, 0), errors.New("some error"))

	data := utils.GenerateInstructionSlice(1)
	s.Save(data[0])

	time.Sleep(750 * time.Millisecond)

	ind := strings.Index(logBuff.String(), "saver dump error: some error")

	if ind == -1 {
		t.Error("Want saver dump error")
	}
}
