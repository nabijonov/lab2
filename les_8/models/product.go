package models

import "sync"

type Product struct {
	Id          int64
	Name        string
	Price       int64
	Description string
}

type ProductCache struct {
	MU *sync.RWMutex
	Products map[int64]*Product
}

type ProductClick struct{
	MU *sync.RWMutex
	ProductClick map[int64]int64
}