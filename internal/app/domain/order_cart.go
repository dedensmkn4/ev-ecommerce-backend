package domain

import "time"

type (
	// AddToCartPayload AddToCartParam parameter for "/cart/add" endpoint
	AddToCartPayload struct {
		UserIDStr     string `json:"user_id"`
		UserID        int
		ProductCodes  []string `json:"product_code"`
		QuantitiesStr []string `json:"qty"`
		Quantities    []int
	}

	// AddToCartResponse is response for "/cart/add" endpoint
	AddToCartResponse struct {
		CartID int `json:"cart_id"`
	}

	// CheckoutPayload CheckoutParam is param for checkout endpoint
	CheckoutPayload struct {
		CartIDStr string `form:"cart_id"`
		CartID    int
	}

	// CartItem is data for cart item
	CartItem struct {
		ProductCode string
		Qty         int
		Date        time.Time
	}

	// CartData containing cart data including its details
	CartData struct {
		CartID int
		UserID int
		Date   time.Time
		Items  []CartItem
	}

	// CheckoutResponse is response for checkout endpoint
	CheckoutResponse struct {
		OrderID int `json:"order_id"`
	}
)
