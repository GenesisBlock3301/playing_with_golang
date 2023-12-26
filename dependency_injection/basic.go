package main

import "fmt"

// DataService 1. Firstly declare a interface
type DataService interface {
	GetData() string
}

// APIService 2. Declare services
type APIService struct{}

// GetData 3. Implement the DataService interface method.
func (as APIService) GetData() string {
	return "Data from API"
}

type Processor struct {
	dataService DataService
}

func (p Processor) ProcessData() {
	data := p.dataService.GetData()
	fmt.Println("Processing:", data)
}

func main() {
	myService := APIService{}
	processor := Processor{myService}
	processor.ProcessData()
}
