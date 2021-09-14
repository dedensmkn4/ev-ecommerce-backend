package app

import (
	"fmt"
	"github.com/dedensmkn4/ev-ecommerce-backend/internal/app/adapter/postgresrepository"
	"github.com/dedensmkn4/ev-ecommerce-backend/internal/app/config"
	"github.com/dedensmkn4/ev-ecommerce-backend/internal/app/domain/usecase"
	"github.com/dedensmkn4/ev-ecommerce-backend/internal/app/handler"
	"github.com/dedensmkn4/ev-ecommerce-backend/internal/app/infra/postgresdb"
	"github.com/dedensmkn4/ev-ecommerce-backend/internal/app/middle"
	"github.com/dedensmkn4/ev-ecommerce-backend/pkg/validation"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)


// Start app
func Start(cfg *config.Config) *echo.Echo {
	e := echo.New()
	v := validation.New()

	readTimeout, _ := strconv.Atoi(cfg.ReadTimeout)
	writeTimeout, _ := strconv.Atoi(cfg.WriteTimeout)
	/* ***** ***** ***** ***** ***** */
	/* setup server
	/* ***** ***** ***** ***** ***** */
	e.HideBanner = true
	e.Server.ReadTimeout = time.Duration(readTimeout) * time.Second
	e.Server.WriteTimeout = time.Duration(writeTimeout) * time.Second
	e.HTTPErrorHandler = middle.HTTPCustomError
	e.Validator = v

	e.Use(middle.Logger())

	postgresDb := postgresdb.NewPgDb()
	db := postgresDb.OpenPostgres(cfg)

	//register all repo
	productRepo 			:= postgresrepository.NewProductRepository(db)
	orderCartRepo			:= postgresrepository.NewOrderCartRepository(db)
	orderCartDetailRepo 	:= postgresrepository.NewOrderCartDetailRepository(db)
	userRepo				:= postgresrepository.NewUserRepository(db)
	cartItemRepo			:= postgresrepository.NewCartItemRepository(db)

	//register all usecase
	productUseCase 	:= usecase.NewProductUseCase(productRepo, postgresDb)
	orderCartUseCase 	:= usecase.NewOrderCartUseCase(orderCartRepo, orderCartDetailRepo, userRepo, productRepo, cartItemRepo, postgresDb)

	// register handler
	h := handler.NewHandler(handler.HandlerConfig{
		Validator:   v,
		ProductUseCase: productUseCase,
		OrderCartUseCase: orderCartUseCase,
	})

	group := e.Group("/ev-ecommerce")

	//product
	group.GET("/product", h.GetAllProduct)
	group.GET("/product/:id", h.FindProductById)

	//order
	group.POST("/cart/add", h.AddToCart)
	group.POST("/order/checkout", h.OrderCheckout)

	return e
}


// Stop app
func Stop() {
	// TODO: change graceful shutdown implementation
	fmt.Printf("Stop app at %s", time.Now())
}


