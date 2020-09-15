package repositories

import (
	"crypto/md5"
	"fmt"
	"warehousing/models"
)

type warehouseRepository struct {
	warehouseCollection []*models.Warehouse
	repo                *WarehouseRepositoryInfrastructure
}

func NewWarehouseRepository(dataPath string) IWarehouseRepository {
	warehouseCollection := make([]*models.Warehouse, 0)
	repo := NewWarehouseRepoInfra(dataPath)
	return &warehouseRepository{warehouseCollection, repo}
}

func (br *warehouseRepository) AddNewWarehouse(warehouse *models.Warehouse) *models.Warehouse {
	data := []byte(warehouse.Name)
	warehouse.Kode = fmt.Sprintf("%x", md5.Sum(data))
	br.repo.saveToFile(warehouse)
	return warehouse
}

func (br *warehouseRepository) FindAllWarehouse() []*models.Warehouse {
	br.warehouseCollection = br.repo.readFile()
	return br.warehouseCollection
}
