package unreserve

import (
	"database/sql"

	"github.com/mrLandyrev/stock/internal/errors"
	"github.com/mrLandyrev/stock/internal/models"
)

type PlacementsRepository interface {
	BeginTx() (*sql.Tx, error)
	CommitTx(tx *sql.Tx) error
	RollbackTx(tx *sql.Tx) error
	GetByProductIds(productIds []models.ProductId, tx *sql.Tx, lockRows bool) ([]models.Placement, error)
	SetReservedByProductIdAndStockId(productId models.ProductId, stockId models.StockId, count models.ProductCount, tx *sql.Tx) error
}

type UnreserveUseCase struct {
	placementsRepository PlacementsRepository
}

func NewUnreserveUseCase(placementsRepository PlacementsRepository) *UnreserveUseCase {
	return &UnreserveUseCase{
		placementsRepository: placementsRepository,
	}
}

func (usecase *UnreserveUseCase) Execute(products models.UnreserveProducts) error {
	productIds := make([]models.ProductId, len(products))
	for productId := range products {
		productIds = append(productIds, productId)
	}

	tx, err := usecase.placementsRepository.BeginTx()
	if err != nil {
		return err
	}
	defer usecase.placementsRepository.RollbackTx(tx)

	placements, err := usecase.placementsRepository.GetByProductIds(productIds, tx, true)
	if err != nil {
		return err
	}

	placementsMap := usecase.transformPlacementsListToMapByProductId(placements)

	for productId, productCount := range products {
		placementsOfProduct := placementsMap[productId]

		if err := usecase.unreserveProduct(productId, productCount, placementsOfProduct, tx); err != nil {
			return err
		}
	}

	return usecase.placementsRepository.CommitTx(tx)
}

func (usecase *UnreserveUseCase) transformPlacementsListToMapByProductId(list []models.Placement) map[models.ProductId][]models.Placement {
	res := make(map[models.ProductId][]models.Placement)

	for _, placement := range list {
		bucket, ok := res[placement.ProductId]
		if !ok {
			bucket = make([]models.Placement, 0)
		}
		res[placement.ProductId] = append(bucket, placement)
	}

	return res
}

func (usecase *UnreserveUseCase) unreserveProduct(productId models.ProductId, count models.ProductCount, placements []models.Placement, tx *sql.Tx) error {
	for _, placement := range placements {
		if placement.Reserved == 0 {
			continue
		}
		if placement.Reserved >= count {
			return usecase.placementsRepository.SetReservedByProductIdAndStockId(productId, placement.StockId, placement.Reserved-count, tx)
		}
		if placement.Reserved < count {
			err := usecase.placementsRepository.SetReservedByProductIdAndStockId(productId, placement.StockId, 0, tx)
			if err != nil {
				return err
			}
			count -= placement.Reserved
		}
	}

	if count > 0 {
		return errors.ErrProductReserveNotFound
	}

	return nil
}
