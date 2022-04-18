/*
for initialization expression:
evaluation_expression;
iteration_expression{
	one or more statement here.
}
*/
package main
import "fmt"

func main(){
	var i int
	for i = 1; i < 11; i++{
		if(i%2 == 0){
			fmt.Println("Even no:", i)
		}else{
			fmt.Println("Odd no:",i)
		}
	}
}