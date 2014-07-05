package main

import (
	"fmt"
)

// interface 1
type person interface {
	getName() string
}

// interface 2
type personState interface {
	getState() string
	getZipcode() int64
}

// normalperson struct implements person interface & getState but not the statePerson's zipcode
// it has its own two fields that it uses in the implemented functions
type normalPerson struct {
	State string
	age   int
}

func (np normalPerson) getState() string {
	if np.State == "" {
		return "State is undefined"
	}
	return np.State
}

func (np normalPerson) getName() string { return "Name of CA-Resident" }

// another function implemented for normalPerson struct
func (np normalPerson) getAge() int { return np.age }

// state person implments both person and state interfaces
type statePerson struct {
	state string
	self  normalPerson
}

type functionalPerson struct {
	person      // inheriting the functions of person's interface
	personState // inheriting the functions of personState interface
}

type functionalStructPerson struct {
	normalPerson // inheriting the functions of normalPerson's struct
	//personState  // inheriting the functions of personState interface
}

func (sp statePerson) getState() string { return sp.state }

func main() {

	sp := statePerson{state: "CA-Resident", self: normalPerson{State: "Furguson Street, San Francisco , California-Resident", age: 20}}

	fmt.Println("State Person >>>", sp)
	fmt.Println("\t \tField literal state >> sp.getState() >>", sp.getState())
	fmt.Println("\t \tSelf-defined (NormalPerson) State >> sp.self.getState() >>", sp.self.getState())

	fmt.Println()
	sp = statePerson{state: "CA-Resident"}

	fmt.Println("State Person 2 >>>", sp)
	fmt.Println("\t \tField literal state >> sp.getState() >>", sp.getState())
	fmt.Println("\t \tSelf-defined (NormalPerson) State >> sp.self.getState() >>", sp.self.getState())

	fmt.Println()
	sp = statePerson{}

	fmt.Println("State Person 3 >>>", sp)
	fmt.Println("\t \tField literal state >> sp.getState() >>", sp.getState())
	fmt.Println("\t \tSelf-defined (NormalPerson) State >> sp.self.getState() >>", sp.self.getState())

	fmt.Println()
	fp := functionalPerson{}
	fmt.Println("Functional Person 3 >>>", fp)
	// both gives errors as theses are interfaces and no concrete implementations
	//fmt.Println("\t \tField literal state >> fp.getState() >>", fp.getState())
	//fmt.Println("\t \tSelf-defined (NormalPerson) State >> fp.self.getState() >>", fp.getName())

	fmt.Println()
	fsp := functionalStructPerson{}
	fmt.Println("Functional struct Person 1 >>>", fsp)
	// both gives errors as theses are interfaces and no concrete implementations
	fmt.Println("\t \tField literal state >> fsp.getState() >>", fsp.getState())
	fmt.Println("\t \tSelf-defined (NormalPerson) State >> fsp.self.getState() >>", fsp.getName())

}
