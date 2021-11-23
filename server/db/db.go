package db

import (
	courseRepository "github.com/Graphity/surge/server/course/repository"
	courseService "github.com/Graphity/surge/server/course/service"
	instructorRepository "github.com/Graphity/surge/server/instructor/repository"
	instructorService "github.com/Graphity/surge/server/instructor/service"
)

var (
	CourseRepository = courseRepository.NewCourseRepository()
	CourseService    = courseService.New(CourseRepository)
	InstructorRepository = instructorRepository.NewInstructorRepository()
	InstructorService = instructorService.New(InstructorRepository)
)