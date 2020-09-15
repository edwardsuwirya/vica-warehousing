package repositories

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"warehousing/models"
)

type WarehouseRepositoryInfrastructure struct {
	dataPath string
}

func NewWarehouseRepoInfra(dataPath string) *WarehouseRepositoryInfrastructure {
	_, err := os.Stat(dataPath)
	var file *os.File
	defer file.Close()
	if err != nil {
		if os.IsNotExist(err) {
			file, err = os.Create(dataPath)
			if err != nil {
				panic(err)
			}
		}
	}
	return &WarehouseRepositoryInfrastructure{dataPath}
}

func (bri *WarehouseRepositoryInfrastructure) saveToFile(warehouseCollection *models.Warehouse) {
	//file, _ := json.MarshalIndent(warehouseCollection, "", " ")
	//_ = ioutil.WriteFile(bri.dataPath, file, 0644)
	file, err := os.OpenFile(bri.dataPath, os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	if _, err := file.WriteString(warehouseCollection.ToString()); err != nil {
		panic(err)
	}
}

func (bri *WarehouseRepositoryInfrastructure) readFile() []*models.Warehouse {
	//file, _ := ioutil.ReadFile(bri.dataPath)
	//_ = json.Unmarshal(file, warehouseCollection)
	file, err := os.OpenFile(bri.dataPath, os.O_RDONLY, 0644)
	result := make([]*models.Warehouse, 0)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		dt := scanner.Text()
		dts := strings.Split(dt, ",")
		kode := dts[0]
		name := dts[1]
		address := dts[2]
		large, _ := strconv.ParseFloat(dts[3], 64)
		information := dts[4]
		price, _ := strconv.ParseFloat(dts[5], 64)

		result = append(result, &models.Warehouse{kode, name, address, large, information, price})
	}
	return result
}
