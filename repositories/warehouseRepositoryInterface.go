package repositories

import "warehousing/models"

type IWarehouseRepository interface {
	AddNewWarehouse(warehouse *models.Warehouse) *models.Warehouse
	FindAllWarehouse() []*models.Warehouse
}
