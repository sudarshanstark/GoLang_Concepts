// Embedding or indirect inheritance
package main

import "fmt"

type Person struct {
	name  string
	phone int
	add   string
}

type Employee struct {
	Person // Embedding Person struct inside the Employee struct.
	id     int
	dept   int
	salary int
}

func (E *Employee) setInfo() {
	fmt.Scan(&E.name, &E.phone, &E.add, &E.id, &E.dept, &E.salary)
}

func (E Employee) getInfo() {
	fmt.Println("Name", E.name)
	fmt.Println("Id", E.id)
	fmt.Println("Dept", E.dept)
	fmt.Println("Phone number", E.phone)
	fmt.Println("Address", E.add)
}
func main() {
	emp := Employee{}
	emp.setInfo()
	emp.getInfo()
}
