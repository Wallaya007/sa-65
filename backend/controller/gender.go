package controller

import (
	"net/http"

	"github.com/Wallaya007/sa-65-example/entity"
	"github.com/gin-gonic/gin"
)

// LIST /genders
func ListGenders(c *gin.Context) {
	var genders []entity.Gender
	if err := entity.DB().Raw("SELECT * FROM genders").Find(&genders).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": genders})
}


// // POST /genders
// func CreateGender(c *gin.Context) {
// 	var gender entity.Gender
// 	if err := c.ShouldBindJSON(&gender); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if err := entity.DB().Create(&gender).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusCreated, gin.H{"data": gender})
// }

// // GET /gender/:id
// func GetGender(c *gin.Context) {
// 	var gender entity.Gender

// 	id := c.Param("id")
// 	if tx := entity.DB().Where("id = ?", id).First(&gender); tx.RowsAffected == 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "gender not found"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": gender})
// }

// // LIST /genders
// func ListGenders(c *gin.Context) {
// 	var genders []entity.Gender
// 	if err := entity.DB().Raw("SELECT * FROM genders").Find(&genders).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": genders})
// }

// // DELETE /genders/:id
// func DeleteGender(c *gin.Context) {
// 	id := c.Param("id")
// 	if tx := entity.DB().Exec("DELETE FROM genders WHERE id = ?", id); tx.RowsAffected == 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "gender not found"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": id})
// }

// // PATCH /genders
// func UpdateGender(c *gin.Context) {
// 	var gender entity.Gender
// 	if err := c.ShouldBindJSON(&gender); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if tx := entity.DB().Where("id = ?", gender.ID).First(&gender); tx.RowsAffected == 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "gender not found"})
// 		return
// 	}

// 	if err := entity.DB().Save(&gender).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": gender})
// }
