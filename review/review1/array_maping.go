package main

import "fmt"

func main() {
	x := make(map[string]int)
	x["key"] = 0
	x["nas"] = 4
	delete(x, "nas")
	if name, ok := x["nas"]; ok {
		fmt.Println("Name is ", name, ok)
	}
	//elements := map[string]string{
	//	"name":        "nas",
	//	"age":         "25",
	//	"designation": "SWE",
	//}
	nested_element := map[string]map[string]string{
		"A": map[string]string{
			"name":  "Hybrid",
			"state": "gas",
		},
		"B": {
			"name":  "Hybrid",
			"state": "gas",
		},
	}
	fmt.Println(nested_element)
}
