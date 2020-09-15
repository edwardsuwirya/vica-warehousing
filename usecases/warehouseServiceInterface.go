package usecases

import "warehousing/models"

type IWarehouseService interface {
	RegisterNewWarehouse(b *models.Warehouse) (*models.Warehouse, error)
	GetAllWarehouse() []*models.Warehouse
}
