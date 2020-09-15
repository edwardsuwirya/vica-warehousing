package usecases

import (
	"warehousing/models"
	"warehousing/repositories"
)

type warehouseService struct {
	r repositories.IWarehouseRepository
}

func NewWarehouseService(repo repositories.IWarehouseRepository) IWarehouseService {
	return &warehouseService{r: repo}
}

func (bs *warehouseService) RegisterNewWarehouse(b *models.Warehouse) {
	bs.r.AddNewWarehouse(b)
}

func (bs *warehouseService) GetAllWarehouse() []*models.Warehouse {
	return bs.r.FindAllWarehouse()
}
