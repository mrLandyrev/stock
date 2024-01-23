package reserve

import (
	"database/sql"
	"fmt"

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

type ReserveUseCase struct {
	placementsRepository PlacementsRepository
}

func NewReserveUseCase(placementsRepository PlacementsRepository) *ReserveUseCase {
	return &ReserveUseCase{
		placementsRepository: placementsRepository,
	}
}

func (usecase *ReserveUseCase) Execute(products models.ReserveProducts) error {
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
		fmt.Println(err)
		return err
	}

	fmt.Println(placements)

	placementsMap := usecase.transformPlacementsListToMapByProductId(placements)

	for productId, productCount := range products {
		placementsOfProduct := placementsMap[productId]

		if err := usecase.reserveProduct(productId, productCount, placementsOfProduct, tx); err != nil {
			return err
		}
	}

	return usecase.placementsRepository.CommitTx(tx)
}

func (super *ReserveUseCase) transformPlacementsListToMapByProductId(list []models.Placement) map[models.ProductId][]models.Placement {
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

func (super *ReserveUseCase) reserveProduct(productId models.ProductId, count models.ProductCount, placements []models.Placement, tx *sql.Tx) error {
	for _, placement := range placements {
		freeSlots := placement.Count - placement.Reserved
		if freeSlots <= 0 {
			continue
		}
		if freeSlots >= count {
			return super.placementsRepository.SetReservedByProductIdAndStockId(productId, placement.StockId, placement.Reserved+count, tx)
		}
		if freeSlots < count {
			err := super.placementsRepository.SetReservedByProductIdAndStockId(productId, placement.StockId, placement.Reserved+count-freeSlots, tx)
			if err != nil {
				return err
			}
			count -= freeSlots
		}
	}

	if count > 0 {
		return errors.ErrProductOutOfStock
	}

	return nil
}
