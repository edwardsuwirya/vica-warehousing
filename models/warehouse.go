package models

import "fmt"

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

func (w *Warehouse) ToString() string {
	return fmt.Sprintf("%s,%s,%s,%f,%s,%f\n", w.Kode, w.Name, w.Address, w.Large, w.Information, w.Price)
}
