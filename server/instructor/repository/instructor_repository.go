package repository

import "github.com/Graphity/surge/server/entity"

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
}
