package usecase

import (
	"context"
	"time"

	"github.com/TranThang-2804/infrastructure-engine/internal/domain"
	"github.com/TranThang-2804/infrastructure-engine/internal/shared/constant"
	"github.com/TranThang-2804/infrastructure-engine/internal/shared/log"
	"github.com/TranThang-2804/infrastructure-engine/internal/utils"
)

type compositeResourceUsecase struct {
	compositeResourceRepository domain.CompositeResourceRepository
	bluePrintUsecase            domain.BluePrintUsecase
	contextTimeout              time.Duration
}

func NewCompositeResourceUsecase(compositeResourceRepository domain.CompositeResourceRepository, bluePrintUsecase domain.BluePrintUsecase, timeout time.Duration) domain.CompositeResourceUsecase {
	return &compositeResourceUsecase{
		compositeResourceRepository: compositeResourceRepository,
		bluePrintUsecase:            bluePrintUsecase,
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

	// Validate BluePrintId
	// bluePrint, err := cu.bluePrintUsecase.GetById(c, createCompositeResourceRequest.BluePrintId)
	//  if err != nil {
	//    log.Logger.Error("Error getting blueprint by id", "error", err.Error())
	//    return domain.CompositeResource{}, err
	//  }

	// Validate Spec With JsonSchema


	// Generate uuid
	log.Logger.Debug("Generating uuidv7")
	uuid, err := utils.GenerateUUIDv7()
	if err != nil {
		log.Logger.Error("Error getting all composite resources", "error", err.Error())
		return domain.CompositeResource{}, err
	}

	// Get current time
	currentTime := time.Now()
	currentDate := currentTime.Format("2006-01-02")

	// Getting user created
	createdBy := "anonymous"

	compositeResource := domain.CompositeResource{
		Name:           createCompositeResourceRequest.Name,
		Description:    createCompositeResourceRequest.Description,
		BluePrintId:    createCompositeResourceRequest.BluePrintId,
		Id:             uuid,
		CreatedAt:      currentDate,
		CreatedBy:      createdBy,
		LastModifiedAt: currentDate,
		LastModifiedBy: createdBy,
		Spec:           createCompositeResourceRequest.Spec,
		Status:         constant.Pending,
		Resources:      nil,
		Metadata:       createCompositeResourceRequest.MetaData,
	}

	log.Logger.Debug("CompositeResourceUsecase", "compositeResource", compositeResource)

	return cu.compositeResourceRepository.Create(ctx, compositeResource)
}
