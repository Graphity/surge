package repository

import (
	"github.com/Graphity/surge/server/instructor/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type InstructorRepository interface {
	Save(instructor entity.Instructor)
	Update(instructor entity.Instructor)
	Delete(instructor entity.Instructor)
	FindById(id int) entity.Instructor
	FindByFirstName(firstName string) []entity.Instructor
	FindByLastName(lastName string) []entity.Instructor
	FindByPhoneNumber(phoneNumber string) entity.Instructor
	FindByEmail(email string) entity.Instructor
	FindByLinkedIn(linkedIn string) entity.Instructor
	FindByRating(rating int) []entity.Instructor
	CloseDB()
}

func NewInstructorRepository() InstructorRepository {
	dsn := "host=localhost user=random password=secret dbname=surge port=5432 sslmode=disable TimeZone=Asia/Tbilisi"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	err = db.AutoMigrate(&entity.Instructor{})
	if err != nil {
		panic("Failed to migrate: Instructor")
	}
	return &databaseInstructor{
		connection: db,
	}
}

type databaseInstructor struct {
	connection *gorm.DB
}

func (db databaseInstructor) Save(instructor entity.Instructor) {

	panic("implement me")
}

func (db databaseInstructor) Update(instructor entity.Instructor) {
	panic("implement me")
}

func (db databaseInstructor) Delete(instructor entity.Instructor) {
	panic("implement me")
}

func (db databaseInstructor) FindById(id int) entity.Instructor {
	panic("implement me")
}

func (db databaseInstructor) FindByFirstName(firstName string) []entity.Instructor {
	panic("implement me")
}

func (db databaseInstructor) FindByLastName(lastName string) []entity.Instructor {
	panic("implement me")
}

func (db databaseInstructor) FindByPhoneNumber(phoneNumber string) entity.Instructor {
	panic("implement me")
}

func (db databaseInstructor) FindByEmail(email string) entity.Instructor {
	panic("implement me")
}

func (db databaseInstructor) FindByLinkedIn(linkedIn string) entity.Instructor {
	panic("implement me")
}

func (db databaseInstructor) FindByRating(rating int) []entity.Instructor {
	panic("implement me")
}

func (db databaseInstructor) CloseDB() {
	panic("implement me")
}
