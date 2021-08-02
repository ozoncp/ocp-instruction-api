package repo

import "github.com/ozoncp/ocp-instruction-api/internal/models"

// Repo - storage interface for models.Instruction
type Repo interface {
	Add(entities []models.Instruction) error
	List(limit, offset uint64) ([]models.Instruction, error)
	Describe(entityId uint64) (*models.Instruction, error)
}

