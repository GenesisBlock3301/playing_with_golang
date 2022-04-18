package main

import (
	"fmt"
)
type Vertex struct{
	Latitude,Longitude float64
}
func main(){

	var m = map[string]Vertex{
		"javapoint":Vertex{40.67889789,90.67568679678},
		"SST":Vertex{95.0594054,-78.053458493},
	}
	fmt.Println(m["javapoint"].Latitude)
}