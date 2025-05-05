package repository

import (
	"context"

	"github.com/TranThang-2804/infrastructure-engine/internal/domain"
	"github.com/TranThang-2804/infrastructure-engine/internal/infrastructure/git"
)

type iacPipelineRepository struct {
	gitStore git.GitStore
}

func NewIacPipelineRepository(gitStore git.GitStore) domain.IacPipelineRepository {
	return &iacPipelineRepository{
		gitStore: gitStore,
	}
}

func (ir *iacPipelineRepository) Trigger(c context.Context, iacPipeline domain.IacPipeline) (domain.IacPipeline, error) {
	return iacPipeline, nil
}

func (ir *iacPipelineRepository) GetPipelineOutputByUrl(c context.Context, iacPipeline domain.IacPipeline) (domain.IacPipelineOutput, error) {
	return domain.IacPipelineOutput{}, nil
}
