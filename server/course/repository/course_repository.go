package repository

import (
	"fmt"
	"github.com/Graphity/surge/server/course/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	//"gorm.io/gorm"
)

type CourseRepository interface {
	Save(course entity.Course)
	Update(course entity.Course)
	Delete(course entity.Course)
	FindById(id int) entity.Course
	FindAll() []entity.Course
	AutoMigration() error
}

func NewCourseRepository() CourseRepository {
	dsn := "host=localhost user=random password=secret dbname=surge port=5432 sslmode=disable TimeZone=Asia/Tbilisi"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	err = db.AutoMigrate(&entity.Course{})
	if err != nil {
		panic("Failed to migrate: Course")
	}
	return &database{
		conn: db,
	}
}

type database struct {
	conn *gorm.DB
}

func (db *database) AutoMigration() error {
	err := db.conn.AutoMigrate(&entity.Course{})
	if err != nil {
		return err
	}
	return nil
}

func (db *database) Save(course entity.Course) {
	db.conn.Create(&course)
}

func (db *database) Update(course entity.Course) {
	db.conn.Save(&course)
}

func (db *database) Delete(course entity.Course) {
	db.conn.Delete(&course)
}

func (db *database) FindById(id int) entity.Course {
	result := entity.Course{}
	data := db.conn.Find(&result, id)
	fmt.Println(data)
	return entity.Course{}
}

func (db *database) FindAll() []entity.Course {
	var courses []entity.Course
	db.conn.Set("gorm:auto_preload", true).Find(&courses)
	return courses
}
