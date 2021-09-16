package usecase

import (
	"context"
	"database/sql"
	"github.com/dedensmkn4/ev-ecommerce-backend/internal/app/domain"
	"github.com/dedensmkn4/ev-ecommerce-backend/internal/app/domain/port"
	"github.com/dedensmkn4/ev-ecommerce-backend/internal/app/infra/postgresdb"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-rest-server/pkg/sqkit"
	"net/http"
	"time"
)

const CHECKOUT = "checkout"

type orderCartUseCaseImpl struct {
	orderCartRepository port.OrderCartRepository
	orderCartDetailRepository port.OrderCartDetailRepository
	userRepository port.UserRepository
	productRepository port.ProductRepository
	cartItemRepository port.CartItemRepository
	pgDb *postgresdb.PgDb
}

func NewOrderCartUseCase(orderRepository port.OrderCartRepository,
	orderCartDetailRepository port.OrderCartDetailRepository,
	userRepository port.UserRepository,
	productRepository port.ProductRepository,
	cartItemRepository port.CartItemRepository, pgDb *postgresdb.PgDb)  port.OrderCartUseCase{
	return &orderCartUseCaseImpl{
		orderCartRepository: orderRepository,
		orderCartDetailRepository: orderCartDetailRepository,
		userRepository: userRepository,
		productRepository: productRepository,
		cartItemRepository: cartItemRepository,
		pgDb: pgDb,
	}
}

func (o orderCartUseCaseImpl) AddCart(ctx context.Context, param domain.AddCartItemPayload) (cartResponse *domain.AddCartItemResponse, err error) {

	_, err = o.validatedUser(ctx, param.UserId)

	if err != nil{
		return cartResponse, err
	}

	dbTx, err := o.pgDb.BeginTrx(ctx , nil)
	if err != nil {
		log.Error(err)
		return cartResponse, echo.NewHTTPError(http.StatusInternalServerError, "Failed Add to Cart")
	}

	//product checking
	_, err = o.validatedProduct(ctx, dbTx, param.ProductId, param.Quantity)
	if err != nil {
		return cartResponse, err
	}

	cartId, err := o.cartItemRepository.Insert(ctx, dbTx, o.toCartItem(param))
	if err != nil {
		_ = dbTx.Rollback()
		return cartResponse, echo.NewHTTPError(http.StatusInternalServerError, "Failed Add to Cart")
	}

	cartResponse = &domain.AddCartItemResponse{
		UserID : param.UserId,
	}

	cartResponse.CartItemDetail = append(cartResponse.CartItemDetail, domain.CartItemDetailResponse{
		ID: cartId,
		ProductId: param.ProductId,
		Quantity: param.Quantity,
		Date: time.Now(),
	})

	return cartResponse, dbTx.Commit()
}

