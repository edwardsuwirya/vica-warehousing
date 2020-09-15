package usecases

import (
	"warehousing/models"
	"warehousing/repositories"
)

type customerOrderService struct {
	r repositories.ICustomerOrderRepository
}

func NewCustomerOrderService(repo repositories.ICustomerOrderRepository) ICustomerOrderService {
	return &customerOrderService{r: repo}
}

func (bs *customerOrderService) RegisterNewCustomerOrder(b *models.CustomerOrder) {
	bs.r.AddNewCustomerOrder(b)
}
