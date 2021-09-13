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
func Start(cfg *config.Config, db *postgresdb.PgDb) *echo.Echo {
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


	//register all repo
	productRepo 			:= postgresrepository.NewProductRepository(db.Pg)
	orderCartRepo			:= postgresrepository.NewOrderCartRepository(db.Pg)
	orderCartDetailRepo 	:= postgresrepository.NewOrderCartDetailRepository(db.Pg)
	userRepo				:= postgresrepository.NewUserRepository(db.Pg)

	//register all usecase
	productUseCase 	:= usecase.NewProductUseCase(productRepo)
	orderCartUseCase 	:= usecase.NewOrderCartUseCase(orderCartRepo, orderCartDetailRepo, userRepo, productRepo)

	// register handler
	h := handler.NewHandler(handler.HandlerConfig{
		Validator:   v,
		ProductUseCase: productUseCase,
		OrderUseCase: orderCartUseCase,
	})

	e.GET("/product/:id", h.FindProductById)
	e.POST("/cart/add", h.AddToCart)

	return e
}


// Stop app
func Stop() {
	// TODO: change graceful shutdown implementation
	fmt.Printf("Stop app at %s", time.Now())
}


