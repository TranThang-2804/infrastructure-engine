package domain

import (
	"context"

	"github.com/TranThang-2804/infrastructure-engine/internal/shared/constant"
)

type Resource struct {
	Name               string                  `json:"name" validate:"required"`
	Id                 string                  `json:"id" validate:"required"`
	Status             constant.ResourceStatus `json:"status" validate:"required"`
	Description        string                  `json:"description" validate:"required"`
	IacTemplateId      string                  `json:"iacTemplateId" validate:"required"`
	IacTemplateVersion string                  `json:"iacTemplateVersion" validate:"required"`
	ResourceValue      string                  `json:"spec" validate:"required"`
	RunId              []string                `json:"runId"`
}

type CompositeResource struct {
	Name             string                    `json:"name" validate:"required"`
	Id               string                    `json:"id" validate:"required"`
	Description      string                    `json:"description" validate:"required"`
	BluePrintId      string                    `json:"bluePrintId" validate:"required"`
	BluePrintVersion string                    `json:"bluePrintVersion" validate:"required"`
	CreatedAt        string                    `json:"createdAt" validate:"required"`
	CreatedBy        string                    `json:"createdBy" validate:"required"`
	LastModifiedAt   string                    `json:"lastModifiedAt" validate:"required"`
	LastModifiedBy   string                    `json:"lastModifiedBy" validate:"required"`
	Spec             map[string]interface{}    `json:"spec" validate:"required"`
	Status           constant.ResourceStatus   `json:"status" validate:"required"`
	Resources        []Resource                `json:"resources" validate:"required"`
	Metadata         CompositeResourceMetadata `json:"metadata,omitempty"`
}

type CompositeResourceMetadata struct {
	Group   string `json:"group"`
	Project string `json:"project"`
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
	Name             string                    `json:"name" validate:"required"`
	Description      string                    `json:"description" validate:"required"`
	Spec             map[string]interface{}    `json:"spec" validate:"required"`
	BluePrintId      string                    `json:"bluePrintId" validate:"required"`
	BluePrintVersion string                    `json:"bluePrintVersion" validate:"required"`
	MetaData         CompositeResourceMetadata `json:"metadata,omitempty"`
}

type CreateCompositeResourceResponse struct {
	CompositeResource CompositeResource `json:"compositeResource"`
	Status            string            `json:"status"`
}

type CompositeResourceRepository interface {
	Create(c context.Context, compositeResource CompositeResource) (CompositeResource, error)
	GetAll(c context.Context) ([]CompositeResource, error)
}

type CompositeResourceUsecase interface {
	GetAll(c context.Context) ([]CompositeResource, error)
	Create(c context.Context, CreateCompositeResourceRequest CreateCompositeResourceRequest) (CompositeResource, error)
}
