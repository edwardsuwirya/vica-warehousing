package repositories

import "warehousing/models"

type ICustomerOrderRepository interface {
	AddNewCustomerOrder(customerOrder *models.CustomerOrder)
}
