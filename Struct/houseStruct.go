package main

import "fmt"

type house struct {
	noRooms int16
	price   float64
	city    string
}
type housList struct {
	houses []house
}

func (a *housList) print() []house {
	response := []house{}
	for _, val := range a.houses {
		house := house{noRooms: val.noRooms, price: val.price, city: val.city}
		response = append(response, house)
	}
	return response
}

func main() {
	houses := housList{
		houses:[]house{
			{
				noRooms: 1,
				price:   44.6,
				city:    "Jamalpur",
			},
			{
				noRooms: 2,
				price:   44.6,
				city:    "Jamalpur",
			},
		},
	}
	h := houses.print()
	fmt.Println(h)

	// for _,val:=range houses.houses{
	// 	fmt.Println(val.city)
	// }
}
