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
