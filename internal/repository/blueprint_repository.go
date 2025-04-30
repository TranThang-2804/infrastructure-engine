package repository

import (
	"context"

	"github.com/TranThang-2804/infrastructure-engine/internal/adapter/git"
	"github.com/TranThang-2804/infrastructure-engine/internal/domain"
)

type bluePrintRepository struct {
	gitStore   git.GitStore
}

func NewBluePrintRepository(gitStore git.GitStore) domain.BluePrintRepository {
	return &bluePrintRepository{
    gitStore: gitStore,
  }
}

func (br *bluePrintRepository) GetAll(c context.Context) ([]domain.BluePrint, error) {
	var bluePrint []domain.BluePrint = []domain.BluePrint{}

	var err error

	return bluePrint, err
}
