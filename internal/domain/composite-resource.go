package domain

import (
	"context"
	"encoding/json"

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
	Spec           json.RawMessage         `json:"spec"`
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
	Spec           json.RawMessage         `json:"spec"`
	Status         constant.ResourceStatus `json:"status"`
	Resources      []resource              `json:"resources"`
}

type CompositeResourceRepository interface {
	Create(c context.Context, compositeResource CompositeResource) (CompositeResource, error)
	GetAll(c context.Context) ([]CompositeResource, error)
}

type CompositeResourceUsecase interface {
	GetAll(c context.Context) ([]CompositeResource, error)
	Create(c context.Context, compositeResource CompositeResource) (CompositeResource, error)
}

type GetCompositeResourceRequest struct {
	Name          string `json:"name,omitempty"`
	BluePrintType string `json:"bluePrintType,omitempty"`
	Id            string `json:"id,omitempty"`
}

type GetCompositeResourceResponse struct {
	CompositeResource []CompositeResource `json:"compositeResources"`
}

type CreateCompositeResourceRequest struct {
	Name          string `json:"name"`
	Description   string
	Spec          json.RawMessage `json:"spec"`
	BluePrintType string          `json:"bluePrintType,omitempty"`
}

type CreateCompositeResourceResponse struct {
	CompositeResource CompositeResource `json:"compositeResource"`
	Status            string            `json:"status"`
}
