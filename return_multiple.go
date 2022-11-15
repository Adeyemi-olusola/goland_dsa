package main 

import (
	"fmt"
)

func foo()(string ,string){
	return "two", "second"

}


func main(){
	fmt.Println(foo())
}