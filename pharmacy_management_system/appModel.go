// package pharmacy_management
package main

import (
	"math"
)

// Base -> Common type
type Base struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phoneNumber"`
}

func (b *Base) BaseInfo() *Base {
	return &Base{
		Id:          b.Id,
		Name:        b.Name,
		Address:     b.Address,
		PhoneNumber: b.PhoneNumber,
	}
}

// CommonMethodInterface is a common type container
type CommonMethodInterface interface {
	salaryDue() float64
	getInfo() *Base
}

type Employee struct {
	Common  Base    `json:"common"`
	Salary  float64 `json:"salary"`
	Payment float64 `json:"payment"`
}

func (e *Employee) getInfo() *Base {
	return &Base{
		Id:          e.Common.Id,
		Name:        e.Common.Name,
		Address:     e.Common.Address,
		PhoneNumber: e.Common.PhoneNumber,
	}
}

func (e *Employee) salaryDue() float64 {
	return math.Abs(e.Salary - e.Payment)
}

type Customer struct {
	Common        Base    `json:"common"`
	MedicinePrice float64 `json:"medicinePrice"`
	Payment       float64 `json:"payment"`
}

func (c *Customer) getInfo() *Base {
	return &Base{
		Id:          c.Common.Id,
		Name:        c.Common.Name,
		Address:     c.Common.Address,
		PhoneNumber: c.Common.PhoneNumber,
	}
}
func (c *Customer) salaryDue() float64 {
	return math.Abs(c.MedicinePrice - c.Payment)
}

func GetUserInfo(u CommonMethodInterface) *Base {
	return u.getInfo()
}

type SalesManagement struct {
	MedicineName         string  `json:"medicine_name"`
	Location             string  `json:"location"`
	NumberOfMedicine     int     `json:"number_of_medicine"`
	OriginalPrice        float64 `json:"original_price"`
	SellingPrice         float64 `json:"selling_price"`
	SoldNumberOfMedicine int     `json:"sold_number_of_medicine"`
	SoldAtaTime          int     `json:"sold_ata_time"`
}

func (s *SalesManagement) soldMedicineValue() float64 {
	return float64(s.SoldNumberOfMedicine) * s.SellingPrice
}

func (s *SalesManagement) expense() float64 {
	return s.OriginalPrice * float64(s.NumberOfMedicine)
}

func (s *SalesManagement) profit() float64 {
	if s.soldMedicineValue() > s.expense() {
		return s.soldMedicineValue() - s.expense()
	}
	return 0
}

func (s *SalesManagement) loss() float64 {
	if s.soldMedicineValue() < s.expense() {
		return s.expense() - s.soldMedicineValue()
	}
	return 0
}
