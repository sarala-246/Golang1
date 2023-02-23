// Calling function multiple times

// package main
// import ("fmt")

// func myMessage() {
//   fmt.Println("Welcome to Go!")
// }

// func main() {
//   myMessage()
//   myMessage()
//   myMessage()
// }


// Functions with arguments

// package main
// import "fmt"

// func main(){
// 	fmt.Println("starts func")
// 	gm()
// 	ga()
// 	ge()
// 	DisplayMessage("welcome to the golang world")
// 	DisplayMessage("Hello")
// 	DisplayMessage("team")
// 	fmt.Println("Ends func")
	
// 	}


// func ge(){
// 	fmt.Println("good evening all of you")
	
// 	}
	

// func gm(){
// 	fmt.Println("good morning all of you")
	
// 	}
	
// func ga(){
// 	fmt.Println("good afternoon all of you")
	
// 	}


// func DisplayMessage(s string) {

// 	fmt.Println("the message is :",s)
		
// }


// fuctions with return type

// package main
// import "fmt"
// func main(){
// 	var res int
// 	res=add(100,30)
// 	fmt.Println("the addition of two values :",res)
// 	}
	
// func add(x int,y int) int {
// 	var z int
// 	z=x+y
// 	return z
// }


// Function with multiple parameters

// package main
// import ("fmt")

// func familyName(fname string, age int) {
//   fmt.Println("Hello", age, "year old", fname)
// }

// func main() {
//   familyName("sarala", 5)
//   familyName("Jenny", 10)
//   familyName("Anjali", 10)
// }

/////////////////////////////////////////

// package main

// import "fmt"

// func main() {

// 	var arr [4]int
// 	fmt.Println("Elements of the Array: ", arr)

// }

///////////////////////////////////////////////////////////////////



// package main
  
// import "fmt"
  
// func main() {
  
//     arr := [4][4]int{{1,2,3,4}, {10,100,110,120},
//         {20,30,50,70}, {3,6,9,7}}
  
//     fmt.Println("Elements of Array",arr)
// }



////////////////////////////////////////////////////////////////



// 1 - dimensional array


// package main
  
// import "fmt"
  
// func main() {
  
// arr:= [5]int{10,20,30,40,50}

// fmt.Println("Elements of the array:",arr)

// }


///////////////////////////////////////////////////////////////////


// package main
  
// import "fmt"
  
// func main() {
     
// arr1:= [2]string{"sarala", "Tammireddy"}
// arr2:= [3]string{"Hello", "Welcome", "Golang"}
// arr3:= [...]int{1,2,3,4,5,6,}
  
// fmt.Println("Length of the array 1 is:", len(arr1), arr1)
// fmt.Println("Length of the array 2 is:", len(arr2), arr2)
// fmt.Println("Length of the array 3 is:", len(arr3), arr3)
// }

///////////////////////////////////////



// Creating Slice

// package main

// import "fmt"

// func main() {
	
// 	arr := [...]string{"S", "A", "R", "A", "L", "A"}

// 	fmt.Println("Array:", arr)

// 	Slice1 := arr[1:4]

// 	fmt.Println("Slice:", Slice1)

// 	fmt.Printf("Length of the array: %d", len(arr))

// 	fmt.Printf("\nLength of the Slice: %d", len(Slice1))

// 	fmt.Printf("\nCapacity of the Slice: %d", cap(Slice1))

// 	fmt.Printf("\nCapacity of the array: %d", cap(arr))

// }

//////////////////////////////////////////////////////////////////


// Creating slice from the slice and using var keyword and shorthand methods


// package main

// import "fmt"

// func main() {
// 	Slice := []int{2,4,6,8}

// 	var slice1 = Slice[0:1]
// 	slice2 := Slice[1:2]
// 	var slice3 = Slice[1:4]
// 	slice4 := slice3[1:2]

