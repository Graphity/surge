package service

import (
	"github.com/Graphity/surge/server/Instructor/repository"
	"github.com/Graphity/surge/server/instructor/entity"
)

type InstructorService interface {
	Save(instructor entity.Instructor) error
	Update(instructor entity.Instructor) error
	Delete(instructor entity.Instructor) error
	FindAll() []entity.Instructor
}

type instructorService struct {
	repository repository.InstructorRepository
}

func (crs *instructorService) FindAll() []entity.Instructor {
	r := crs.repository.FindAll()
	return r
}

func New(instructorRepository repository.InstructorRepository) InstructorService {
	return &instructorService{
		repository: instructorRepository,
	}
}

func (crs *instructorService) Save(instructor entity.Instructor) error {
	crs.repository.Save(instructor)
	return nil
}

func (crs *instructorService) Update(instructor entity.Instructor) error {
	crs.repository.Update(instructor)
	return nil
}

func (crs *instructorService) Delete(instructor entity.Instructor) error {
	crs.repository.Delete(instructor)
	return nil
}