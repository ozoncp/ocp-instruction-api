package repoService

import (
	"context"
	"github.com/ozoncp/ocp-instruction-api/internal/models"
	"github.com/ozoncp/ocp-instruction-api/internal/repo"
)

type IRepoService interface {
	Add(ctx context.Context, entities []models.Instruction) error
	List(ctx context.Context, limit, offset uint64) ([]models.Instruction, error)
	Describe(ctx context.Context, id uint64) (*models.Instruction, error)
	Remove(ctx context.Context, id uint64) error
	Update(ctx context.Context, entity models.Instruction) error
}

type RepoService struct {
	r repo.Repo
}

func (r *RepoService) Add(ctx context.Context, entities []models.Instruction) error {
	return r.r.Add(ctx, entities)
}
func (r *RepoService) List(ctx context.Context, limit, offset uint64) ([]models.Instruction, error) {
	return r.r.List(ctx, limit, offset)
}
func (r *RepoService) Describe(ctx context.Context, id uint64) (*models.Instruction, error) {
	return r.r.Describe(ctx, id)
}
func (r *RepoService) Remove(ctx context.Context, id uint64) error {
	return r.r.Remove(ctx, id)
}
func (r *RepoService) Update(ctx context.Context, entity models.Instruction) error {
	return r.r.Update(ctx, entity)
}

func BuildRequestService() IRepoService {
	return NewRequestService(repo.NewRepo())
}

func NewRequestService(r repo.Repo) *RepoService {
	return &RepoService{
		r: r,
	}
}
