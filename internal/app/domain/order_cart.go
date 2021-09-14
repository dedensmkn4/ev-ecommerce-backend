package domain

import "time"

type (

	OrderCart     struct {
		ID    				int
		UserId  			int
		TotalPrice  		int
		OrderStatus			string
		Date 				time.Time
		TimeLimitPayment	time.Time
	}

	OrderCheckoutPayload struct {
		UserId				int		`json:"userID" validate:"required"`
		OrderCartDetail[] 	OrderCartDetailPayload `json:"orderCartDetail" validate:"required"`
	}

	OrderCheckoutResponse struct {
		OrderId 			int						`json:"orderId"`
		UserID 				int						`json:"UserId"`
		Date   				time.Time				`json:"date"`
		TotalPrice			int						`json:"totalPrice"`
		OrderStatus			string					`json:"OrderStatus"`
		OrderCartDetail[]	OrderCartDetailResponse `json:"orderCartDetail"`
	}


)

var (
	OrderCartTableName = "tbl_order__cart"
	OrderCartTable     = struct {
		ID    				string
		UserId  			string
		TotalPrice  		string
		OrderStatus			string
		Date 				string
		TimeLimitPayment	string
	}{
		ID:     		"id",
		UserId:   		"user_id",
		TotalPrice: 	"total_price",
		OrderStatus: 	"order_status",
		Date: 			"date",
		TimeLimitPayment: "time_limit_payment",
	}
)
