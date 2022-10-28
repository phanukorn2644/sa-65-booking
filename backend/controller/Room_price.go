package controller

import (
	"github.com/phanukorn2644/sa-65-example/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// POST /Room_types

func CreateRoom_price(c *gin.Context) {

	var room_prices entity.Room_price

	if err := c.ShouldBindJSON(&room_prices); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if err := entity.DB().Create(&room_prices).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": room_prices})

}

// GET /Room_type/:id

func GetRoom_price(c *gin.Context) {

	var room_prices entity.Room_price

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM room_prices WHERE id = ?", id).Scan(&room_prices).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": room_prices})

}

// GET /Room_types

func ListRoom_prices(c *gin.Context) {

	var room_prices []entity.Room_price

	if err := entity.DB().Raw("SELECT * FROM Room_prices").Scan(&room_prices).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": room_prices})

}

// DELETE /Room_types/:id

func DeleteRoom_price(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM room_prices WHERE id = ?", id); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Room_prics not found"})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": id})

}

// PATCH /Room_types

func UpdateRoom_price(c *gin.Context) {

	var room_prices entity.Room_price

	if err := c.ShouldBindJSON(&room_prices); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if tx := entity.DB().Where("id = ?", room_prices.ID).First(&room_prices); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Room_price not found"})

		return

	}

	if err := entity.DB().Save(&room_prices).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": room_prices})

}
