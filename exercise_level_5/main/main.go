package main

import "fmt"

func main() {
	fmt.Println("Examples : slices and map")
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

	fmt.Println("Exerice level 4 ==========   starts HERE  ===================")
	fmt.Println()
	fmt.Println("Exerice level 4 #1")
	arr := [5]int{1, 2, 3, 4, 5}
	for i, v := range arr {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", arr)
	fmt.Println()

	fmt.Println("Exerice level 4 #2")
	sl := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	for i, v := range sl {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", sl)
	fmt.Println()

	fmt.Println("Exerice level 4 #3")
	sl1 := sl[:5]
	fmt.Println(sl1)
	sl2 := sl[5:]
	fmt.Println(sl2)
	sl3 := sl[2:7]
	fmt.Println(sl3)
	sl4 := sl[1:6]
	fmt.Println(sl4)
	fmt.Println(sl)

	fmt.Println()

	fmt.Println("Exerice level 4 #4")
	fmt.Println("sl", sl)
	sl = append(sl, 52)
	fmt.Println("appended 52 to sl", sl)
	sl = append(sl, 53, 54, 55)
	fmt.Println("appended 53 54 55 to sl", sl)

	sl2append := []int{56, 57, 58, 59, 60}
	sl = append(sl, sl2append...)
	fmt.Println("sl2append", sl2append)
	fmt.Println("appended sl2append to sl", sl)
	fmt.Println()

	fmt.Println("Exerice level 4 #5")
	//deleting some part of the slice with append
	slX := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	slDel := append(slX[:3], slX[6:]...)
	fmt.Println("slDel", slDel)
	fmt.Println("slX", slX)
	fmt.Println()

	fmt.Println("Exerice level 4 #6")
	//deleting some part of the slice with append
	slUsa := make([]string, 50, 50)
	slUsa = []string{"Alabama",
		"Alaska",
		"Arizona",
		"Arkansas",
		"California",
		"Colorado",
		"Connecticut",
		"Delaware",
		"Florida",
		"Georgia",
		"Hawaii",
		"Idaho",
		"Illinois",
		"Indiana",
		"Iowa",
		"Kansas",
		"Kentucky",
		"Louisiana",
		"Maine",
		"Maryland",
		"Massachusetts",
		"Michigan",
		"Minnesota",
		"Mississippi",
		"Missouri",
		"Montana",
		"Nebraska",
		"Nevada",
		"New Hampshire",
		"New Jersey",
		"New Mexico",
		"New York",
		"North Carolina",
		"North Dakota",
		"Ohio",
		"Oklahoma",
		"Oregon",
		"Pennsylvania",
		"Rhode Island",
		"South Carolina",
		"South Dakota",
		"Tennessee",
		"Texas",
		"Utah",
		"Vermont",
		"Virginia",
		"Washington",
		"West Virginia",
		"Wisconsin",
		"Wyoming"}
	fmt.Println(slUsa, "len ", len(slUsa), "cap ", cap(slUsa))
	for i := 0; i < len(slUsa); i++ {
		fmt.Println(slUsa[i], i)
	}
	fmt.Println()

	fmt.Println("Exerice level 4 #7")

	slOfSlices := [][]string{{"James", "Bond", "Shaken, not stirred"}, {"Miss", "Moneypenny", "Hellooooo, James"}}
	fmt.Println(slOfSlices)
	for i, v := range slOfSlices {
		fmt.Println("Record: ", i, " Value: ", v)
		for j, v1 := range v {
			fmt.Println("\t", j, "\t", v1)
		}
	}
	fmt.Println()

	fmt.Println("Exerice level 4 #8")

	mapEx8 := map[string][]string{
		"Mila":     []string{"ball", "walk", "food"},
		"Igor":     []string{"sailing", "biking", "Mila"},
		"Katarina": []string{"books", "cooking", "Mila"}}
	for k, v := range mapEx8 {
		fmt.Println(k, v)
		for i, val := range v {
			fmt.Println("\t", i, "\t", val)
		}

	}
	fmt.Println()

	fmt.Println("Exerice level 4 #9 - add `Cucki`")

	mapEx8["Cucki"] = []string{"guitar", "drawing", "Mila"}
	for k, v := range mapEx8 {
		fmt.Println(k, v)
		for i, val := range v {
			fmt.Println("\t", i, "\t", val)
		}

	}
	fmt.Println()

	fmt.Println("Exerice level 4 #10 - delete `Cucki`")

	delete(mapEx8, "Cucki")
	for k, v := range mapEx8 {
		fmt.Println(k, v)
		for i, val := range v {
			fmt.Println("\t", i, "\t", val)
		}

	}
	fmt.Println()

	fmt.Println("Exerice level 5 ==========   starts HERE  ===================")
	fmt.Println()
	fmt.Println("Exerice level 1 #5")
	fmt.Println("type struct")
	type person struct {
		firstName        string
		lastName         string
		favoriteIceCream []string
	}

	p1 := person{
		firstName:        "igor",
		lastName:         "raspopovic",
		favoriteIceCream: []string{"rumraisn", "vanila", "chocolate"},
	}
	p2 := person{
		firstName:        "james",
		lastName:         "bond",
		favoriteIceCream: []string{"orio", "chocolatemint", "cokiedough"},
	}
	fmt.Println("p1 favourite ice creams are:")
	for _, v := range p1.favoriteIceCream {
		fmt.Println("\t", v)
	}
	fmt.Println("p2 favourite ice creams are:")
	for _, v := range p2.favoriteIceCream {
		fmt.Println("\t", v)
	}

	fmt.Println()

	fmt.Println("Exerice level 5 #2")
	fmt.Println("incorporate struct type into map")
	nameToPersonMap := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}
	fmt.Println(nameToPersonMap)

	for k, v := range nameToPersonMap {
		fmt.Println("Key: ", k, "\t", "Value: ", v)
		for _, slv := range v.favoriteIceCream {
			fmt.Println("\t", slv)
		}
	}
	fmt.Println()

	fmt.Println("Exerice level 5 #3")
	fmt.Println("incorporate struct type into another struct type. resembel inheritance")
	type vehicle struct {
		doors int
		color string
	}

	type truck struct {
		vehicle
		fourWheel bool
	}
	type sedan struct {
		vehicle
		luxury bool
	}

	myTruck := truck{
		vehicle: vehicle{
			doors: 5,
			color: "black",
		},
		fourWheel: true,
	}
	fmt.Println("Truck: ")
	fmt.Println("doors: ", myTruck.doors, "\tcolor: ", myTruck.color, "\tfourWheel: ", myTruck.fourWheel)

	mySedan := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "red",
		},
		luxury: true,
	}

	fmt.Println("Sedan: ")
	fmt.Println("doors: ", mySedan.doors, "\tcolor: ", mySedan.color, "\tluxury: ", mySedan.luxury)

	fmt.Println()

	fmt.Println("Exerice level 5 #4")
	fmt.Println("Create and use anonymous struct")
	anonymousStruct := struct {
		mapField   map[string]int
		sliceField []string
	}{
		mapField: map[string]int{
			"igor": 1966,
			"kaca": 1967,
			"mila": 2018,
		},
		sliceField: []string{"bla", "blah", "blahh"},
	}
	fmt.Println(anonymousStruct)
	fmt.Println("MapField: ", anonymousStruct.mapField, "\tsliceField: ", anonymousStruct.sliceField)
	for k, v := range anonymousStruct.mapField {
		fmt.Println(k, v)
	}
	for i, v := range anonymousStruct.sliceField {
		fmt.Println(i, v)
	}

	fmt.Println()

	//NOTES:
	//function - everyting in Go is passed by VALUE

	//variadic parameter has to the final paramete in the function

	//defer key word d

	//methods (reveiver part of the function definition)

	//GENERAL observarin - to declare anything (key word, identifier, type)
	//var myVar int
	//type Car struct
	//type Human interface

	//value can be of more than one type

	//interface (if your have this method you are of my `type` :)

	//empty inteface - any type implements empty inteface and therefore any type is of empty interface
	// so lots of function in fmt package like printLn function accepts (a ...inerface{}) so variadic number of parameters of any type

	//compsition in go - doc by william kennedy -read it
}
