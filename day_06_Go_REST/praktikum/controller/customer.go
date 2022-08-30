package controller

import (
	"log"
	"net/http"

	"github.com/devianwahyu/go_rest/model"
	"github.com/labstack/echo/v4"
)

var customers = []model.Customer{
	{
		ID:   "1",
		Name: "Joko",
		Age:  20,
	},
	{
		ID:   "2",
		Name: "Ucup",
		Age:  22,
	},
}

func UseSubroute(group *echo.Group) {
	group.GET("/", getCustomers)
	group.GET("/:id", getCustomerByID)
	group.POST("/", addCustomer)
	group.PUT("/:id", updateCustomer)
	group.DELETE("/:id", deleteCustomer)
}

func getCustomers(c echo.Context) error {
	customerList := customers
	log.Println("Success to get list of customers")
	return c.JSON(http.StatusOK, customerList)
}

func getCustomerByID(c echo.Context) error {
	id := c.Param("id")
	var customer model.Customer

	for _, value := range customers {
		if value.ID == id {
			customer = value
		}
	}

	if customer.ID == "" {
		log.Println("Failed to get customer with id: " + id)
		return c.JSON(http.StatusNotFound, model.Error{
			Message: "Not found customer with id: " + id,
		})
	}

	log.Println("Success to get customer with id: " + id)
	return c.JSON(http.StatusOK, customer)
}

func addCustomer(c echo.Context) error {
	newUser := new(model.Customer)
	if err := c.Bind(newUser); err != nil {
		log.Println("Failed to add new customer")
		return c.JSON(http.StatusBadRequest, model.Error{
			Message: "Failed to add new user",
		})
	}

	idUsed := false
	for _, value := range customers {
		if value.ID == newUser.ID {
			idUsed = true
			break
		}
	}

	if idUsed {
		log.Println("Failed to add new customer with id: " + newUser.ID)
		return c.JSON(http.StatusConflict, model.Error{
			Message: "User with id " + newUser.ID + " already exist",
		})
	}

	customers = append(customers, *newUser)

	log.Println("Success to add new customer with id: " + newUser.ID)
	return c.JSON(http.StatusOK, customers)
}

func updateCustomer(c echo.Context) error {
	id := c.Param("id")

	newUser := new(model.Customer)
	if err := c.Bind(newUser); err != nil {
		log.Println("Failed to update customer with id: " + id)
		return c.JSON(http.StatusBadRequest, model.Error{
			Message: "Failed to update user id: " + id,
		})
	}

	for i, value := range customers {
		if value.ID == id {
			customers[i] = *newUser
		}
	}

	log.Println("SUccess to update customer with id: " + id)
	return c.JSON(http.StatusOK, newUser)
}

func deleteCustomer(c echo.Context) error {
	id := c.Param("id")
	message := ""

	for i := range customers {
		if customers[i].ID == id {
			customers = append(customers[:i], customers[i+1:]...)
			break
		}
	}

	if message != "" {
		log.Println("Failed to delete customer with id: " + id)
		return c.JSON(http.StatusNotFound, model.Error{
			Message: "Not found customer with id: " + id,
		})
	}

	log.Println("Success to delete customer with id: " + id)
	return c.JSON(http.StatusOK, model.Success{
		Message: "Customer with id " + id + " deleted",
	})
}
