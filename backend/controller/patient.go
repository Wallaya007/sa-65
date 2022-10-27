package controller

import (
	"net/http"

	"github.com/Wallaya007/sa-65-example/entity"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

// POST /watch_videos
func CreatePatient(c *gin.Context) {

	var patient entity.Patient
	var disease entity.Disease
	var blood entity.Blood
	var gender entity.Gender
	var title entity.Title
	var user entity.User

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", patient.TitleID).First(&title); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "title not found"})
		return
	}

	if tx := entity.DB().Where("id = ?", patient.GenderID).First(&gender); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gender not found"})
		return
	}

	if tx := entity.DB().Where("id = ?", patient.BloodID).First(&blood); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "blood not found"})
		return
	}

	if tx := entity.DB().Where("id = ?", patient.DiseaseID).First(&disease); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "disease not found"})
		return
	}

	if tx := entity.DB().Where("id = ?", patient.UserID).First(&user); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	// สร้าง Patient
	pt := entity.Patient{ 
		Title:        title,               // โยงความสัมพันธ์กับ Entity Title
		Gender:       gender,               // โยงความสัมพันธ์กับ Entity Gender
		Blood:        blood,                // โยงความสัมพันธ์กับ Entity Blood
		Disease:      disease,              // โยงความสัมพันธ์กับ Entity Disease
		User: 		  user,
		PersonalID:   patient.PersonalID,   // ตั้งค่าฟิลด์ personalID
		Allergy:      patient.Allergy,      // ตั้งค่าฟิลด์ allergy
		Tel:          patient.Tel,          // ตั้งค่าฟิลด์ tel
		BirthdayTime: patient.BirthdayTime, // ตั้งค่าฟิลด์ birthdayTime

	}

	// ขั้นตอนการ validate ที่นำมาจาก unit test
	if _, err := govalidator.ValidateStruct(pt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 13: บันทึก
	if err := entity.DB().Create(&pt).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": pt})
}

// // GET /watchvideo/:id
// func GetPatient(c *gin.Context) {
// 	var patient entity.Patient
// 	id := c.Param("id")
// 	if err := entity.DB().Preload("User").Preload("Disease").Preload("Blood").Preload("Gender").Preload("Title").Raw("SELECT * FROM patients WHERE id = ?", id).Find(&patient).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"data": patient})
// }

// GET /watch_videos
func ListPatient(c *gin.Context) {
	var patient []entity.Patient
	if err := entity.DB().Preload("User").Preload("Disease").Preload("Blood").Preload("Gender").Preload("Title").Raw("SELECT * FROM patients").Find(&patient).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": patient})
}

// // DELETE /watch_videos/:id
// func DeletePatient(c *gin.Context) {
// 	id := c.Param("id")
// 	if tx := entity.DB().Exec("DELETE FROM patients WHERE id = ?", id); tx.RowsAffected == 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "patient not found"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": id})
// }

// // PATCH /watch_videos
// func UpdatePatient(c *gin.Context) {
// 	var patient entity.Patient
// 	if err := c.ShouldBindJSON(&patient); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if tx := entity.DB().Where("id = ?", patient.ID).First(&patient); tx.RowsAffected == 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "patient not found"})
// 		return
// 	}

// 	if err := entity.DB().Save(&patient).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": patient})
// }
