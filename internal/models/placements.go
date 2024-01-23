package models

type Placement struct {
	StockId   StockId
	ProductId ProductId
	Count     ProductCount
	Reserved  ProductCount
}

type ReserveProducts = map[ProductId]ProductCount
type UnreserveProducts = map[ProductId]ProductCount
