package main
import (
	"fmt"
)

func main(){

	var val int
	fmt.Printf("Enter the number: ")
	fmt.Scanln(&val)
	switch val {
	case 10:
		fmt.Println("Val is 10")
	default:
		fmt.Println("Val is not 10")
	}
}