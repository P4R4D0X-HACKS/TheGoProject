package main

import "fmt"

// Go is a pass by value language 

type contactInfo struct{
	email string
	zipCode int
}
// Initializing struts
type person struct {
	firstName string
	lastName string
	contact contactInfo
}

func main() {
	// alex := person{"Alex", "Anderson"} ONE WAY TO DEFINE 
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)
	// var alex person
	// fmt.Printf("%+v", alex)

	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"
	// fmt.Printf("%+v", alex)

	jim := person {
		firstName: "Jim",
		lastName: "Party",
		contact : contactInfo{
			email: "jim@gmail.com",
			zipCode: 741235,
		},
	}

	//jimPointer := &jim // gime me the memory address of the value this variale is pointing at
	// fmt.Println(jimPointer)
	//jimPointer.updateName("sam")
	jim.updateName("sam")
	//fmt.Printf("%+v", jim)
	jim.print()
}

// *pointer -> give me the value this memory address is pointing at 

// But but but *person here is just a type description it means we're working with a pointer to a person
// *pointerToPerson -> this is an operator it eans we want to manipulate the value the pointer is referencing 
// Turn address into value with *address
// Turn value into address with &value

func ( pointerToPerson *person  ) updateName(newFirstName string){
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print(){
	fmt.Printf("%+v", p)
}

/*
Value Type in go : int, float, string, bool, structs { use pointers to change these things in a function }
Reference Type in go : slices, maps, channels, pointers, functions { don't worry about pointers with these }
*/