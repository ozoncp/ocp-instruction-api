package flusher

import (
	"github.com/ozoncp/ocp-instruction-api/internal/models"
	"github.com/ozoncp/ocp-instruction-api/internal/repo"
	"github.com/ozoncp/ocp-instruction-api/internal/utils"
)

// Flusher - interface for flush data to the storage
type Flusher interface {
	Flush(entities []models.Instruction) ([]models.Instruction, error)
}

// NewFlusher возвращает Flusher с поддержкой батчевого сохранения
func NewFlusher(chunkSize int, repo repo.Repo) Flusher {
	return &flusher{
		chunkSize: chunkSize,
		repo:  repo,
	}
}

type flusher struct {
	chunkSize int
	repo  repo.Repo
}

func (f *flusher) Flush(data []models.Instruction) ([]models.Instruction, error) {
	chunks, err := utils.BatchInstructionSlice(data, f.chunkSize)
	if err != nil {
		return data, err
	}

	for i := 0; i < len(chunks); i++ {
		err := f.repo.Add(chunks[i])
		if err != nil {
			return data[(i*f.chunkSize): ], err
		}
	}

	return make([]models.Instruction, 0), nil
}
