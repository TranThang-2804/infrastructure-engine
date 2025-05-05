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

func (iu *iacPipelineUsecase) Trigger(c context.Context, iacPipeline domain.IacPipeline) (string, error) {
	ctx, cancel := context.WithTimeout(c, iu.contextTimeout)
	defer cancel()
	return iu.iacPipelineRepository.Trigger(ctx, iacPipeline)
}

func (iu *iacPipelineUsecase) GetPipelineOutputByUrl(c context.Context, iacPipeline domain.IacPipeline) ([]byte, error) {
	ctx, cancel := context.WithTimeout(c, iu.contextTimeout)
	defer cancel()
	return iu.iacPipelineRepository.GetPipelineOutputByUrl(ctx, iacPipeline)
}

func (iu *iacPipelineUsecase) GetPipelineStatus(c context.Context, iacIacPipeline domain.IacPipeline) (string, error) {
	ctx, cancel := context.WithTimeout(c, iu.contextTimeout)
	defer cancel()
	return iu.iacPipelineRepository.GetPipelineStatus(ctx, iacIacPipeline)
}

func (iu *iacPipelineUsecase) GetPipelineLog(c context.Context, iacPipeline domain.IacPipeline) ([]byte, error) {
	ctx, cancel := context.WithTimeout(c, iu.contextTimeout)
	defer cancel()
	return iu.iacPipelineRepository.GetPipelineLog(ctx, iacPipeline)
}
