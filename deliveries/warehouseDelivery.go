package deliveries

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"warehousing/models"
	"warehousing/repositories"
	"warehousing/usecases"
	"warehousing/utils"
)

type WarehouseDelivery struct {
	warehouseService usecases.IWarehouseService
}

func NewWarehouseDelivery() *WarehouseDelivery {
	repo := repositories.NewWarehouseRepository()
	warehouseService := usecases.NewWarehouseService(repo)
	return &WarehouseDelivery{warehouseService}
}

func (bd *WarehouseDelivery) RegistrationWarehouseForm(backToMainMenu callbackFn) {
	utils.ConsoleClear()
	var name string
	var address string
	var large float64
	var information string
	var price float64
	var confirmation string
	fmt.Println()
	fmt.Printf("%s\n", "Warehouse Registration Form")
	fmt.Printf("%s\n", strings.Repeat("-", 30))
	scanner := bufio.NewReader(os.Stdin)
	fmt.Print("Name : ")
	sName, _ := scanner.ReadString('\n')
	name = strings.TrimSpace(sName)
	fmt.Print("Address : ")
	sAddress, _ := scanner.ReadString('\n')
	address = strings.TrimSpace(sAddress)
	fmt.Print("Large (/m2): ")
	sLarge, _ := scanner.ReadString('\n')
	large, _ = strconv.ParseFloat(strings.TrimSpace(sLarge), 64)
	fmt.Print("Information : ")
	sInformation, _ := scanner.ReadString('\n')
	information = strings.TrimSpace(sInformation)
	fmt.Print("Price (Rp.): ")
	sPrice, _ := scanner.ReadString('\n')
	price, _ = strconv.ParseFloat(strings.TrimSpace(sPrice), 64)

	fmt.Println("Save to collection? :Y/N")
	fmt.Scanln(&confirmation)

	utils.ConsoleClear()
	if confirmation == "Y" {
		newWarehouse := models.NewWarehouse(name, address, large, information, price)
		wh, err := bd.warehouseService.RegisterNewWarehouse(&newWarehouse)
		if err != nil {
			fmt.Println("Error Notification", err)
		} else {
			fmt.Printf("Success Notification %s is successfully registered\n", wh.Name)
		}
	}
	backToMainMenu()
}

func (bd *WarehouseDelivery) ListWarehouseForm(backToMainMenu callbackFn) {
	utils.ConsoleClear()
	warehouses := bd.warehouseService.GetAllWarehouse()
	for _, w := range warehouses {
		fmt.Println(*w)
	}
	backToMainMenu()
}
