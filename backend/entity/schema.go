package entity

import (
	"time"

	// "github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"uniqueIndex" `
	Password string	`json:"-"`
	Role     string

	// 1 user เป็นเจ้าของได้หลาย playlist
	Patients []Patient `gorm:"foreignKey:UserID"`
}

type Title struct {
	gorm.Model
	Name string

	Patients []Patient `gorm:"foreignKey:TitleID"`
}

type Gender struct {
	gorm.Model
	Name string

	Patients []Patient `gorm:"foreignKey:GenderID"`
}

type Blood struct {
	gorm.Model
	Name string

	Patients []Patient `gorm:"foreignKey:BloodID"`
}

type Disease struct {
	gorm.Model
	Name string

	Patients []Patient `gorm:"foreignKey:DiseaseID"`
}

type Patient struct {
	gorm.Model
	PersonalID   string `gorm:"uniqueIndex" `
	Allergy      string
	Tel          string
	BirthdayTime time.Time

	// DiseaseID ทำหน้าที่เป็น FK
	DiseaseID *uint
	Disease   Disease `gorm:"references:id" valid:"-"`

	// BloodID ทำหน้าที่เป็น FK
	BloodID *uint
	Blood   Blood `gorm:"references:id" valid:"-"`

	// GenderID ทำหน้าที่เป็น FK
	GenderID *uint
	Gender   Gender `gorm:"references:id" valid:"-"`

	// TitleID ทำหน้าที่เป็น FK
	TitleID *uint
	Title   Title `gorm:"references:id" valid:"-"`

	// UserID ทำหน้าที่เป็น FK
	UserID *uint
	User   User `gorm:"references:id" valid:"-"`
}

// func init() {
// 	govalidator.CustomTypeTagMap.Set("past", func(i interface{}, context interface{}) bool {
// 		t := i.(time.Time)
// 		now := time.Now()
// 		return now.After(t)
// 	})
// 	govalidator.CustomTypeTagMap.Set("future", func(i interface{}, context interface{}) bool {
// 		t := i.(time.Time)
// 		now := time.Now()
// 		return now.Before(time.Time(t))
// 	})
// }
