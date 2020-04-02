package main

import "fmt"

var x int = 66
var y string = "Kicking Horse"
var z bool = true

var i int
var j string
var k bool

//T1 ... there are no cast in go only coversion - so let see we can cast this custom type to int
type T1 int

var t1 T1

func main() {
	// a := 42
	// b := "James Bond"
	// c := true
	a, b, c := 42, "James Brown", true
	fmt.Println(a, b, c)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Print("\n") //new line with fmt.print

	fmt.Println(x, y, z)
	fmt.Println() //new line

	fmt.Println("zero type value for int is", i)
	fmt.Println("zero type value for string is", `"`, j, `"`, "lenght: ", len(j))
	fmt.Println("zero type value for bool is", k)
	fmt.Println() //new line
	//%d integer, %s string %t for bool
	s := fmt.Sprintf("%d %s %t", x, y, z)
	fmt.Println(s)
	fmt.Println()
	//%v print the value of the variable no matter what the type integer, and \t tab for formating
	s = fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)
	fmt.Println()
	//%T to print the type and \n to and new line
	fmt.Println(t1)
	fmt.Printf("%T\n", t1)
	t1 = 42
	fmt.Println(t1)

	//convert t1 to underlying type (int)
	t2 := int(t1)
	fmt.Printf("type=%T\tvalue=%v\n", t2, t2)

	//priont number in decimal, binary and hex
	n := 66
	fmt.Printf("dec=%d\tbin=%b\thex=%#x\n", n, n, n)

}
