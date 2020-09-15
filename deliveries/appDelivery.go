package deliveries

import (
	"fmt"
	"sort"
	"strings"
	"warehousing/config"
)

type callbackFn func()

type AppDelivery struct {
	warehouseDelivery     *WarehouseDelivery
	customerOrderDelivery *CustomerOrderDelivery
}

func NewAppDelivery(c *config.AppConfig) *AppDelivery {
	return &AppDelivery{
		warehouseDelivery:     NewWarehouseDelivery(),
		customerOrderDelivery: NewCustomerOrderDelivery(),
	}
}

func (bd *AppDelivery) Create() {
	var isExist = false
	var userChoice string

	bd.mainMenuForm()
	for isExist == false {
		fmt.Printf("\n%s", "Your Choice: ")
		fmt.Scan(&userChoice)
		switch {
		case userChoice == "01":
			bd.warehouseDelivery.RegistrationWarehouseForm(bd.mainMenuForm)
		case userChoice == "02":
			bd.warehouseDelivery.ListWarehouseForm(bd.mainMenuForm)
		case userChoice == "03":
			bd.customerOrderDelivery.CustomerOrderForm(bd.mainMenuForm)
		case userChoice == "q":
			isExist = true
		default:
			fmt.Println("Unknown Menu Code")

		}
	}
}

func (bd *AppDelivery) menuChoiceOrdered(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func (bd *AppDelivery) mainMenuForm() {
	var appMenu = map[string]string{
		"01": "Create new warehouse",
		"02": "Warehouse List",
		"03": "Customer Order Form",
		"q":  "Quit aplication",
	}
	fmt.Printf("%s\n", strings.Repeat("*", 30))
	fmt.Printf("%26s\n", "Warehouse Application")
	fmt.Printf("%s\n", strings.Repeat("*", 30))
	for _, menuCode := range bd.menuChoiceOrdered(appMenu) {
		fmt.Printf("%s. %s\n", menuCode, appMenu[menuCode])
	}
}

func (bd *AppDelivery) Run() {
	bd.Create()
}
