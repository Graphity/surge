package db

import (
	"github.com/Graphity/surge/server/course/repository"
	"github.com/Graphity/surge/server/course/service"
)

var (
	CourseRepository = repository.NewCourseRepository()
	CourseService    = service.New(CourseRepository)
)