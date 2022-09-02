package router

import (
	"github.com/devianwahyu/efishery/api/controller"
	"github.com/devianwahyu/efishery/domain/service"
	"github.com/devianwahyu/efishery/repository"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRoute(db *gorm.DB) *echo.Echo {
	e := echo.New()

	CustomerRepository := repository.NewCustomerRepository(db)
	CustomerService := service.NewCustomerService(*CustomerRepository)
	CustomerController := controller.NewControllerCustomer(*CustomerService)

	e.GET("/", CustomerController.GetCustomersController)
	e.GET("/:id", CustomerController.GetCustomerByIDController)
	e.POST("/", CustomerController.AddCustomerController)
	e.PUT("/:id", CustomerController.UpdateCustomerController)
	e.DELETE("/:id", CustomerController.DeleteCustomerController)

	return e
}
