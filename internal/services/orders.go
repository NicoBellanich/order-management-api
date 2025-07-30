package services

import (
	"github.com/nicobellanich/order-management-api/internal/domain"
	"github.com/nicobellanich/order-management-api/internal/platform/repository"
)

type IOrderService interface {
	GetByID(id int) (*domain.Order, error)
	GetAll() (*domain.Orders, error)
	Create(order *domain.Order) error
	Update(id int, order *domain.Order) error
	Delete(id int) error
}

type OrderService struct {
	ordersRepository repository.IOrdersRepository
}

func NewOrderService(or repository.IOrdersRepository) IOrderService {
	return &OrderService{ordersRepository: or}
}

func (os *OrderService) GetByID(id int) (*domain.Order, error) {
	return os.ordersRepository.GetByID(id)
}
func (os *OrderService) GetAll() (*domain.Orders, error) {
	return os.ordersRepository.GetAll()
}
func (os *OrderService) Create(order *domain.Order) error {
	return nil
}
func (os *OrderService) Update(id int, order *domain.Order) error {
	return nil
}
func (os *OrderService) Delete(id int) error {
	return nil
}
