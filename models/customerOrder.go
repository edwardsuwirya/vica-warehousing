package models

type CustomerOrder struct {
	Id           string
	Name         string
	Goods        string
	LargeUsed    float64
	warehouseId  string
	DateOfEntry  string
	TotalDayRent int64
	TotalPrice   float64
}

func NewCustomerOrder(name string, goods string, largeUsed float64, warehouse string, dateOfEntry string, totalDayRent int64, totalPrice float64) CustomerOrder {
	return CustomerOrder{
		Name:         name,
		Goods:        goods,
		LargeUsed:    largeUsed,
		warehouseId:  warehouse,
		DateOfEntry:  dateOfEntry,
		TotalDayRent: totalDayRent,
		TotalPrice:   totalPrice,
	}
}
