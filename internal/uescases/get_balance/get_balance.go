package get_balance

import "github.com/mrLandyrev/stock/internal/models"

type PlacementsRepository interface {
	GetByStockIds(stockId []models.StockId) (result []models.Placement, err error)
}

type GetBalanceUseCase struct {
	placementsRepository PlacementsRepository
}

func NewGetBalanceUseCase(placementsRepository PlacementsRepository) *GetBalanceUseCase {
	return &GetBalanceUseCase{
		placementsRepository: placementsRepository,
	}
}

func (usecase *GetBalanceUseCase) Execute(stockId models.StockId) (models.Balance, error) {
	placements, err := usecase.placementsRepository.GetByStockIds([]models.StockId{stockId})
	if err != nil {
		return nil, err
	}

	result := make(models.Balance, len(placements))
	for _, placement := range placements {
		result[placement.ProductId] = placement.Count - placement.Reserved
	}

	return result, nil
}
