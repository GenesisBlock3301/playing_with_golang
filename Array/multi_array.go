package main
import "fmt"

func main(){
	a := [2][3]int{{1,2,3},{4,5,6}}
	fmt.Println(a)
	for i:= 0; i < len(a);i++{
		for j:=0; j < len(a[0]); j++{
			fmt.Printf("%d ",a[i][j])
		}
		fmt.Println()
	}
	fmt.Println(a[0][:2])
}