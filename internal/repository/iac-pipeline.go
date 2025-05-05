package repository

import (
	"context"
	"encoding/json"

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
	pipelineData := map[string]interface{}{
		"action":   iacPipeline.Action,
		"filepath": iacPipeline.Name,
	}
	pipelinePayload, err := json.Marshal(pipelineData)
	if err != nil {
		return domain.IacPipeline{}, err
	}

  // Trigger pipeline
	ir.gitStore.TriggerPipeline(
		"TranThang-2804",
		"platform-iac-resource",
		pipelinePayload,
	)
	return iacPipeline, nil
}

func (ir *iacPipelineRepository) GetPipelineOutputByUrl(c context.Context, iacPipeline domain.IacPipeline) (domain.IacPipelineOutput, error) {
	return domain.IacPipelineOutput{}, nil
}
