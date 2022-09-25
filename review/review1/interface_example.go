package main

import "fmt"

type taxSystem interface {
	calculateTax() int
}

type indianTax struct {
	taxPercentage int
	income        int
}

func (i *indianTax) calculateTax() int {
	tax := i.income * i.taxPercentage / 100
	return tax
}

type singapurTax struct {
	taxPercentage int
	income        int
}

func (s *singapurTax) calculateTax() int {
	tax := s.income * s.taxPercentage / 100
	return tax
}

func main() {
	indianTax := &indianTax{taxPercentage: 30, income: 1000}

	singapurTax := &singapurTax{taxPercentage: 10, income: 2000}

	taxSystems := []taxSystem{indianTax, singapurTax}
	totalTax := calculateTotalTax(taxSystems)

	fmt.Println(totalTax)
}

func calculateTotalTax(taxSystems []taxSystem) int {
	totalTax := 0
	for _, t := range taxSystems {
		totalTax += t.calculateTax()
	}
	return totalTax
}
