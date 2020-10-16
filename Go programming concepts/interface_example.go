// Interfaces
// Interface describes the behaviour.
// Named collection of method signatures.
// Interface is only going to have method signatures.
// Whenever this method signature is implemented in any of the structure, then the structure is said to implement this interface.
package main

import "fmt"

type Human interface {
	getInfo()
	setInfo()
}

type Employee struct {
	id     int
	name   string
	salary int
}

type Student struct {
	rollNo int
	name   string
	total  int
}

func (S *Student) setInfo() {
	fmt.Scan(&S.rollNo, &S.name, &S.total)
}
func (E *Employee) setInfo() {
	fmt.Scan(&E.id, &E.name, &E.salary)
}
func (S Student) getInfo() {
	fmt.Println(S.rollNo, S.name, S.total)
}

func (E Employee) getInfo() {
	fmt.Println(E.id, E.name, E.salary)
}
func main() {
	/*
		stu := Student{}
		emp := Employee{}
		humans := []Human{&stu, &emp}
	*/
	humans := []Human{new(Student), new(Employee)}
	for _, hum := range humans {
		hum.setInfo()
		hum.getInfo()
	}
}
