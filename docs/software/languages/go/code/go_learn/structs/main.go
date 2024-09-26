package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {

	kobbi := person{
		firstName: "Kobbi",
		lastName:  "Gal",
		contactInfo: contactInfo{
			zip:   66653,
			email: "kgal@paloaltonetworks.com"},
	}

	kobbi.print()
	kobbi.updateName("Kobi")
	kobbi.print()

	kobbiMemoryAddress := &kobbi
	println(kobbiMemoryAddress)

	kobbiValue := *kobbiMemoryAddress
	fmt.Printf("%+v\n", kobbiValue)

}

func (pointerToPersonType *person) updateName(newFirstName string) {
	pointerToPersonType.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
