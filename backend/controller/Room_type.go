package controller

import (
	"github.com/phanukorn2644/sa-65-example/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// POST /Room_types

func CreateRoom_type(c *gin.Context) {

	var room_types entity.Room_type

	if err := c.ShouldBindJSON(&room_types); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if err := entity.DB().Create(&room_types).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": room_types})

}

// GET /Room_type/:id

func GetRoom_type(c *gin.Context) {

	var room_types entity.Room_type

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM Room_types WHERE id = ?", id).Scan(&room_types).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": room_types})

}

// GET /Room_types

func ListRoom_types(c *gin.Context) {

	var room_types []entity.Room_type

	if err := entity.DB().Raw("SELECT * FROM Room_types").Scan(&room_types).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": room_types})

}

// DELETE /Room_types/:id

func DeleteRoom_type(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM room_types WHERE id = ?", id); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Room_type not found"})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": id})

}

// PATCH /Room_types

func UpdateRoom_type(c *gin.Context) {

	var room_types entity.Room_type

	if err := c.ShouldBindJSON(&room_types); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if tx := entity.DB().Where("id = ?", room_types.ID).First(&room_types); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Room_type not found"})

		return

	}

	if err := entity.DB().Save(&room_types).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": room_types})

}
