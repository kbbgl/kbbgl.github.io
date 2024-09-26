# Struct

Data structure, collection of different properties that are related. Very similar to `object`/`dict`.

## Define

```go
package main
type contactInfo struct {
  email string
  zip   int
}

type person struct {
  firstName string
  lastName string
}

// ...
```

## Creating an Instance

```go
func main() {
  kobbi := person{firstName: "Kobbi", lastName:  "Gal"}
  
  // will assign zero-value ("")
  var kobbi person
}
```

## Updating Values

```go
kobbi.firstName = "New Name"
```

## Nested Structs

```go
kobbi := person{
  firstName: "Kobbi",
  lastName:  "Gal",
  contactInfo:  contactInfo{
    zip:   66653,
    email: "kgal@paloaltonetworks.com"}
```

## Receiver Functions

```go
func (p person) print() {
  fmt.Printf("%+v", p)
}

kobbi.print()
```

Since Go is a pass by value language, we cannot update a `struct` from a receiver function directly, we would need to reference the memory address.

```go
// Won't update the firstName
func (p person) updateName(newFirstName string){
  p.firstName = newFirstName
}

// will update firstName, reference to pointer

func (p *person) updateName(newFirstName string) {
  p.firstName = newFirstName
}

kobbi.updateName("Kobi")
kobbi.print(kobbi.firstName)
// Kobbi
```

## Pointer Operators

```go
// get the memory address of variable 'kobbi'
kobbiMemoryAddress := &kobbi
println(kobbiMemoryAddress)

// 0xc000092f08

// get value of the variable's memory address
kobbiValue = *kobbiMemoryAddress
fmt.Printf("%+v\\", kobbiValue)
// {firstName:Kobi lastName:Gal contactInfo:{email:kgal@paloaltonetworks.com zip:66653}}
```
