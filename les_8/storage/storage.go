package storage

import "github.com/nabijonov/lab2/les_8/models"

type IStorage interface {
	//
	GetProductByID(id int64) (*models.Product, error)
}