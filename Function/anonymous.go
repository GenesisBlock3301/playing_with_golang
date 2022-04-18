package main
import "fmt"

func main(){
	var area = func(l int, m int) int{
		return l*m
	}
	fmt.Println(area(5,6))
	fmt.Println("------------")
	func(l int, m int) {
		fmt.Println(l*m)
		return
	}(7,7)
	
}