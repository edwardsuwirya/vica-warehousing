package repositories

import (
	"crypto/md5"
	"fmt"
	"warehousing/models"
)

type warehouseRepository struct {
	warehouseCollection []*models.Warehouse
	repo                IWarehouseFileRepo
}

func NewWarehouseRepository(repoInfra IWarehouseFileRepo) IWarehouseRepository {
	warehouseCollection := make([]*models.Warehouse, 0)
	return &warehouseRepository{warehouseCollection, repoInfra}
}

func (br *warehouseRepository) AddNewWarehouse(warehouse *models.Warehouse) *models.Warehouse {
	data := []byte(warehouse.Name)
	warehouse.Kode = fmt.Sprintf("%x", md5.Sum(data))
	br.repo.SaveToFile(warehouse)
	return warehouse
}

func (br *warehouseRepository) FindAllWarehouse() []*models.Warehouse {
	br.warehouseCollection = br.repo.ReadFile()
	return br.warehouseCollection
}
