package domain

type Orders struct {
	Orders []Order
}

func (o *Orders) GetAllOrders() []Order {
	return o.Orders
}

func (o *Orders) GetTotalPrice() (total int) {
	for _, order := range o.Orders {
		total += order.Price
	}
	return
}

type Order struct {
	Id      int
	Content []string
	Price   int
}
