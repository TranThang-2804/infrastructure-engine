package usecase

import (
	"context"
	"time"

	"github.com/TranThang-2804/infrastructure-engine/internal/domain"
	"github.com/TranThang-2804/infrastructure-engine/internal/utils"
)

type iacPipelineUsecase struct {
	iacPipelineRepository domain.IacPipelineRepository
	contextTimeout        time.Duration
}

func NewIacPipelineUsecase(iacPipelineRepository domain.IacPipelineRepository) domain.IacPipelineUsecase {
	return &iacPipelineUsecase{
		iacPipelineRepository: iacPipelineRepository,
		contextTimeout:        utils.GetContextTimeout(),
	}
}

func (iu *iacPipelineUsecase) Trigger(c context.Context) (domain.IacPipeline, error) {
	ctx, cancel := context.WithTimeout(c, iu.contextTimeout)
	defer cancel()
	return iu.iacPipelineRepository.Trigger(ctx)
}

func (iu *iacPipelineUsecase) GetPipelineOutputByUrl(c context.Context) (domain.IacPipelineOutput, error) {
	ctx, cancel := context.WithTimeout(c, iu.contextTimeout)
	defer cancel()
	return iu.iacPipelineRepository.GetPipelineOutputByUrl(ctx)
}
