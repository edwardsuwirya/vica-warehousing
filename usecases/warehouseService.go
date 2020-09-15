package usecases

import (
	"errors"
	"warehousing/appLogging"
	"warehousing/models"
	"warehousing/repositories"
)

type warehouseService struct {
	r repositories.IWarehouseRepository
}

func NewWarehouseService(repo repositories.IWarehouseRepository) IWarehouseService {
	return &warehouseService{r: repo}
}

func (bs *warehouseService) RegisterNewWarehouse(b *models.Warehouse) (*models.Warehouse, error) {
	if b.Name == "" {
		err := errors.New("Warehouse name can not be empty")
		appLogging.Logger.LogError("WarehouseDelivery", "Registration New Warehouse", err.Error())
		return nil, err
	}
	newWarehouse := bs.r.AddNewWarehouse(b)
	appLogging.Logger.LogDebug("WarehouseDelivery", "Registration New Warehouse", newWarehouse)

	return newWarehouse, nil
}

func (bs *warehouseService) GetAllWarehouse() []*models.Warehouse {
	return bs.r.FindAllWarehouse()
}
