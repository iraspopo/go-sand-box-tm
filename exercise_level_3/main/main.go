package main

import "fmt"

func main() {
	fmt.Println("Exerice level 3")
	//init ;condition; post statement
	for i := 1; i < 1001; i++ {
		fmt.Println(i)
	}
	//init and post statement are optional (while loop version in go)
	fmt.Println()
	j := 1
	for j < 1001 {
		fmt.Println(j)
		j++
	}
	//infinite loop unless you break properly :)
	fmt.Println()
	counter := 1966
	for {
		if counter > 2020 {
			break
		}
		fmt.Println(counter)
		counter++
	}

	//print all upercase asci chars as rune - unicode code points
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}

	}

	//print out the mod (of 4) reminder of every number between 10 and 100
	for k := 10; k <= 100; k++ {
		fmt.Println(k % 4)
	}

	//switch with no switch statments
	switch {
	case false:
		fmt.Println("it is totaly false")
	case true:
		fmt.Println("it is true!")
	}

	favSport := "tennis"
	switch favSport {
	case "swimming":
		fmt.Printf("Favorite sport is %v \n", favSport)
	case "biking":
		fmt.Printf("Favorite sport is %v \n", favSport)
	case "tennis":
		fmt.Printf("Favorite sport is %v \n", favSport)
	default:
		fmt.Println("no match")
	}

	//SLICES
	//array size is build into a variable type
	var x [5]int
	var y [11]int
	fmt.Printf("%T\t%T\n", x, y)

	//composite literals - use them to initialize arrays, slices, maps..
	//slice z
	z := []int{11, 22, 33}
	fmt.Println("Going ove the slice with range")
	//range
	for i, v := range z {
		fmt.Println(i, v)
	}

	//just printing the slice or array fmt is doing it by adding comma separted values of the collection between square brackets
	fmt.Println(z)

	fmt.Println()

	a := []int{1, 2, 3}
	b := []int{4, 5, 6}

	//append and use od variadic parameters
	a = append(a, 100, 101, 102)
	a = append(a, b...)
	fmt.Println(a)

	//examples of deleting of slice elements by using append(x[1:4],y[8:]...)
	//append is from builtin package https://golang.org/pkg/builtin/#append
	// a slice looks like this now [1,2,3,100,101,102,4,5,6]
	// remove 100,101,102
	fmt.Println()
	a = append(a[:3], a[6:]...) //slice upper bound does not iclude the element of that index=> 3 upper bound is up to but not including element of index 3
	fmt.Println(a)

	//slice internaly has lenght and pointer to array. you can also use built function make to create slice of predetermine size
	//to optimize and make thing faster whe using append since  underlying array are always descrpyed and recreated if they are not
	//speified size? every time we add to slice that is going over original capcity it doubles the capacity automatically

	//make example
	fmt.Println()
	c := make([]int, 5, 10)  //slice type, length , capacity)
	c = append(c, 6, 7, 8)   //this will make length to be 8 and cap is still 10
	c = append(c, 9, 10, 11) //this will make length to be 11 and cap is doubled to 20

	fmt.Println(c, len(c), cap(c)) //buit in method on slices called len (for length) and cap (for capacity)

	//multidimansional slices (slice of slices)
	fmt.Println()
	d := []int{1, 2, 3}
	e := []int{11, 12, 13}
	f := []int{21, 22, 23}
	de := [][]int{d, e}
	de = append(de, f)

	fmt.Println(de, len(de), cap(de))

	//MAP
	fmt.Println("MAP")
	m := map[string]int{
		"Mila":     2018,
		"Igor":     1966,
		"Katarina": 1967,
		"Luka":     1997,
		"Matija":   2001}
	//use pange to iterate over, just like i slice or any other collection
	fmt.Println(m)
	for k, v := range m {
		fmt.Println(k, v)
	}

	//checking if key exist in the map ()
	fmt.Println(`Check if "Pera" key exist`)
	v, ok := m["Pera"]
	println(v, ok)

	fmt.Println(`Check if key "Igor" exist with if statment`)
	if v, ok := m["Igor"]; ok {
		println(v, ok)
	}

	//DELETE entry for the map
	fmt.Println(`Delete entry with key "Mila"`)
	delete(m, "Mila")
	fmt.Println(m)
	if v, ok := m["Mila"]; !ok {
		println("Mila expected value = 0, actual value = ", v, ". Mila delete operation result =", ok)
	}
}
