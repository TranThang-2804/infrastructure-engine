package repository

import (
	"context"
	"encoding/json"

	"github.com/TranThang-2804/infrastructure-engine/internal/domain"
	"github.com/TranThang-2804/infrastructure-engine/internal/infrastructure/git"
	"github.com/TranThang-2804/infrastructure-engine/internal/shared/log"
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
	pipelineUrl, err := ir.gitStore.TriggerPipeline(
		"TranThang-2804",
		"platform-iac-resource",
		pipelinePayload,
	)
	if err != nil {
		log.Logger.Error("Error Triggering pipeline", "error", err)
		return domain.IacPipeline{}, nil
	}

	// Get Pipeline URL and attach it
	iacPipeline.URL = pipelineUrl
	return iacPipeline, nil
}

func (ir *iacPipelineRepository) GetPipelineOutputByUrl(c context.Context, iacPipeline domain.IacPipeline) (domain.IacPipelineOutput, error) {
	return domain.IacPipelineOutput{}, nil
}
