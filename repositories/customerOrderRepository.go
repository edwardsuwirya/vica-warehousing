package repositories

import (
	"crypto/md5"
	"fmt"
	"time"
	"warehousing/models"
)

type customerOrderRepository struct {
	customerCollection []*models.CustomerOrder
}

func NewCustomerOrderRepository() ICustomerOrderRepository {
	customerCollection := make([]*models.CustomerOrder, 0)
	return &customerOrderRepository{
		customerCollection,
	}
}
func (ord *customerOrderRepository) AddNewCustomerOrder(customerOrder *models.CustomerOrder) {
	data := []byte(time.Now().String())
	customerOrder.Id = fmt.Sprintf("%x", md5.Sum(data))
	ord.customerCollection = append(ord.customerCollection, customerOrder)
}
