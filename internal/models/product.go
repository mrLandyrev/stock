package models

type ProductId string
type ProductCount int32
type ProductSize struct {
	X int32
	Y int32
	Z int32
}

type Product struct {
	Id   ProductId
	Name string
	Size ProductSize
}
