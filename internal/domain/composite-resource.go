package domain

import (
	"context"

	"github.com/TranThang-2804/infrastructure-engine/internal/shared/constant"
)

type Resource struct {
	Name               string                  `json:"name" yaml:"name" validate:"required"`
	Id                 string                  `json:"id" yaml:"id" validate:"required"`
	Status             constant.ResourceStatus `json:"status" yaml:"status" validate:"required"`
	Description        string                  `json:"description" yaml:"description" validate:"required"`
	IacTemplateId      string                  `json:"iacTemplateId" yaml:"iacTemplateId" validate:"required"`
	IacTemplateVersion string                  `json:"iacTemplateVersion" yaml:"iacTemplateVersion" validate:"required"`
	ResourceValue      string                  `json:"spec" yaml:"spec" validate:"required"`
	RunId              []string                `json:"runId" yaml:"runId"`
}

type CompositeResource struct {
	Name             string                    `json:"name" yaml:"name" validate:"required"`
	Id               string                    `json:"id" yaml:"id" validate:"required"`
	Description      string                    `json:"description" yaml:"description" validate:"required"`
	BluePrintId      string                    `json:"bluePrintId" yaml:"bluePrintId" validate:"required"`
	BluePrintVersion string                    `json:"bluePrintVersion" yaml:"bluePrintVersion" validate:"required"`
	CreatedAt        string                    `json:"createdAt" yaml:"createdAt" validate:"required"`
	CreatedBy        string                    `json:"createdBy" yaml:"createdBy" validate:"required"`
	LastModifiedAt   string                    `json:"lastModifiedAt" yaml:"lastModifiedAt" validate:"required"`
	LastModifiedBy   string                    `json:"lastModifiedBy" yaml:"lastModifiedBy" validate:"required"`
	Spec             map[string]interface{}    `json:"spec" yaml:"spec" validate:"required"`
	Status           constant.ResourceStatus   `json:"status" yaml:"status" validate:"required"`
	Resources        []Resource                `json:"resources" yaml:"resources" validate:"required"`
	Metadata         CompositeResourceMetadata `json:"metadata,omitempty" yaml:"metadata,omitempty"`
}

type CompositeResourceMetadata struct {
	Group   string `json:"group" yaml:"group"`
	Project string `json:"project" yaml:"project"`
}

type GetCompositeResourceRequest struct {
	Name          string `json:"name,omitempty" yaml:"name,omitempty" validate:"required"`
	BluePrintType string `json:"bluePrintType,omitempty" yaml:"bluePrintType,omitempty" validate:"required"`
	Id            string `json:"id,omitempty" yaml:"bluePrintType,omitempty" validate:"required"`
}

type GetCompositeResourceResponse struct {
	CompositeResource []CompositeResource `json:"compositeResources" yaml:"compositeResources"`
}

type CreateCompositeResourceRequest struct {
	Name             string                    `json:"name" yaml:"name" validate:"required"`
	Description      string                    `json:"description" yaml:"description" validate:"required"`
	Spec             map[string]interface{}    `json:"spec" yaml:"spec" validate:"required"`
	BluePrintId      string                    `json:"bluePrintId" yaml:"bluePrintId" validate:"required"`
	BluePrintVersion string                    `json:"bluePrintVersion" yaml:"bluePrintVersion" validate:"required"`
	MetaData         CompositeResourceMetadata `json:"metadata,omitempty" yaml:"metadata,omitempty"`
}

type CreateCompositeResourceResponse struct {
	CompositeResource CompositeResource `json:"compositeResource" yaml:"compositeResource"`
	Status            string            `json:"status" yaml:"status"`
}

type UpdateCompositeResourceRequest struct {
	Name             string                    `json:"name" yaml:"name" validate:"required"`
	Description      string                    `json:"description" yaml:"description" validate:"required"`
	Spec             map[string]interface{}    `json:"spec" yaml:"spec" validate:"required"`
	BluePrintId      string                    `json:"bluePrintId" yaml:"bluePrintId" validate:"required"`
	BluePrintVersion string                    `json:"bluePrintVersion" yaml:"bluePrintVersion" validate:"required"`
	MetaData         CompositeResourceMetadata `json:"metadata,omitempty" yaml:"metadata,omitempty"`
}

type UpdateCompositeResourceResponse struct {
	CompositeResource CompositeResource `json:"compositeResource" yaml:"compositeResource"`
	Status            string            `json:"status" yaml:"status"`
}

type DeleteCompositeResourceRequest struct {
	Name             string                    `json:"name" yaml:"name" validate:"required"`
	Description      string                    `json:"description" yaml:"description" validate:"required"`
	Spec             map[string]interface{}    `json:"spec" yaml:"spec" validate:"required"`
	BluePrintId      string                    `json:"bluePrintId" yaml:"bluePrintId" validate:"required"`
	BluePrintVersion string                    `json:"bluePrintVersion" yaml:"bluePrintVersion" validate:"required"`
	MetaData         CompositeResourceMetadata `json:"metadata,omitempty" yaml:"metadata,omitempty"`
}

type DeleteCompositeResourceResponse struct {
	CompositeResource CompositeResource `json:"compositeResource" yaml:"compositeResource"`
	Status            string            `json:"status" yaml:"status"`
}

type CompositeResourceRepository interface {
	GetAll(c context.Context) ([]CompositeResource, error)
	Create(c context.Context, compositeResource CompositeResource) (CompositeResource, error)
	Update(c context.Context, compositeResource CompositeResource) (CompositeResource, error)
	Delete(c context.Context, compositeResource CompositeResource) (CompositeResource, error)
}

type CompositeResourceEventPublisher interface {
  PublishToPendingSubject(c context.Context, compositeResource CompositeResource) (error)
  PublishToProvisioningSubject(c context.Context, compositeResource CompositeResource) (error)
  PublishToDeletingSubject(c context.Context, compositeResource CompositeResource) (error)
}

type CompositeResourceEventConsumer interface {
  StartConsumer() (error)
}

type CompositeResourceUsecase interface {
	GetAll(c context.Context) ([]CompositeResource, error)
	Create(c context.Context, CreateCompositeResourceRequest CreateCompositeResourceRequest) (CompositeResource, error)
	Update(c context.Context, CreateCompositeResourceRequest UpdateCompositeResourceRequest) (CompositeResource, error)
	Delete(c context.Context, CreateCompositeResourceRequest DeleteCompositeResourceRequest) (CompositeResource, error)
	HandlePending(message string) (error)
	HandleProvisioning(message string) (error)
	HandleDeleting(message string) (error)
}
