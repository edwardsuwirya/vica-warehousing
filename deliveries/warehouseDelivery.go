package deliveries

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"warehousing/appConstant"
	"warehousing/config"
	"warehousing/models"
	"warehousing/repositories"
	"warehousing/usecases"
	"warehousing/utils"
)

type WarehouseDelivery struct {
	warehouseService usecases.IWarehouseService
}

func NewWarehouseDelivery(c *config.AppConfig) *WarehouseDelivery {
	//repoInfra := repositories.NewWarehouseRepoInfra(c.DataPath)
	repoInfra := repositories.NewWarehouseRepoAwsInfra(c.BucketName, c.DataPath, c.Aws)
	repo := repositories.NewWarehouseRepository(repoInfra)
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
	fmt.Printf("%s\n", appConstant.WarehouseRegistrationFormLabel)
	fmt.Printf("%s\n", strings.Repeat("-", 30))
	scanner := bufio.NewReader(os.Stdin)
	fmt.Print(appConstant.NameLabel, ":")
	sName, _ := scanner.ReadString('\n')
	name = strings.TrimSpace(sName)
	fmt.Print(appConstant.AddressLabel, ":")
	sAddress, _ := scanner.ReadString('\n')
	address = strings.TrimSpace(sAddress)
	fmt.Print(appConstant.LargeLabel, ":")
	sLarge, _ := scanner.ReadString('\n')
	large, _ = strconv.ParseFloat(strings.TrimSpace(sLarge), 64)
	fmt.Print(appConstant.InformationLabel, ":")
	sInformation, _ := scanner.ReadString('\n')
	information = strings.TrimSpace(sInformation)
	fmt.Print(appConstant.PriceLabel, ":")
	sPrice, _ := scanner.ReadString('\n')
	price, _ = strconv.ParseFloat(strings.TrimSpace(sPrice), 64)

	fmt.Printf(appConstant.SaveConfirmationFormat, "warehouse")
	fmt.Scanln(&confirmation)

	utils.ConsoleClear()
	if confirmation == "Y" {
		newWarehouse := models.NewWarehouse(name, address, large, information, price)
		wh, err := bd.warehouseService.RegisterNewWarehouse(&newWarehouse)
		if err != nil {
			fmt.Printf(appConstant.ErrorNotificationFormat, err.Error())
		} else {
			fmt.Printf(appConstant.SuccessRegistrationNotificationFormat, wh.Name)
		}
	}
	backToMainMenu()
}

func (bd *WarehouseDelivery) ListWarehouseForm(backToMainMenu callbackFn) {
	utils.ConsoleClear()
	warehouses := bd.warehouseService.GetAllWarehouse()
	fmt.Println("Warehouses")
	fmt.Printf("%s\n", strings.Repeat("=", 140))
	fmt.Printf("%-40s%-25s%-25s%10s %-10s%20s\n", appConstant.WarehouseIDLabel, appConstant.NameLabel, appConstant.AddressLabel, appConstant.LargeLabel, appConstant.InformationLabel, appConstant.PriceLabel)
	fmt.Printf("%s\n", strings.Repeat("-", 140))
	for _, w := range warehouses {
		fmt.Printf("%-40s%-25s%-25s%10.2f %-10s%15.2f\n", w.Kode, w.Name, w.Address, w.Large, w.Information, w.Price)
	}
	backToMainMenu()
}
