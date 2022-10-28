package entity

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"gorm.io/driver/sqlite"
)

var db *gorm.DB

func DB() *gorm.DB {

	return db

}

func SetupDatabase() {

	database, err := gorm.Open(sqlite.Open("SA-65.db"), &gorm.Config{})

	if err != nil {

		panic("failed to connect database")

	}

	// Migrate the schema

	database.AutoMigrate(
		&User{},
		&Time{},
		&Gender{},
		&Role{},
		&Program{},
		&Room_type{},
		&Room_price{},
		&Furniture{},
		&Set_of_furniture{},
		&Semester{},
		&Job_Position{},
		&Province{},
		&Equipment{},
		&Borrowcard{},
		&Employee{},
		&Student{},
		&Room{},
		&Booking{},
		&ListData{},
		&Payment_Bill{},
		&Repair{},
	)

	db = database

	// password, err := bcrypt.GenerateFromPassword([]byte("123456"), 14)

	// db.Model(&Student{}).Create(&Student{
	// 	STUDENT_NUMBER: "B611308",
	// 	Password:       string(password),
	// })
	// db.Model(&Student{}).Create(&Student{
	// 	STUDENT_NUMBER: "B6428853",
	// 	STUDENT_NAME:   "SriSuk Malee",
	// 	PERSONAL_ID:    1234569875159,
	// 	Password:       string(password),
	// })

	// 	var chanwit User
	// 	var name User
	// 	db.Raw("SELECT * FROM users WHERE email = ?", "tanapon@gmail.com").Scan(&chanwit)
	// 	db.Raw("SELECT * FROM users WHERE email = ?", "name@example.com").Scan(&name)

	// 	var phanukorn Student
	// 	db.Raw("SELECT * FROM students WHERE STUDENT_NUMBER = ?", "B6311308").Scan(&phanukorn)

	T1 := Time{
		Time_number: "1 เทอม",
	}
	db.Model(&Time{}).Create(&T1)
	T2 := Time{
		Time_number: "2 เทอม",
	}
	db.Model(&Time{}).Create(&T2)
	T3 := Time{
		Time_number: "1 ปี",
	}
	db.Model(&Time{}).Create(&T3)
	T4 := Time{
		Time_number: "1 ปี 1 เทอม",
	}
	db.Model(&Time{}).Create(&T4)

	// 	Room1 := Room{}
	// 	db.Model(&Room{}).Create(&Room1)
	// 	Room2 := Room{}
	// 	db.Model(&Room{}).Create(&Room2)
	// 	Room3 := Room{}
	// 	db.Model(&Room{}).Create(&Room3)

	// 	var target User
	// 	db.Model(&User{}).Find(&target, db.Where("email = ?", "tanapon@gmail.com"))

	type1 := Room_type{
		Room_type_name: "Single",
	}
	db.Model(&Room_type{}).Create(&type1)

	type2 := Room_type{
		Room_type_name: "Twin",
	}

	db.Model(&Room_type{}).Create(&type2)

	//Resolution Data
	price1 := Room_price{
		Price: "3000",
	}
	db.Model(&Room_price{}).Create(&price1)

	price2 := Room_price{
		Price: "5000",
	}
	db.Model(&Room_price{}).Create(&price2)

	//set_of
	set1 := Set_of_furniture{
		Set_of_furniture_title: "Set1",
	}
	db.Model(&Set_of_furniture{}).Create(&set1)
	set2 := Set_of_furniture{
		Set_of_furniture_title: "Set2",
	}
	db.Model(&Set_of_furniture{}).Create(&set2)

	furniture1 := Furniture{
		Furniture_type: "Table1",

		Set_of_furniture: set1,
	}
	db.Model(&Furniture{}).Create(&furniture1)

	furniture2 := Furniture{
		Furniture_type:   "Table2",
		Set_of_furniture: set2,
	}
	db.Model(&Furniture{}).Create(&furniture2)

	db.Model(&Furniture{}).Create(&Furniture{
		Furniture_type:   "table3",
		Set_of_furniture: set1,
	})

	// === Query
	//
	r1 := Room{
		Room_type:        type1,
		Room_price:       price1,
		Set_of_furniture: set1,
	}
	db.Model(&Room{}).Create(&r1)

	db.Model(&Room{}).Create(&Room{
		Room_type:        type1,
		Room_price:       price1,
		Set_of_furniture: set1,
	})

	db.Model(&Room{}).Create(&Room{
		Room_type:        type2,
		Room_price:       price1,
		Set_of_furniture: set1,
	})
	db.Model(&Booking{}).Create(&Booking{
		Room: r1,
		Time: T1,
	})
	//add example

	// ======================================================================================================================
	// ======================================  Employee  =====================================================================
	// ======================================================================================================================

	//Gender
	gender1 := Gender{
		Name: "Male",
	}

	db.Model(&Gender{}).Create(&gender1)

	gender2 := Gender{
		Name: "FeMale",
	}

	db.Model(&Gender{}).Create(&gender2)

	//insert job_position
	job_position1 := Job_Position{
		Name: "Admin",
	}
	db.Model(&Job_Position{}).Create(&job_position1)

	job_position2 := Job_Position{
		Name: "Housekeeper",
	}
	db.Model(&Job_Position{}).Create(&job_position2)

	job_position3 := Job_Position{
		Name: "Security Guard",
	}
	db.Model(&Job_Position{}).Create(&job_position3)

	job_position4 := Job_Position{
		Name: "Mechanic",
	}
	db.Model(&Job_Position{}).Create(&job_position4)

	//province
	roiet := Province{
		Name: "Roiet",
	}
	db.Model(&Province{}).Create(&roiet)
	korat := Province{
		Name: "Nakhon Ratchasima",
	}
	db.Model(&Province{}).Create(&korat)

	chon := Province{
		Name: "Chonburi",
	}
	db.Model(&Province{}).Create(&chon)

	bangkok := Province{
		Name: "Bangkok",
	}
	db.Model(&Province{}).Create(&bangkok)

	password1, err := bcrypt.GenerateFromPassword([]byte("abc12456"), 14)
	password2, err := bcrypt.GenerateFromPassword([]byte("123456"), 14)
	password3, err := bcrypt.GenerateFromPassword([]byte("1111111111111"), 14)
	password4, err := bcrypt.GenerateFromPassword([]byte("adas8485"), 14)

	//insert employee'
	em1 := Employee{
		Personal_ID: "1456287463254",
		Email:       "ana@gmail.com",
		Name:        "Ana poul",
		Password:    string(password1),

		Gender:       gender2,
		Job_Position: job_position1,
		Province:     korat,
	}
	db.Model(&Employee{}).Create(&em1)

	em2 := Employee{
		Personal_ID: "1456287463255",
		Email:       "ana345@gmail.com",
		Name:        "Ana poul",
		Password:    string(password1),

		Gender:       gender2,
		Job_Position: job_position1,
		Province:     korat,
	}
	db.Model(&Employee{}).Create(&em2)

	em3 := Employee{
		Personal_ID: "5874621453054",
		Email:       "kerkkiat@gmail.com",
		Name:        "Kerkkiat Prabmontree",
		Password:    string(password3),

		Gender:       gender1,
		Job_Position: job_position3,
		Province:     bangkok,
	}
	db.Model(&Employee{}).Create(&em3)

	em4 := Employee{
		Personal_ID: "4587652145385",
		Email:       "matinez@gmail.com",
		Name:        "Devid Matinez",
		Password:    string(password2),

		Gender:       gender1,
		Job_Position: job_position4,
		Province:     chon,
	}
	db.Model(&Employee{}).Create(&em4)

	em5 := Employee{
		Personal_ID: "5847532016420",
		Email:       "akira@gmail.com",
		Name:        "Akira komisu",
		Password:    string(password4),

		Gender:       gender1,
		Job_Position: job_position1,
		Province:     roiet,
	}
	db.Model(&Employee{}).Create(&em5)

	// ======================================================================================================================
	// ======================================  Student  =====================================================================
	// ======================================================================================================================

	// --- Program Data
	p1 := Program{
		Program_name: "Computer engineering",
	}
	db.Model(&Program{}).Create(&p1)
	p2 := Program{
		Program_name: "Telecommunication engineering",
	}
	db.Model(&Program{}).Create(&p2)
	p3 := Program{
		Program_name: "Program in Biology",
	}
	db.Model(&Program{}).Create(&p3)
	p4 := Program{
		Program_name: "Institute of Nursing",
	}
	db.Model(&Program{}).Create(&p4)

	// --- Role Data

	role1 := Role{
		Role_name: "Student",
	}
	db.Model(&Role{}).Create(&role1)

	db.Model(&Student{}).Create(&Student{
		STUDENT_NUMBER: "B62457815",
		STUDENT_NAME:   "Supachai srikawe",
		PERSONAL_ID:    "1786542390457",
		Password:       string(password4),

		Gender:   gender1,
		Program:  p3,
		Province: roiet,
		Role:     role1,
		Employee: em1,
	})
	db.Model(&Student{}).Create(&Student{
		STUDENT_NUMBER: "B6311308",
		STUDENT_NAME:   "Phanuukorn Kongpet",
		PERSONAL_ID:    "1786542390457",
		Password:       string(password2),

		Gender:   gender1,
		Program:  p3,
		Province: roiet,
		Role:     role1,
		Employee: em1,
	})
	db.Model(&Furniture{}).Create(&Furniture{
		Furniture_type: "Fan",
	})

	province1 := Province{
		Name: "Kamphaeng Phet",
	}
	db.Model(&Province{}).Create(&province1)
	province2 := Province{
		Name: "Chiang Rai",
	}
	db.Model(&Province{}).Create(&province2)
	province3 := Province{
		Name: "Angthong",
	}
	db.Model(&Province{}).Create(&province3)

	// Semester Data ------------------------------------------------------
	semester1 := Semester{
		Semester: "1/2564",
	}
	db.Model(&Semester{}).Create(&semester1)

	semester2 := Semester{
		Semester: "2/2564",
	}
	db.Model(&Semester{}).Create(&semester2)

	semester3 := Semester{
		Semester: "3/2564",
	}
	db.Model(&Semester{}).Create(&semester3)

	semester4 := Semester{
		Semester: "1/2565",
	}
	db.Model(&Semester{}).Create(&semester4)
}
