package usecases

import "warehousing/models"

type ICustomerOrderService interface {
	RegisterNewCustomerOrder(b *models.CustomerOrder)
}
