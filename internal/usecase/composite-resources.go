package usecase

import (
	"context"
	"time"

	"github.com/TranThang-2804/infrastructure-engine/internal/domain"
)

type compositeResourceUsecase struct {
	compositeResourceRepository domain.CompositeResourceRepository
	contextTimeout              time.Duration
}

func NewCompositeResourceUsecase(compositeResourceRepository domain.CompositeResourceRepository, timeout time.Duration) domain.CompositeResourceUsecase {
	return &compositeResourceUsecase{
		compositeResourceRepository: compositeResourceRepository,
		contextTimeout:              timeout,
	}
}

func (cu *compositeResourceUsecase) GetAll(c context.Context) ([]domain.CompositeResource, error) {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.compositeResourceRepository.GetAll(ctx)
}

func (cu *compositeResourceUsecase) Create(c context.Context, createCompositeResourceRequest domain.CreateCompositeResourceRequest) (domain.CompositeResource, error) {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()

  compositeResource := domain.CompositeResource{
  }
	return cu.compositeResourceRepository.Create(ctx, compositeResource)
}
