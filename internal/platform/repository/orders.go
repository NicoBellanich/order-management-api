package repository

import (
	"errors"
	"fmt"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/nicobellanich/order-management-api/config"
	"github.com/nicobellanich/order-management-api/internal/domain"
	"github.com/nicobellanich/order-management-api/internal/mocks"
)

type IOrdersRepository interface {
	GetByID(id int) (*domain.Order, error)
	GetAll() (*domain.Orders, error)
	Create(order *domain.Order) error
	Update(id int, order *domain.Order) error
	Delete(id int) error
}

func NewOrdersRepository(conf *config.Config) (IOrdersRepository, error) {
	var or IOrdersRepository

	switch {
	case conf.IsProdEnv():
		or = newPordRepository()
	case conf.IsTestEnv():
		or = newTestRepository()
	case conf.IsLocalEnv():
		or = newInmemoryRepository()
	default:
		return nil, errors.New("something went wrong loading environments for creating MessageRepository")
	}

	return or, nil
}

func newPordRepository() IOrdersRepository {
	return nil
}

func newTestRepository() IOrdersRepository {
	return nil
}

type InmemoryRepository struct {
	mu     sync.Mutex
	orders domain.Orders
}

func newInmemoryRepository() IOrdersRepository {
	return &InmemoryRepository{
		orders: mocks.Orders,
	}
}

func (ir *InmemoryRepository) GetByID(id int) (*domain.Order, error) {
	ir.mu.Lock()
	defer ir.mu.Unlock()

	for _, order := range ir.orders.Orders {
		if order.Id == id {
			return &order, nil
		}
	}

	return nil, fiber.NewError(fiber.ErrNotFound.Code, fmt.Sprintf("order not found %v", id))
}
func (ir *InmemoryRepository) GetAll() (*domain.Orders, error) {
	return &ir.orders, nil
}
func (ir *InmemoryRepository) Create(order *domain.Order) error {
	return nil
}
func (ir *InmemoryRepository) Update(id int, order *domain.Order) error {
	return nil
}
func (ir *InmemoryRepository) Delete(id int) error {
	return nil
}
