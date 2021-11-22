package repository

import (
	"github.com/Graphity/surge/server/course/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type CourseRepository interface {
	Save(course entity.Course)
	Update(course entity.Course)
	Delete(course entity.Course)
	FindById(id int) entity.Course
	FindByTitle(title string) entity.Course
	FindByCreditsQuantity(credits int) []entity.Course
	FindAll() []entity.Course
	FindAllMandatory() []entity.Course
	FindAllNonMandatory() []entity.Course
	FindBySemester(semester int) []entity.Course
	CloseDB()
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
	panic("implement me")
}

func (db *database) FindByTitle(title string) entity.Course {
	panic("implement me")
}

func (db *database) FindByCreditsQuantity(credits int) []entity.Course {
	panic("implement me")
}

func (db *database) FindAll() []entity.Course {
	panic("implement me")
}

func (db *database) FindAllMandatory() []entity.Course {
	panic("implement me")
}

func (db *database) FindAllNonMandatory() []entity.Course {
	panic("implement me")
}

func (db *database) FindBySemester(semester int) []entity.Course {
	panic("implement me")
}

func (db *database) CloseDB() {

}
