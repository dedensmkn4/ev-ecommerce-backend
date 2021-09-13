package domain

type (
	OrderCartDetail struct {
		ID    			string
		ProductCode  	string
		Quantity  		string
		OrderId 		string
		Date 			string
	}
)

var (
	OrderCartDetailTableName = "tbl_order__cart_detail"
	OrderCartDetailTable     = struct {
		ID    			string
		ProductCode  	string
		Quantity  		string
		OrderId 		string
		Date 			string
	}{
		ID:     		"id",
		ProductCode:   	"product_code",
		Quantity: 		"quantity",
		OrderId: 		"order_id",
		Date: 			"date",
	}
)

