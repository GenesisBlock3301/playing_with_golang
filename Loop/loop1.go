package main
import "fmt"

func main(){
	// same as while loop
	sum:=1
	for sum <= 100{
		sum += sum
		fmt.Println("Sum is : ",sum)
	}
	
}