package repository

import (
	"fmt"
	"github.com/Graphity/surge/server/instructor/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type InstructorRepository interface {
	Save(instructor entity.Instructor)
	Update(instructor entity.Instructor)
	Delete(instructor entity.Instructor)
	FindById(id int) entity.Instructor
	FindAll() []entity.Instructor
}
type database struct {
	conn *gorm.DB
}

func (db *database) Save(instructor entity.Instructor) {
	db.conn.Create(&instructor)
}

func (db *database) Update(instructor entity.Instructor) {
	db.conn.Save(&instructor)
}

func (db *database) Delete(instructor entity.Instructor) {
	db.conn.Delete(&instructor)
}

func (db *database) FindById(id int) entity.Instructor {
	result := entity.Instructor{}
	data := db.conn.Find(&result, id)
	fmt.Println(data)
	return entity.Instructor{}
}

func (db *database) FindAll() []entity.Instructor {
	var courses []entity.Instructor
	db.conn.Set("gorm:auto_preload", true).Find(&courses)
	return courses
}

func NewInstructorRepository() InstructorRepository {
	dsn := "host=localhost user=random password=secret dbname=surge port=5432 sslmode=disable TimeZone=Asia/Tbilisi"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	err = db.AutoMigrate(&entity.Instructor{})
	if err != nil {
		panic("Failed to migrate: Course")
	}
	return &database{
		conn: db,
	}
}

func (db *database) AutoMigration() error {
	err := db.conn.AutoMigrate(&entity.Instructor{})
	if err != nil {
		return err
	}
	return nil
}
