// Interface of Parameter type/argument type.
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

func process(H Human) {
	H.setInfo()
	H.getInfo()
}

func main() {
	humans := []Human{new(Student), new(Employee)}
	for _, hum := range humans {
		process(hum)
	}
}
