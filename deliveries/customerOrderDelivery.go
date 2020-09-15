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

type CustomerOrderDelivery struct {
	customerOrderService usecases.ICustomerOrderService
}

func NewCustomerOrderDelivery() *CustomerOrderDelivery {
	repo := repositories.NewCustomerOrderRepository()
	customerOrderService := usecases.NewCustomerOrderService(repo)
	return &CustomerOrderDelivery{customerOrderService}
}

func (bd *CustomerOrderDelivery) CustomerOrderForm(backToMainMenu callbackFn) {
	utils.ConsoleClear()
	var name string
	var goods string
	var warehouse string
	var large float64
	var checkIn string
	var totalDays int64
	var price float64
	var confirmation string
	fmt.Println()
	fmt.Printf("%s\n", "Customer Order Form")
	fmt.Printf("%s\n", strings.Repeat("-", 30))
	scanner := bufio.NewReader(os.Stdin)
	fmt.Print("Customer Name : ")
	sName, _ := scanner.ReadString('\n')
	name = strings.TrimSpace(sName)
	fmt.Print("Goods : ")
	sGoods, _ := scanner.ReadString('\n')
	goods = strings.TrimSpace(sGoods)
	fmt.Print("Large (/m2): ")
	sLarge, _ := scanner.ReadString('\n')
	large, _ = strconv.ParseFloat(strings.TrimSpace(sLarge), 64)
	fmt.Print("Warehouse ID : ")
	sWarehouse, _ := scanner.ReadString('\n')
	warehouse = strings.TrimSpace(sWarehouse)
	fmt.Print("Date Check In (YYYY/mm/dd) : ")
	sCheckIn, _ := scanner.ReadString('\n')
	checkIn = strings.TrimSpace(sCheckIn)
	fmt.Print("Total days rent : ")
	sTotalDays, _ := scanner.ReadString('\n')
	totalDays, _ = strconv.ParseInt(strings.TrimSpace(sTotalDays), 10, 64)
	fmt.Print("Price (Rp.): ")
	sPrice, _ := scanner.ReadString('\n')
	price, _ = strconv.ParseFloat(strings.TrimSpace(sPrice), 64)

	fmt.Println("Save customer order ? :Y/N")
	fmt.Scanln(&confirmation)

	if confirmation == "Y" {
		newOrder := models.NewCustomerOrder(name, goods, large, warehouse, checkIn, totalDays, price)
		bd.customerOrderService.RegisterNewCustomerOrder(&newOrder)
	}
	utils.ConsoleClear()
	backToMainMenu()
}
