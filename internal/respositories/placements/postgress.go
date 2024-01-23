package placements

import (
	"context"
	"database/sql"
	"strings"

	"github.com/lib/pq"
	"github.com/mrLandyrev/stock/internal/models"
)

type PlacementsRepository struct {
	db *sql.DB
}

func NewPlacementsRepository(db *sql.DB) *PlacementsRepository {
	return &PlacementsRepository{
		db: db,
	}
}

func (super *PlacementsRepository) BeginTx() (*sql.Tx, error) {
	return super.db.BeginTx(context.TODO(), nil)
}

func (super *PlacementsRepository) CommitTx(tx *sql.Tx) error {
	return tx.Commit()
}

func (super *PlacementsRepository) RollbackTx(tx *sql.Tx) error {
	return tx.Rollback()
}

func (super *PlacementsRepository) GetByProductIds(productIds []models.ProductId, tx *sql.Tx, lockRows bool) (result []models.Placement, err error) {
	var builder strings.Builder
	builder.WriteString("SELECT stock_id, product_id, count, reserved FROM placements WHERE product_id::text = ANY($1)")
	if tx != nil && lockRows {
		builder.WriteString(" FOR UPDATE")
	}

	query := builder.String()
	params := pq.Array(productIds)

	var rows *sql.Rows

	if tx != nil {
		rows, err = tx.QueryContext(context.TODO(), query, params)
	} else {
		rows, err = super.db.QueryContext(context.TODO(), query, params)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var record models.Placement
		err = rows.Scan(&record.StockId, &record.ProductId, &record.Count, &record.Reserved)
		if err != nil {
			return nil, err
		}
		result = append(result, record)
	}

	return result, nil
}

func (super *PlacementsRepository) GetByStockIds(stockIds []models.StockId) (result []models.Placement, err error) {
	rows, err := super.db.QueryContext(context.TODO(), "SELECT stock_id, product_id, count, reserved FROM placements WHERE stock_id::text = ANY($1)", pq.Array(stockIds))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var record models.Placement
		err = rows.Scan(&record.StockId, &record.ProductId, &record.Count, &record.Reserved)
		if err != nil {
			return nil, err
		}
		result = append(result, record)
	}

	return result, nil
}

func (super *PlacementsRepository) SetReservedByProductIdAndStockId(productId models.ProductId, stockId models.StockId, count models.ProductCount, tx *sql.Tx) (err error) {
	query := "UPDATE placements SET reserved = $1 WHERE stock_id = $2 AND product_id::text = $3"
	params := []any{count, stockId, productId}

	if tx != nil {
		_, err = tx.ExecContext(context.TODO(), query, params...)
	} else {
		_, err = super.db.ExecContext(context.TODO(), query, params...)
	}

	return
}
