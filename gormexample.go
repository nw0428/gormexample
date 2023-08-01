package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type Car struct {
	gorm.Model
	Make           string
	Manufacturer   string
	Year           int8
	NumberOfWheels int8
	Price          decimal.Decimal
}

type Student struct {
	gorm.Model
	Name      string
	Birthdate time.Time
}

func main() {
	dsn := "host=localhost user=davidzabner dbname=gormexample port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		fmt.Printf("%s\n", err)
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Car{})
	db.AutoMigrate(&Student{})

	david := &Student{Name: "David Zabner", Birthdate: time.Date(1991, time.June, 28, 0, 0, 0, 0, time.Local)}
	db.Create(&david)
	db.Create(&Student{Name: "George Smith", Birthdate: time.Date(2005, time.July, 4, 0, 0, 0, 0, time.Local)})

	// Read
	var student Student
	db.First(&student, "name like ?", "David%")
	fmt.Printf("Student with a name like David is %+v\n", student)

	var student2 Student
	db.First(&student2)
	fmt.Printf("First student is %+v\n", student2)

	var student3 Student
	db.Model(&student3).Update("Name", "Sally M")

	// How does changing @student to student2 or 3 change
	db.Model(&student).Updates(Student{Name: "Ellie", Birthdate: time.Now()}) // non-zero fields

	// Delete - delete product
	db.Delete(&student, student.ID)

}
