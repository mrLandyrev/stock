package respositories

import (
	"context"
	"database/sql"
	"strings"

	"github.com/lib/pq"
)

type Placement struct {
	StockId   string
	ProductId string
	Count     int64
	Reseved   int64
}

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

func (super *PlacementsRepository) GetByProductIds(productIds []string, tx *sql.Tx, lockRows bool) (result []Placement, err error) {
	var builder strings.Builder
	builder.WriteString("SELECT stock_id, product_id, count, reserved FROM placements WHERE product_id = ANY($1)")
	if lockRows {
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
		var record Placement
		err = rows.Scan(&record.StockId, &record.ProductId, &record.Count, &record.Reseved)
		if err != nil {
			return nil, err
		}
		result = append(result, record)
	}

	return result, nil
}

func (super *PlacementsRepository) ReserveByProductIdAndStockId(productId string, stockId string, tx *sql.Tx) (err error) {
	query := "UPDATE placements SET count = count + 1 WHERE stock_id = $1, product_id = $2"
	params := []any{productId, stockId}

	if tx != nil {
		_, err = tx.ExecContext(context.TODO(), query, params...)
	} else {
		_, err = super.db.ExecContext(context.TODO(), query, params...)
	}

	return
}
