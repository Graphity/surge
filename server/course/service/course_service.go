package service

import (
	"github.com/Graphity/surge/server/course/entity"
	"github.com/Graphity/surge/server/course/repository"
)

type CourseService interface {
	Save(course entity.Course) error
	Update(course entity.Course) error
	Delete(course entity.Course) error
	FindAll() []entity.Course
}

type courseService struct {
	repository repository.CourseRepository
}

func (crs *courseService) FindAll() []entity.Course {
	r := crs.repository.FindAll()
	return r
}

func New(courseRepository repository.CourseRepository) CourseService {
	return &courseService{
		repository: courseRepository,
	}
}

func (crs *courseService) Save(course entity.Course) error {
	crs.repository.Save(course)
	return nil
}

func (crs *courseService) Update(course entity.Course) error {
	crs.repository.Update(course)
	return nil
}

func (crs *courseService) Delete(course entity.Course) error {
	crs.repository.Delete(course)
	return nil
}
