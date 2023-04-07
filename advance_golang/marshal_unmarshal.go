package main

import (
	"encoding/json"
	"fmt"
)

type Course struct {
	Name  string
	Grade float64
}

type Student struct {
	Name   string
	Age    int
	Course Course
}

func MarshalingSimpleObject(s1 *Student) string {
	bytes, err := json.MarshalIndent(s1, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func MarshalingComplexObject(student *[]Student) string {
	bytes, err := json.MarshalIndent(student, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func main() {
	var s Student
	var ss []Student

	s1 := Student{
		Name:   "Mr Nas",
		Age:    25,
		Course: Course{Name: "CSE32", Grade: 3.79},
	}
	simpleJson := MarshalingSimpleObject(&s1)
	student := []Student{
		{Name: "Nas", Age: 44, Course: Course{Name: "CSE32", Grade: 3.79}},
		{Name: "Mr. ALi", Age: 50, Course: Course{Name: "CSE32", Grade: 3.79}},
	}
	complexJson := MarshalingComplexObject(&student)

	err := json.Unmarshal([]byte(simpleJson), &s)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
	err = json.Unmarshal([]byte(complexJson), &ss)
	if err != nil {
		panic(err)
	}
	fmt.Println(ss)

}
