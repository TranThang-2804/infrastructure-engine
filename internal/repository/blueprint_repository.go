package repository

import (
	"context"

	"github.com/TranThang-2804/infrastructure-engine/internal/domain"
)

type bluePrintRepository struct {
	// database   mongo.Database
}

func NewBluePrintRepository() domain.BluePrintRepository {
	return &bluePrintRepository{}
}

func (br *bluePrintRepository) GetAll(c context.Context) ([]domain.BluePrint, error) {
	var bluePrint []domain.BluePrint

	var err error

	return bluePrint, err
}
