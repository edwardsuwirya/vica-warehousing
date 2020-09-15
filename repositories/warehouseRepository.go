package repositories

import (
	"crypto/md5"
	"fmt"
	"warehousing/models"
)

type warehouseRepository struct {
	warehouseCollection []*models.Warehouse
}

func NewWarehouseRepository() IWarehouseRepository {
	warehouseCollection := make([]*models.Warehouse, 0)
	return &warehouseRepository{warehouseCollection}
}

func (br *warehouseRepository) AddNewWarehouse(warehouse *models.Warehouse) *models.Warehouse {
	data := []byte(warehouse.Name)
	warehouse.Kode = fmt.Sprintf("%x", md5.Sum(data))
	br.warehouseCollection = append(br.warehouseCollection, warehouse)
	return warehouse
}

func (br *warehouseRepository) FindAllWarehouse() []*models.Warehouse {
	return br.warehouseCollection
}