// 	fmt.Println("Original Slice:", Slice)
// 	fmt.Println("New slice1:", slice1)
// 	fmt.Println("New slice2:", slice2)
// 	fmt.Println("New slice3:", slice3)
// 	fmt.Println("New slice4:", slice4)

// }

///////////////////////////////////////////////////////////////////////////

// Slice using make function

// package main
 
// import "fmt"
 
// func main() {

// 	var slice1 = make([]int, 2, 5)
//     fmt.Printf("slice1 = %v, \nlength = %d, \ncapacity = %d\n",
// 	slice1, len(slice1), cap(slice1))

	
// 	slice2 := make([]int, 4)
//     fmt.Printf("slice2 = %v, \nlength = %d, \ncapacity = %d\n",
// 	slice2, len(slice2), cap(slice2))
// }	

/////////////////////////////////////////////////////////////////////////////////


// Sorting Method


// package main
 
// import (
//     "fmt"
//     "sort"
// )
 
// func main() {

// 	slc1 := []int{9,4,3,2,6,4,5}
// 	slc2 := []string{"H","E","L","L","O"}

// 	fmt.Println("Before sorting:")
// 	fmt.Println("Slice 1: ", slc1)
// 	fmt.Println("Slice 2: ", slc2)

// 	sort.Strings(slc2)
// 	sort.Ints(slc1)

	
//     fmt.Println("\nAfter sorting:")
//     fmt.Println("Slice 1: ", slc1)
//     fmt.Println("Slice 2: ", slc2)
// }

///////////////////////////////////////////////////

// comparing or check the equality of two slice

// package main
  
// import (
//     "bytes"
//     "fmt"
// )
  
// func main() {

//     slc1 := []byte{'1', '2', '3', 'a'}
      
//     slc2 := []byte{'1', '2', '3', 'a'}
      
//     slc3 := []byte{'1', '2', '3', 'A',}

//     r1 := bytes.Equal(slc1, slc2)
//     r2 := bytes.Equal(slc1, slc3)
//     r3 := bytes.Equal(slc2, slc3)
    
//     fmt.Println("Result 1:", r1)
//     fmt.Println("Result 2:", r2)
//     fmt.Println("Result 3:", r3)
    
// }

/////////////////////////////////////////////////////////////

// Converting lower case to upper case and vice versa

// package main
// import (
// 	"bytes"
// 	"fmt"
// )
// func main() {
// 	s1 := []byte{'s', 'a', 'r', 'a', 'l', 'a'}
// 	s2 := []byte{'T', 'A', 'M', 'M', 'I', 'R', 'E', 'D', 'D', 'Y'}

// 	fmt.Println("Original slice:")
// 	fmt.Printf("Slice 1: %s", s1)
// 	fmt.Printf("\nSlice 2: %s", s2)

// 	r1 := bytes.ToUpper(s1)
// 	r2 := bytes.ToLower(s2)
// 	r3 := bytes.ToUpper([]byte("HeLLo"))
// 	r4 := bytes.ToLower([]byte("WeLcoMe"))

// 	fmt.Printf("\n\nNew Slice:")
// 	fmt.Printf("\nSlice 1: %s", r1)
// 	fmt.Printf("\nSlice 2: %s", r2)
// 	fmt.Printf("\nSlice 3: %s", r3)
// 	fmt.Printf("\nSlice 4: %s", r4)
// }

////////////////////////////////////////////////////

// checking specified element in slice

// package main

// import (
// 	"bytes"
// 	"fmt"
// )

// func main() {

// 	s1 := []byte{'h','e','l','l','o'}

// 	r1 := bytes.Contains(s1, []byte{'v'})
// 	r2 := bytes.Contains([]byte("Welcome Golang"), []byte("lang"))
// 	r3 := bytes.Contains(s1, []byte{'l'})

// 	fmt.Println("Result 1:", r1)
// 	fmt.Println("Result 2:", r2)
// 	fmt.Println("Result 3:", r3)

// }