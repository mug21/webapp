package routes

import (
	"fmt"
	"net/http"
	"webapp/e-commerce/models"
	"webapp/e-commerce/service"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

func RegisterRoutes() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	{
		api.POST("/addUser", handleAddUser)
		api.GET("/getUsers", handleGetUsers)
		api.PUT("/updateUser/:userId", handleUpdateUser)
		api.DELETE("/deleteUser/:userId", handleDeleteUser)

		api.GET("/user/:userId/orders", handleGetUserWithOrders)
		api.POST("/user/addOrder", handleAddOrder)
		api.GET("/user/:userId/getOrder/:orderId", handleGetUserOrder)
		api.DELETE("/user/deleteOrder/:orderId", handleDeleteOrder)
	}
	return r
}

func handleAddUser(ctx *gin.Context) {
	fmt.Println("Inside Add User handler...")
	var user models.User
	ctx.BindJSON(&user)
	// Validate request payload
	if err := validateRequest(user); err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"Error": err.Error()})
		return
	}
	err := service.AddUser(&user)
	if err != nil {
		ctx.JSON(http.StatusNotFound, map[string]interface{}{"Error": err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, "Message: User created successfully!!")
	}
}

func handleGetUsers(ctx *gin.Context) {
	fmt.Println("Inside Get All Users handler...")
	res, err := service.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusNotFound, map[string]interface{}{"Error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, res)
	}
}

func handleUpdateUser(ctx *gin.Context) {
	fmt.Println("Inside Update User handler...")
	id := ctx.Params.ByName("userId")
	var user models.User
	ctx.BindJSON(&user)
	err := service.UpdateUser(id, user)
	if err != nil {
		ctx.JSON(http.StatusNotFound, map[string]interface{}{"Error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, map[string]interface{}{"Message": "User updated successfully!!"})
	}
}

func handleDeleteUser(ctx *gin.Context) {
	fmt.Println("Inside Delete user handler...")
	id := ctx.Params.ByName("userId")
	err := service.DeleteUser(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, map[string]interface{}{"Error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, map[string]interface{}{"Message": "User deleted successfully!!"})
	}
}

func handleGetUserWithOrders(ctx *gin.Context) {
	fmt.Println("Inside GetUserWithOrder handler...")
	id := ctx.Params.ByName("userId")
	res, err := service.GetUserWithOrders(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, map[string]interface{}{"Error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, map[string]interface{}{"Message": "Successfully fetched user with orders", "Response": res})
	}
}

func handleAddOrder(ctx *gin.Context) {
	fmt.Println("Inside AddOrder handler...")
	var orders []models.Order
	ctx.BindJSON(&orders)
	err := service.AddOrder(orders)
	if err != nil {
		ctx.JSON(http.StatusNotFound, map[string]interface{}{"Error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, "Message: New orders added successfully!!")
	}
}

func handleGetUserOrder(ctx *gin.Context) {
	fmt.Println("Inside GetUserOrder handler...")
	orderid := ctx.Params.ByName("orderId")
	res, err := service.GetUserOrder(orderid)
	if err != nil {
		ctx.JSON(http.StatusNotFound, map[string]interface{}{"Error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, res)
	}
}

func handleDeleteOrder(ctx *gin.Context) {
	fmt.Println("Inside DeleteOrder handler...")
	orderid := ctx.Params.ByName("orderId")
	err := service.DeleteOrder(orderid)
	if err != nil {
		ctx.JSON(http.StatusNotFound, map[string]interface{}{"Error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, map[string]interface{}{"Message": "Order deleted successfully!!"})
	}
}

func validateRequest(val interface{}) error {
	v := validator.New()
	err := v.Struct(val)
	return err
}
