package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/devianwahyu/efishery/domain/model"
	"github.com/devianwahyu/efishery/domain/service"
	"github.com/devianwahyu/efishery/utils"
	"github.com/labstack/echo/v4"
)

type ICustomerController interface {
	GetCustomersController(c echo.Context) error
	GetCustomersByIDController(c echo.Context) error
	AddCustomerController(c echo.Context) error
	UpdateCustomerController(c echo.Context) error
	DeleteCustomerController(c echo.Context) error
}

type ControllerCustomer struct {
	serv service.CustomerService
}

func NewControllerCustomer(serv service.CustomerService) *ControllerCustomer {
	return &ControllerCustomer{serv}
}

func (cc *ControllerCustomer) GetCustomersController(c echo.Context) error {
	customers, err := cc.serv.GetCustomersService()

	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusNotFound, utils.ErrorResponse{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	log.Println("Get customers success")
	return c.JSON(http.StatusOK, utils.SuccessResponseArray{
		Status:  "success",
		Message: "Get customers success",
		Data:    customers,
	})
}

func (cc *ControllerCustomer) GetCustomerByIDController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	customer, errGet := cc.serv.GetCustomerByIDService(uint(id))

	if errGet != nil {
		log.Println(errGet.Error())
		return c.JSON(http.StatusNotFound, utils.ErrorResponse{
			Status:  "failed",
			Message: errGet.Error(),
		})
	}

	log.Println("Get customer success")
	return c.JSON(http.StatusOK, utils.SuccessResponse{
		Status:  "success",
		Message: "Get customer success",
		Data:    customer,
	})
}

func (cc *ControllerCustomer) AddCustomerController(c echo.Context) error {
	var newCustomer model.Customer
	if err := c.Bind(&newCustomer); err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	customer, errCreate := cc.serv.AddCustomerService(newCustomer)

	if errCreate != nil {
		log.Println(errCreate.Error())
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Status:  "failed",
			Message: errCreate.Error(),
		})
	}

	log.Println("Add customer success")
	return c.JSON(http.StatusCreated, utils.SuccessResponse{
		Status:  "success",
		Message: "Add customer success",
		Data:    customer,
	})
}

func (cc *ControllerCustomer) UpdateCustomerController(c echo.Context) error {
	id, errId := strconv.Atoi(c.Param("id"))

	if errId != nil {
		log.Println(errId.Error())
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Status:  "failed",
			Message: errId.Error(),
		})
	}

	var newCustomerData model.Customer
	if err := c.Bind(&newCustomerData); err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	customer, errUpdate := cc.serv.UpdateCustomerService(uint(id), newCustomerData)

	if errUpdate != nil {
		log.Println(errUpdate.Error())
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Status:  "failed",
			Message: errUpdate.Error(),
		})
	}

	log.Println("Update customer success")
	return c.JSON(http.StatusOK, utils.SuccessResponse{
		Status:  "success",
		Message: "Update customer success",
		Data:    customer,
	})
}

func (cc *ControllerCustomer) DeleteCustomerController(c echo.Context) error {
	id, errId := strconv.Atoi(c.Param("id"))

	if errId != nil {
		log.Println(errId.Error())
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Status:  "failed",
			Message: errId.Error(),
		})
	}

	message, errDelete := cc.serv.DeleteCustomerService(uint(id))

	if errDelete != nil {
		log.Println(errDelete.Error())
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Status:  "failed",
			Message: errDelete.Error(),
		})
	}

	type response struct {
		status  string
		message string
	}

	log.Println("Delete customer success")
	return c.JSON(http.StatusOK, response{
		status:  "success",
		message: message,
	})
}
