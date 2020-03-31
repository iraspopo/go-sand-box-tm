package main

import (
	"fmt"

	"github.com/iraspopo/go-sand-box-tm/using_go_modules/hello"
)

var varWithScopeOfTheWholePackag = 100

//by default all delcalred vriables of int type without explicit initializatio will be set zero type value which is 0 for int
var varOfTypeInt int

//Hello func return Hello world!
func main() {
	//declare and assign with short declaration operator :=
	n, err := fmt.Println(hello.Hello())
	fmt.Println(n, err)
	//difference between := and var declaration => you can't use := outside of function at the pck level
	fmt.Println(varWithScopeOfTheWholePackag, varOfTypeInt)
	foo()
	varOfTypeInt = 100
	fmt.Println("varOfTypeInt = ", varOfTypeInt)
}

func foo() {
	fmt.Println(varWithScopeOfTheWholePackag)
}
