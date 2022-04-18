package main
import "fmt"

func main(){
	fmt.Printf("Enter your value: ")
	var k int
	fmt.Scanln(&k)
	switch k{
	case 10:
		fmt.Println("was<=10");fallthrough;
	case 20:
		fmt.Println("was <=20")
		fallthrough
	default:
		fmt.Println("Default case")
	}
}