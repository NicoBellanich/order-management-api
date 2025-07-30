package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/nicobellanich/order-management-api/internal/services"
)

type IOrderController interface {
	GetOrderByID(c *fiber.Ctx) error
	GetAllOrders(c *fiber.Ctx) error
	CreateOrder(c *fiber.Ctx) error
	UpdateOrder(c *fiber.Ctx) error
	DeleteOrder(c *fiber.Ctx) error
}

type OrderController struct {
	orderService services.IOrderService
}

func NewOrderController(os services.IOrderService) IOrderController {
	return &OrderController{orderService: os}
}

func (oc *OrderController) GetOrderByID(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	order, err := oc.orderService.GetByID(id)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(order)
}

func (oc *OrderController) GetAllOrders(c *fiber.Ctx) error {
	orders, err := oc.orderService.GetAll()
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(orders)
}

func (oc *OrderController) CreateOrder(c *fiber.Ctx) error {
	return nil
}
func (oc *OrderController) UpdateOrder(c *fiber.Ctx) error {
	return nil
}
func (oc *OrderController) DeleteOrder(c *fiber.Ctx) error {
	return nil
}
