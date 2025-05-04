package usecase

import (
	"context"
	"time"

	"github.com/TranThang-2804/infrastructure-engine/internal/domain"
	"github.com/TranThang-2804/infrastructure-engine/internal/utils"
)

type iacTemplateUsecase struct {
	iacTemplateRepository domain.IacTemplateRepository
	contextTimeout        time.Duration
}

func NewIacTemplateUsecase(iacTemplateRepository domain.IacTemplateRepository) domain.IacTemplateUsecase {
	return &iacTemplateUsecase{
		iacTemplateRepository: iacTemplateRepository,
		contextTimeout:        utils.GetContextTimeout(),
	}
}

func (iu *iacTemplateUsecase) GetAll(c context.Context) ([]domain.IacTemplate, error) {
	ctx, cancel := context.WithTimeout(c, iu.contextTimeout)
	defer cancel()
	return iu.iacTemplateRepository.GetAll(ctx)
}
