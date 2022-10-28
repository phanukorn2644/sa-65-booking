package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/phanukorn2644/sa-65-example/entity"

	"net/http"
)

// POST /Set_of_furnitures

func CreateSet_of_furniture(c *gin.Context) {

	var set_of_furnitures entity.Set_of_furniture

	if err := c.ShouldBindJSON(&set_of_furnitures); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if err := entity.DB().Create(&set_of_furnitures).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": set_of_furnitures})

}

// GET /Set_of_furniture/:id

func GetSet_of_furniture(c *gin.Context) {

	var set_of_furnitures entity.Set_of_furniture

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM Set_of_furnitures WHERE id = ?", id).Scan(&set_of_furnitures).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": set_of_furnitures})

}

// GET /Set_of_furnitures

func ListSet_of_furnitures(c *gin.Context) {

	var set_of_furnitures []entity.Set_of_furniture

	if err := entity.DB().Raw("SELECT * FROM Set_of_furnitures").Scan(&set_of_furnitures).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": set_of_furnitures})

}

// DELETE /Set_of_furnitures/:id

func DeleteSet_of_furniture(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM Set_of_furnitures WHERE id = ?", id); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Set_of_furniture not found"})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": id})

}

// PATCH /Set_of_furnitures

func UpdateSet_of_furniture(c *gin.Context) {

	var set_of_furnitures entity.Set_of_furniture

	if err := c.ShouldBindJSON(&set_of_furnitures); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if tx := entity.DB().Where("id = ?", set_of_furnitures.ID).First(&set_of_furnitures); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Set_of_furniture not found"})

		return

	}

	if err := entity.DB().Save(&set_of_furnitures).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": set_of_furnitures})

}
