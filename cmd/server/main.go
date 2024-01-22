package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/lib/pq"
	r "github.com/mrLandyrev/stock/internal/repositories"
)

func main() {
	main1()
}

type TestRecord struct {
	Name  string `db:"name"`
	Count int64  `db:"count"`
}

func getRows(tx *sql.Tx, names []string) (result []TestRecord, err error) {
	fmt.Println("get records")
	rows, err := tx.Query("SELECT * FROM test WHERE name = ANY($1) FOR UPDATE", pq.Array(names))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var record TestRecord
		err = rows.Scan(&record.Name, &record.Count)
		if err != nil {
			return nil, err
		}
		result = append(result, record)
	}

	return result, nil
}

func updateRecord(tx *sql.Tx, record TestRecord) error {
	_, err := tx.Exec("UPDATE test SET count = $1 WHERE name = $2", record.Count+1, record.Name)
	return err
}

func main1() {
	psqlConnectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "db", 5432, "postgres", "local", "test")
	fmt.Println(psqlConnectionString)

	db, err := sql.Open("postgres", psqlConnectionString)
	if err != nil {
		fmt.Println(err)
		return
	}
	r.NewPlacementsRepository(db)

	tx, err := db.BeginTx(context.TODO(), nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer tx.Rollback()

	records, err := getRows(tx, []string{"test", "test1"})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(records)

	for _, record := range records {
		time.Sleep(time.Second * 1)
		fmt.Println("update record")
		err := updateRecord(tx, record)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Println("ok")

	tx.Commit()
}
