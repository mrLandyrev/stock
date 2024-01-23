package models

type Balance map[ProductId]ProductCount
type StockId string

type Stock struct {
	Id          StockId
	Name        string
	IsAvailable bool
}
