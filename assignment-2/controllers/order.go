package controllers

import (
	"assignment-2/databases"
	"assignment-2/models"
	"fmt"
	"log"
	"net/http"

	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateOrder(c *gin.Context) {
	db := databases.GetDB()
	var data models.CreateOrder

	var items []models.Item
	if err := c.ShouldBindJSON(&data); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	fmt.Println(data)
	orders := models.Order{
		CustomerName: data.CustomerName,
		OrderedAt:    data.OrderedAt,
	}

	db.Create(&orders)
	id_order := orders.ID
	for _, v := range data.Item {
		item := models.Item{
			ItemCode:    v.ItemCode,
			Description: v.Description,
			Quantity:    v.Quantity,
			OrderID:     orders.ID,
		}

		items = append(items, item)
	}
	fmt.Println(items)
	result := db.Create(&items)
	log.Println(id_order, result.RowsAffected)

	var returnData interface{}

	returnData = models.Order{
		ID:           items[0].OrderID,
		CustomerName: data.CustomerName,
		OrderedAt:    time.Now(),
		Item:         items,
	}

	c.JSON(http.StatusOK, returnData)
}

func GetOrder(c *gin.Context) {
	db := databases.GetDB()

	orders := []models.Order{}

	err := db.Preload("Item").Find(&orders).Error

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK, orders)
}

func GetOrderByID(c *gin.Context) {
	db := databases.GetDB()

	order := models.Order{}

	var id_order = c.Param("orderID")
	err := db.Where("id = ?", id_order).Find(&order).Error
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	items := []models.Item{}

	err = db.Where("order_id = ?", id_order).Find(&items).Error
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	resp := models.CreateOrder{}
	resp.CustomerName = order.CustomerName
	resp.OrderedAt = order.OrderedAt
	resp.Item = items
	c.JSON(http.StatusOK, resp)
}

func UpdateOrder(c *gin.Context) {
	db := databases.GetDB()
	firstID := c.Param("orderID")
	secondID, _ := strconv.Atoi(firstID)

	var data = models.CreateOrder{}

	if err := c.ShouldBindJSON(&data); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	order := models.Order{
		ID:           uint(secondID),
		CustomerName: data.CustomerName,
		OrderedAt:    data.OrderedAt,
		Item:         data.Item,
	}
	db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&order)
	c.JSON(http.StatusOK, order)
}

func DeleteOrder(c *gin.Context) {
	db := databases.GetDB()
	firstID := c.Param("orderID")
	secondID, _ := strconv.Atoi(firstID)

	orders := models.Order{}
	items := models.Item{}

	err := db.Where("order_id = ?", secondID).Delete(&items).Error

	rows := db.Where("ID = ?", secondID).Delete(&orders).RowsAffected

	if rows == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Message": "Opss...Data Not Found",
		})
		return
	}
	if err != nil {
		log.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Message": "Data Success For Delete",
	})
}
