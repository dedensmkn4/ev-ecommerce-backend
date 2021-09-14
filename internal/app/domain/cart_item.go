package domain

import "time"

type (
	AddCartItemPayload struct {
		UserId			int	`json:"userID" validate:"required"`
		ProductId	  	int	`json:"productId" validate:"required"`
		Quantity  		int	`json:"quantity" validate:"required"`
	}

	CartItem struct {
		ID    			int 		`json:"id"`
		UserID    		int 		`json:"UserID"`
		ProductId	  	int			`json:"productId"`
		Quantity  		int			`json:"quantity"`
		Date 			time.Time	`json:"date"`
	}

	AddCartItemResponse struct {
		UserID    			int 					`json:"UserID"`
		CartItemDetail[]	CartItemDetailResponse  `json:"CartItemDetail"`
		Date 				time.Time				`json:"date"`
	}

	CartItemDetailResponse struct {
		ID    			int 		`json:"cartDetailId"`
		ProductId	  	int			`json:"productId"`
		Quantity  		int			`json:"quantity"`
		Date 			time.Time	`json:"date"`
	}
)

var (
	CartItemTableName = "tbl_cart_item"
	CartItemTable     = struct {
		ID    			string
		UserId			string
		ProductId	  	string
		Quantity  		string
		Date 			string
	}{
		ID:     	"id",
		UserId: 	"user_id",
		ProductId:  "product_id",
		Quantity:   "quantity",
		Date: 		"date",
	}
)