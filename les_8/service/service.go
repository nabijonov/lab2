package service

import (
	"sync"

	"github.com/nabijonov/lab2/les_8/models"
	"github.com/nabijonov/lab2/les_8/storage"
)

type IService interface {
	GetProductByID(id int64) (*models.Product, error)
}

type service struct{
	ProductCache *models.ProductCache
	clicks *models.ProductClick
	storage storage.IStorage
}

func NewService(strorage storage.IStorage) IService{
	productCache := &models.ProductCache{
		MU: &sync.RWMutex{},
		Products: make(map[int64]*models.Product),
	}
	clicks := &models.ProductClick{
		MU: &sync.RWMutex{},
		ProductClick: make(map[int64]int64),
	}
	return &service{ProductCache: productCache,clicks: clicks,storage: strorage}
}

func(s *service)GetProductByID(id int64) (*models.Product, error){
	return s.storage.GetProductByID(id)
}
