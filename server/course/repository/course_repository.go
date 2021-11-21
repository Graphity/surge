package repository

import (
	"github.com/Graphity/surge/server/entity"
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
		connection: db,
	}
}

type database struct {
	connection *gorm.DB
}

func (d database) Save(course entity.Course) {
	panic("implement me")
}

func (d database) Update(course entity.Course) {
	panic("implement me")
}

func (d database) Delete(course entity.Course) {
	panic("implement me")
}

func (d database) FindById(id int) entity.Course {
	panic("implement me")
}

func (d database) FindByTitle(title string) entity.Course {
	panic("implement me")
}

func (d database) FindByCreditsQuantity(credits int) []entity.Course {
	panic("implement me")
}

func (d database) FindAll() []entity.Course {
	panic("implement me")
}

func (d database) FindAllMandatory() []entity.Course {
	panic("implement me")
}

func (d database) FindAllNonMandatory() []entity.Course {
	panic("implement me")
}

func (d database) FindBySemester(semester int) []entity.Course {
	panic("implement me")
}

func (d database) CloseDB() {
	panic("implement me")
}
