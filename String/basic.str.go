package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main(){
	var x string = "Hello world"
	fmt.Println(x)
	fmt.Println(reflect.TypeOf(x))
	fmt.Println(len(x))
	// show ascii value of character
	fmt.Println((x[0]))
	fmt.Println(strings.ToUpper(x))
	// check prefix & sufix
	fmt.Println((strings.HasPrefix(x,"Hel")))
	fmt.Println(strings.HasSuffix(x,"ld"))

	// join
	var arr = []string{"a","b","c","d"}
	fmt.Println(strings.Join(arr,"="))

	// Repeat
	str := "Hi.... Sifat "
	fmt.Println(strings.Repeat(str,4))

	// Constrain
	fmt.Println(strings.Contains(str,"Hi"))

	// index
	fmt.Println(strings.Index(str,"Si"))

	// Count
	fmt.Println(strings.Count(str,"i"))

	// Replace
	fmt.Println(strings.Replace(str,"i","ZZ",1))
	split_str := "I,Love,You"
	// Split
	fmt.Println(strings.Split(split_str,","))
	splited_str := strings.Split(split_str,",")
	fmt.Println(reflect.TypeOf(splited_str))
	for i,v := range splited_str{
		fmt.Println("index: ",i,"value: ",v)
	}
	// Trim same as python strip
	trim := "Nur Amin Sifat                "
	fmt.Println(len(trim))
	fmt.Println(strings.TrimSpace(trim))
	fmt.Println(len(strings.TrimSpace(trim)))
}