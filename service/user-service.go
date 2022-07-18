package service

import (
	"fmt"
	"webapp/e-commerce/config"
	"webapp/e-commerce/models"
)

func AddUser(user *models.User) error {
	result := config.DB.Create(&user)
	return result.Error
}

func GetUsers() ([]models.User, error) {
	var users []models.User
	result := config.DB.Find(&users)
	return users, result.Error
}

func UpdateUser(id string, user models.User) error {
	var res models.User
	err := config.DB.Where("user_Id= ?", id).First(&res).Error
	if err != nil {
		return fmt.Errorf("user with UserID [%s] does not exist", id)
	}
	res.Username = user.Username
	res.EmailID = user.EmailID
	res.Address = user.Address
	config.DB.Save(&res)
	return nil
}

func DeleteUser(id string) error {
	user := models.User{UserID: id}
	result := config.DB.Delete(&user)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected < 1 {
		return fmt.Errorf("user with UserID [%s] does not exist", id)
	}
	return nil
}

func GetUserWithOrders(id string) (models.User, error) {
	var user models.User
	err := config.DB.Where("user_id= ?", id).Preload("Orders").First(&user).Error
	return user, err
}

func AddOrder(orders []models.Order) error {
	result := config.DB.Create(&orders)
	return result.Error
}

func GetUserOrder(orderid string) (models.Order, error) {
	var order models.Order
	err := config.DB.Where("order_id= ?", orderid).First(&order).Error
	return order, err
}

func DeleteOrder(orderid string) error {
	order := models.Order{OrderID: orderid}
	result := config.DB.Delete(&order)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected < 1 {
		return fmt.Errorf("order with OrderID [%s] does not exist", orderid)
	}
	return nil
}