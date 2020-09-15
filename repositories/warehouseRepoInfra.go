package repositories

import (
	"encoding/json"
	"io/ioutil"
	"warehousing/models"
)

type WarehouseRepositoryInfrastructure struct {
	dataPath string
}

func NewWarehouseRepoInfra(dataPath string) *WarehouseRepositoryInfrastructure {
	return &WarehouseRepositoryInfrastructure{dataPath}
}

func (bri *WarehouseRepositoryInfrastructure) saveToFile(warehouseCollection *[]*models.Warehouse) {
	file, _ := json.MarshalIndent(warehouseCollection, "", " ")
	_ = ioutil.WriteFile(bri.dataPath, file, 0644)
}

func (bri *WarehouseRepositoryInfrastructure) readFile(warehouseCollection *[]*models.Warehouse) {
	file, _ := ioutil.ReadFile(bri.dataPath)
	_ = json.Unmarshal(file, warehouseCollection)
}
