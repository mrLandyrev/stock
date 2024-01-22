package services

import (
	"database/sql"
	"errors"
)

var (
	ErrProductOutOfStock = errors.New("product out of stock")
)

type Placement struct {
	StockId   string
	ProductId string
	Count     int64
	Reseved   int64
}

type PlacementsRepository interface {
	BeginTx() (*sql.Tx, error)
	CommitTx(tx *sql.Tx) error
	RollbackTx(tx *sql.Tx) error
	GetByProductIds(productIds []string, tx *sql.Tx, lockRows bool) ([]Placement, error)
	ReserveByProductIdAndStockId(productId string, stockId string, tx *sql.Tx) error
}

type PlacementsService struct {
	placementsRepository PlacementsRepository
}

func NewPlacementsService(placementsRepository PlacementsRepository) *PlacementsService {
	return &PlacementsService{
		placementsRepository: placementsRepository,
	}
}

func (super *PlacementsService) ReserveProducts(productIds []string) error {
	tx, err := super.placementsRepository.BeginTx()

	if err != nil {
		return err
	}

	defer super.placementsRepository.RollbackTx(tx)

	placements, err := super.placementsRepository.GetByProductIds(productIds, tx, true)
	if err != nil {
		return err
	}

	placementsMap := super.transformPlacementsListToMapByProductId(placements)

	for _, productId := range productIds {
		placementsForProduct := placementsMap[productId]

		if err := super.reserveProduct(productId, placementsForProduct, tx); err != nil {
			return err
		}
	}

	return super.placementsRepository.CommitTx(tx)
}

func (super *PlacementsService) transformPlacementsListToMapByProductId(list []Placement) map[string][]Placement {
	res := make(map[string][]Placement)

	for _, placement := range list {
		bucket, ok := res[placement.ProductId]
		if !ok {
			bucket = make([]Placement, 0)
		}
		res[placement.ProductId] = append(bucket, placement)
	}

	return res
}

func (super *PlacementsService) reserveProduct(productId string, placements []Placement, tx *sql.Tx) error {
	for _, placement := range placements {
		if placement.Reseved < placement.Count {
			if err := super.placementsRepository.ReserveByProductIdAndStockId(productId, placement.StockId, tx); err != nil {
				return err
			}
			placement.Reseved++

			return nil
		}
	}

	return ErrProductOutOfStock
}
