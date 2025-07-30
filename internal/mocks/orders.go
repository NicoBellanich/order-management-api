package mocks

import "github.com/nicobellanich/order-management-api/internal/domain"

var Orders = domain.Orders{
	Orders: []domain.Order{
		{
			Id:      1,
			Content: []string{"Mouse", "Keyboard"},
			Price:   150,
		},
		{
			Id:      2,
			Content: []string{"Monitor"},
			Price:   300,
		},
		{
			Id:      3,
			Content: []string{"Laptop Stand", "HDMI Cable"},
			Price:   90,
		},
	},
}