func (o orderCartUseCaseImpl) Checkout(ctx context.Context, param domain.OrderCheckoutPayload) (checkoutResponse *domain.OrderCheckoutResponse, err error) {
	_, err = o.validatedUser(ctx, param.UserId)

	if err != nil{
		return checkoutResponse, err
	}

	dbTx, err := o.pgDb.BeginTrx(ctx , nil)
	if err != nil {
		log.Error(err)
		return checkoutResponse, echo.NewHTTPError(http.StatusInternalServerError, "Failed Add to Cart")
	}

	//insert order__cart
	orderId, err := o.orderCartRepository.Insert(ctx, dbTx, o.toOrderCart(param.UserId, CHECKOUT))

	if err != nil {
		log.Error(err)
		return checkoutResponse, echo.NewHTTPError(http.StatusInternalServerError, "Failed Create OrderCart")
	}

	checkoutResponse = &domain.OrderCheckoutResponse{
		OrderId: orderId,
		UserID: param.UserId,
		OrderStatus: CHECKOUT,
		Date: time.Now(),
	}

	//initialize total price
	var totalPrice int

	for _,carDetailParam := range param.OrderCartDetail {

		cartItemsDb, err := o.cartItemRepository.Find(ctx, sqkit.Eq{domain.CartItemTable.ID : carDetailParam.CartDetailId})
		if err != nil {
			return checkoutResponse, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		} else if len(cartItemsDb) < 1 {
			return checkoutResponse, echo.NewHTTPError(http.StatusNotFound, "CartDetailId Not Found")
		}

		if cartItemsDb[0].ProductId != carDetailParam.ProductId {
			_ = dbTx.Rollback()
			return checkoutResponse, echo.NewHTTPError(http.StatusBadRequest, "ProductId Not Valid")
		}


		if cartItemsDb[0].UserID != param.UserId {
			_ = dbTx.Rollback()
			return checkoutResponse, echo.NewHTTPError(http.StatusBadRequest, "UserId Not Valid")
		}

		//product checking
		productDb, err := o.validatedProduct(ctx, dbTx, carDetailParam.ProductId, carDetailParam.Quantity)
		if err != nil {
			return checkoutResponse, err
		}

		orderDetailId, err := o.orderCartDetailRepository.Insert(ctx, dbTx, o.toOrderCartDetail(orderId, productDb.Price, carDetailParam))
		if err != nil {
			_ = dbTx.Rollback()
			return checkoutResponse, echo.NewHTTPError(http.StatusInternalServerError, "Failed Add to Cart")
		}

		//decrease stock
		newStock := productDb.Stock - carDetailParam.Quantity
		productDb.Stock = newStock
		_,err = o.updateStockProduct(ctx, dbTx, carDetailParam.ProductId, productDb)
		if err != nil {
			_ = dbTx.Rollback()
			return checkoutResponse, echo.NewHTTPError(http.StatusInternalServerError, "Failed Decrease Stock Product")
		}

		//delete cart items
		_,err = o.cartItemRepository.Delete(ctx, dbTx, sqkit.Eq{domain.CartItemTable.ID :	carDetailParam.CartDetailId})
		if err != nil {
			_ = dbTx.Rollback()
			return checkoutResponse, echo.NewHTTPError(http.StatusInternalServerError, "Failed Delete Cart Items")
		}

		oderCartDetailResponse := domain.OrderCartDetailResponse{
			OrderDetailId	:	orderDetailId,
			ProductId 		: 	carDetailParam.ProductId,
			Quantity 		:	carDetailParam.Quantity,
			Price			:	carDetailParam.Price,
			Date			:	time.Now(),
		}

		checkoutResponse.OrderCartDetail = append(checkoutResponse.OrderCartDetail, oderCartDetailResponse)

		totalPrice += carDetailParam.Price*carDetailParam.Quantity
	}

	checkoutResponse.TotalPrice = totalPrice

	return checkoutResponse, dbTx.Commit()
}

func (o orderCartUseCaseImpl) validatedUser(ctx context.Context, paramUserId int) (users []*domain.User, err error) {
	users, err = o.userRepository.Find(ctx, sqkit.Eq{domain.UserTable.ID:paramUserId})
	if err != nil{
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}else if len(users) < 1{
		return nil, echo.NewHTTPError(http.StatusNotFound, "User Not Found")
	}
	return
}

func (o orderCartUseCaseImpl) validatedProduct(ctx context.Context, dbTx *sql.Tx, paramProductId int, paramQty int)(product *domain.Product, err error) {
	//product checking
	products, err := o.productRepository.Find(ctx, dbTx, true, sqkit.Eq{domain.ProductTable.ID :paramProductId})
	if err != nil {
		_ = dbTx.Rollback()
		return product, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	} else if len(products) < 1 {
		_ = dbTx.Rollback()
		return product, echo.NewHTTPError(http.StatusNotFound, "Product Not Found")
	}

	if products[0].Stock < paramQty {
		return product, echo.NewHTTPError(http.StatusBadRequest, "Product Out Of Stock")
	}
	return products[0], nil
}

func (o orderCartUseCaseImpl) updateStockProduct(ctx context.Context, dbTx *sql.Tx, paramProductId int, product *domain.Product) (int, error) {
	return o.productRepository.Update(ctx, dbTx, product, sqkit.Eq{domain.ProductTable.ID: paramProductId})
}

func(o orderCartUseCaseImpl) toCartItem(param domain.AddCartItemPayload) *domain.CartItem{
	return &domain.CartItem{
		UserID: param.UserId,
		ProductId: param.ProductId,
		Quantity: param.Quantity,
	}
}

func (o orderCartUseCaseImpl) toOrderCart(paramUserId int, paramStatus string) *domain.OrderCart{
	return &domain.OrderCart{
		UserId: paramUserId,
		OrderStatus: paramStatus,
	}
}

func(o orderCartUseCaseImpl) toOrderCartDetail(paramOrderId int, paramPrice int, param domain.OrderCartDetailPayload) *domain.OrderCartDetail{
	return &domain.OrderCartDetail{
		OrderId: paramOrderId,
		ProductId: param.ProductId,
		Quantity: param.Quantity,
		Price: paramPrice,
		Date: time.Now(),
	}
}

