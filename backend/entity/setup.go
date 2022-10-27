package entity

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("sa-65.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	database.AutoMigrate(
		&User{},
		&Title{},
		&Gender{},
		&Blood{},
		&Disease{},
		&Patient{},
	)

	db = database

	password, _ := bcrypt.GenerateFromPassword([]byte("123456"), 14)

	db.Model(&User{}).Create(&User{
		Name:     "Wallaya Patisang",
		Email:    "wallaya.1999@gmail.com",
		Password: string(password),
		Role:     "user",
	})
	db.Model(&User{}).Create(&User{
		Name:     "Name",
		Email:    "name@gmail.com",
		Password: string(password),
		Role:     "employee",
	})

	var wallaya User
	var name User
	db.Raw("SELECT * FROM users WHERE email = ?", "wallaya@gmail.com").Scan(&wallaya)
	db.Raw("SELECT * FROM users WHERE email = ?", "name@gmail.com").Scan(&name)

	// --- Title Data
	mr := Title{
		Name: "นาย",
	}
	db.Model(&Title{}).Create(&mr)

	miss := Title{
		Name: "นางสาว",
	}
	db.Model(&Title{}).Create(&miss)

	mrs := Title{
		Name: "นาง",
	}
	db.Model(&Title{}).Create(&mrs)

	// GENDER Data
	male := Gender{
		Name: "ชาย",
	}
	db.Model(&Gender{}).Create(&male)

	female := Gender{
		Name: "หญิง",
	}
	db.Model(&Gender{}).Create(&female)

	// Blood Data
	a := Blood{
		Name: "A",
	}
	db.Model(&Blood{}).Create(&a)

	b := Blood{
		Name: "B",
	}
	db.Model(&Blood{}).Create(&b)

	ab := Blood{
		Name: "AB",
	}
	db.Model(&Blood{}).Create(&ab)

	o := Blood{
		Name: "O",
	}
	db.Model(&Blood{}).Create(&o)

	// disease Data
	diabetes := Disease{
		Name: "เบาหวาน",
	}
	db.Model(&Disease{}).Create(&diabetes)

	hypertension := Disease{
		Name: "ความดันโลหิตสูง",
	}
	db.Model(&Disease{}).Create(&hypertension)

	tuberculosis := Disease{
		Name: " วัณโรค",
	}
	db.Model(&Disease{}).Create(&tuberculosis)

	none := Disease{
		Name: " ไม่มี",
	}
	db.Model(&Disease{}).Create(&none)

	// patient 1
	// db.Model(&Patient{}).Create(&Patient{

	// 	Title:   miss,
	// 	Gender:  female,
	// 	Blood:   a,
	// 	Disease: none,
	// })

	//
	// === Query
	//

	// 	var target User
	// 	db.Model(&User{}).Find(&target, db.Where("email = ?", "chanwit@gmail.com"))

	// 	var watchedPlaylist Playlist
	// 	db.Model(&Playlist{}).Find(&watchedPlaylist, db.Where("title = ? and owner_id = ?", "Watched", target.ID))

	// 	var watchedList []*WatchVideo
	// 	db.Model(&WatchVideo{}).
	// 		Joins("Playlist").
	// 		Joins("Resolution").
	// 		Joins("Video").
	// 		Find(&watchedList, db.Where("playlist_id = ?", watchedPlaylist.ID))

	// 	for _, wl := range watchedList {
	// 		fmt.Printf("Watch Video: %v\n", wl.ID)
	// 		fmt.Printf("%v\n", wl.Playlist.Title)
	// 		fmt.Printf("%v\n", wl.Resolution.Value)
	// 		fmt.Printf("%v\n", wl.Video.Name)
	// 		fmt.Println("====")
	// 	}
	// }

	/*import (
	  	"gorm.io/gorm"

	  	"gorm.io/driver/sqlite"
	  )

	  var db *gorm.DB

	  func DB() *gorm.DB {

	  	return db

	  }

	  func SetupDatabase() {

	  	database, err := gorm.Open(sqlite.Open("sa-65.db"), &gorm.Config{})

	  	if err != nil {

	  		panic("failed to connect database")

	  	}

	  	// Migrate the schema

	  	database.AutoMigrate(&User{})

	  	db = database

	  }*/
}
