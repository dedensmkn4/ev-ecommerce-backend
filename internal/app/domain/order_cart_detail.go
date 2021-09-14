package domain

import "time"

type (
	OrderCartDetail struct {
		ID    			int
		OrderDetailId  	int
		ProductId  		int
		Quantity  		int
		Price			int
		OrderId 		int
		Date 			time.Time
	}

	OrderCartDetailPayload struct {
		CartDetailId	int		`json:"cart_detail_id" validate:"required"`
		ProductId  		int 	`json:"productId"  validate:"required"`
		Quantity  		int 	`json:"quantity"  validate:"required"`
		Price  			int 	`json:"price"  validate:"required"`
	}

	OrderCartDetailResponse struct {
		OrderDetailId	int			`json:"orderDetailId"`
		ProductId  		int 		`json:"productId"`
		Quantity  		int 		`json:"quantity"`
		Price  			int 		`json:"price"`
		Date			time.Time	`json:"date"`
	}
)

var (
	OrderCartDetailTableName = "tbl_order__cart_detail"
	OrderCartDetailTable     = struct {
		ID    			string
		ProductId  		string
		Quantity  		string
		Price			string
		OrderId 		string
		Date 			string
	}{
		ID:     		"id",
		ProductId:   	"product_id",
		Quantity: 		"quantity",
		Price:			"price",
		OrderId: 		"order_id",
		Date: 			"date",
	}
)

