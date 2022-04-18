package main
import "fmt"

func main(){
	var x[5]int
	var i int
	for i = 0; i < 5; i++{
		x[i] = i+5
	}

	for i = 0; i < len(x); i++{
		fmt.Printf("Element[%d] = %d\n",i,x[i])
	}
}