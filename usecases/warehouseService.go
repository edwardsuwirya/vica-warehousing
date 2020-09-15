package usecases

import (
	"errors"
	"fmt"
	"warehousing/appConstant"
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
		err := errors.New(fmt.Sprintf(appConstant.MessageEmptyFormat, "Warehouse Name"))
		appLogging.Logger.LogError(appConstant.WarehouseService, appConstant.WarehouseRegistration, err.Error())
		return nil, err
	}
	newWarehouse := bs.r.AddNewWarehouse(b)
	appLogging.Logger.LogDebug(appConstant.WarehouseService, appConstant.WarehouseRegistration, newWarehouse)

	return newWarehouse, nil
}

func (bs *warehouseService) GetAllWarehouse() []*models.Warehouse {
	return bs.r.FindAllWarehouse()
}
