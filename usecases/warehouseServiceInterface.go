package usecases

import "warehousing/models"

type IWarehouseService interface {
	RegisterNewWarehouse(b *models.Warehouse)
	GetAllWarehouse() []*models.Warehouse
}
