package domain

import (
	"context"

	"github.com/TranThang-2804/infrastructure-engine/internal/shared/constant"
)

type resource struct {
	Name           string                  `json:"name"`
	Id             string                  `json:"id"`
	Status         constant.ResourceStatus `json:"status"`
	Description    string                  `json:"description"`
	CreatedAt      string                  `json:"createdAt"`
	CreatedBy      string                  `json:"createdBy"`
	LastModifiedAt string                  `json:"lastModifiedAt"`
	LastModifiedBy string                  `json:"lastModifiedBy"`
	Spec           string                  `json:"spec"`
	BluePrintName  string                  `json:"bluePrintName"`
}

type CompositeResource struct {
	Name           string                  `json:"name"`
	Id             string                  `json:"id"`
	Description    string                  `json:"description"`
	CreatedAt      string                  `json:"createdAt"`
	CreatedBy      string                  `json:"createdBy"`
	LastModifiedAt string                  `json:"lastModifiedAt"`
	LastModifiedBy string                  `json:"lastModifiedBy"`
	Spec           string                  `json:"spec"`
	Status         constant.ResourceStatus `json:"status"`
	Resources      []resource              `json:"resources"`
}

type CompositeResourceRepository interface {
	Create(c context.Context, compositeResource CompositeResource) (CompositeResource, error)
	GetAll(c context.Context) ([]CompositeResource, error)
}

type CompositeResourceUsecase interface {
	GetAll(c context.Context) ([]CompositeResource, error)
}
