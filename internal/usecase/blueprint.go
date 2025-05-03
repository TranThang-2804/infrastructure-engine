package usecase

import (
	"context"
	"time"

	"github.com/TranThang-2804/infrastructure-engine/internal/domain"
)

type bluePrintUsecase struct {
	bluePrintRepository domain.BluePrintRepository
	contextTimeout      time.Duration
}

func NewBluePrintUsecase(bluePrintRepository domain.BluePrintRepository, timeout time.Duration) domain.BluePrintUsecase {
	return &bluePrintUsecase{
		bluePrintRepository: bluePrintRepository,
		contextTimeout:      timeout,
	}
}

func (bp *bluePrintUsecase) GetAll(c context.Context) ([]domain.BluePrint, error) {
	ctx, cancel := context.WithTimeout(c, bp.contextTimeout)
	defer cancel()
	return bp.bluePrintRepository.GetAll(ctx)
}

func (bp *bluePrintUsecase) GetById(c context.Context, id string) (domain.BluePrint, error) {
  ctx, cancel := context.WithTimeout(c, bp.contextTimeout)
  defer cancel()
  return bp.bluePrintRepository.GetById(ctx, id)
}

func (bp *bluePrintUsecase) GetByIdAndVersion(c context.Context, id string, version string) (domain.BluePrintVersion, error) {
  ctx, cancel := context.WithTimeout(c, bp.contextTimeout)
  defer cancel()
  return bp.bluePrintRepository.GetByIdAndVersion(ctx, id, version)
}
