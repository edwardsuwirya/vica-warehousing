package repositories

import "warehousing/models"

type IWarehouseFileRepo interface {
	SaveToFile(warehouseCollection *models.Warehouse)
	ReadFile() []*models.Warehouse
}
