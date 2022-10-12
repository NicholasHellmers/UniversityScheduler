
package main

import (
	"fmt"
	"os"
	"encoding/json"
)
// todo: Generate Graph with weigths

// Constraints are that connect if there's a possiblity of a class path, 
// meaning: If the class is not the same class and the time is not in conflict, add an edge.
// Weights of the edges from source are the diff between start of day vs. start time of class 
// If it is not the first in the day it is the diff between the preffered wait time between classes.
// to solve: How to deal with classes that go on two different days.

type Recitation struct {
	Section int 				`json:"section"`
	Instructor string 			`json:"instructor"`
	Credits int 				`json:"credits"`
	Days []string 				`json:"days"`
	Time [][]int 				`json:"time"`
	Location string				`json:"location"`
}
type Lecture struct {
	Section int 				`json:"section"`
	Instructor string 			`json:"instructor"`
	Credits int					`json:"credits"`
	Days []string 				`json:"days"`
	Time [][]int 				`json:"time"`
	Location string 			`json:"location"`
	RecitationMandatory bool 	`json:"recitationMandatory"`
	REC []Recitation			`json:"recitations"`
}
type Course struct {
	Department string 			`json:"department"`
	CourseName string  			`json:"course_name"`
	CourseNumber int 			`json:"course_number"`
	LEC []Lecture 				`json:"lectures"`
}
type CourseList struct {
	Courses []Course 			`json:"courses"`
}

func LoadCourseList(filename string) (CourseList, error) {
	var courseList CourseList
	courseListFile, err := os.Open(filename)
	defer courseListFile.Close()
	if err != nil {
		return courseList, err
	}
	jsonParser := json.NewDecoder(courseListFile)
	jsonParser.Decode(&courseList)
	return courseList, nil
}

// Prints all the courses in the course list
func PrintCourses(courseList CourseList) {
	for i := 0; i < len(courseList.Courses); i++ {
		fmt.Println(courseList.Courses[i].Department, courseList.Courses[i].CourseName,courseList.Courses[i].CourseNumber)
	}
}

func main() {
	fmt.Println("Initializing...")
	courseList, err := LoadCourseList("data.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	PrintCourses(courseList)
}

