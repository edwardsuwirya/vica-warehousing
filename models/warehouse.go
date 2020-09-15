package models

type Warehouse struct {
	Kode        string
	Name        string
	Address     string
	Large       float64
	Information string
	Price       float64
}

func NewWarehouse(name string, address string, large float64, information string, price float64) Warehouse {
	return Warehouse{
		Name:        name,
		Address:     address,
		Large:       large,
		Information: information,
		Price:       price,
	}
}
