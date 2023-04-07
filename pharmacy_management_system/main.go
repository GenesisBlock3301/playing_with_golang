package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func printMenu() {
	fmt.Println(`1. Add Employee.
                     2. Show Employee.`)
}
func OpenFile() []int {
	jsonFile, _ := os.Open("test.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	fmt.Println(byteValue)
	defer jsonFile.Close()
	var updatedVal []Employee
	json.Unmarshal(byteValue, &updatedVal)
	return updatedVal
}
func main() {
	printMenu()
	OpenFile()
	var EmployeeObjectList []Employee
	//var CustomerObjectList []Customer
	//var SalesObjectData []SalesManagement
	for true {
		var val int
		menuChoice, _ := fmt.Scanln(&val)
		if menuChoice == 1 {
			var ID int
			var name, address, phone string
			var salary, payment float64
			fmt.Println("Enter Employee ID: ")
			fmt.Scanln(&ID)
			fmt.Println("Enter Employee Name : ")
			fmt.Scanln(&name)
			fmt.Println("Enter Employee address : ")
			fmt.Scanln(&address)
			fmt.Println("Enter Employee phone : ")
			fmt.Scanln(&phone)
			fmt.Println("Enter Employee salary : ")
			fmt.Scanln(&salary)
			fmt.Println("Enter Employee payment : ")
			fmt.Scanln(&payment)
			employee := Employee{
				Common: Base{
					Id:          ID,
					Name:        name,
					Address:     address,
					PhoneNumber: phone,
				},
				Salary:  salary,
				Payment: payment,
			}
			oldList + arr
			EmployeeObjectList = append(EmployeeObjectList, employee)
			fmt.Println(EmployeeObjectList)

			//file, _ := json.MarshalIndent(EmployeeObjectList, "", " ")
			//
			//_ = ioutil.WriteFile("test.json", file, 0644)
		} else {
			fmt.Println("Terminate successfully")
			break
		}

	}
}
