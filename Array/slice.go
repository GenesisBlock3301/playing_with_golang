package main
import "fmt"

func printSlice(s[] int){
	fmt.Println("length =",len(s),"Capasity =",cap(s),s)
}
func main(){
	slice := []int{1,2,3,4,5,6}
	printSlice(slice)
	slice2 := slice[:0]
	printSlice(slice2)
	delete := append(slice[:2],slice[2+1:]...)
	fmt.Println("Deleted array: ", delete)
}