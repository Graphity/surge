package main

import (
	"fmt"
	"github.com/Graphity/surge/server/db"
)

func main() {
	err := db.CourseRepository.AutoMigration()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	//err = db.CourseService.Save(entity.Course{
	//	Title:     "Course 5",
	//	Semester:  1,
	//	Credits:   6,
	//	Mandatory: false,
	//})
	if err != nil {
		fmt.Printf("Error %v\n", err)
	}

	courseData := db.CourseService.FindAll()
	for _, part := range courseData {
		fmt.Println(part.Title)
	}
}
