package main
import "fmt"

// func name(argument_list) (return_types){}
func testFunc(fname string, lName string) (string){
	return fname+" "+lName;
}

func main(){
	var name string = testFunc("N A","Sifat")
	fmt.Println(name);
}