package fakedb

import (
	"time"

	"github.com/nabijonov/lab2/les_8/models"
	"github.com/nabijonov/lab2/les_8/storage"
)

type storagen struct {
}

func NewFakeStorage()storage.IStorage{
	return &storagen{}
}

func (s *storagen) GetProductByID(id int64) (*models.Product, error) {
	
	time.Sleep(1*time.Second)
	return &models.Product{Id: id,Name: "Name_test",Price: 23,Description: "bla"}, nil
}