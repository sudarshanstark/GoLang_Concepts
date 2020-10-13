// Composition
package main

import "fmt"

type Date struct {
	day   int
	month string
	year  int
}

func (d *Date) setDate() {
	fmt.Scan(&d.day, &d.month, &d.year)
}

func (d Date) getDate() {
	fmt.Println("Date of birth", d.day, d.month, d.year)
}

type Person struct {
	name    string
	phone   int
	address string
	dob     Date // Composition
}

func (P *Person) setPer() {
	fmt.Scan(&P.name, &P.phone, &P.address)
	P.dob.setDate()
}

func (P Person) getPer() {
	fmt.Println("Name : ", P.name)
	P.dob.getDate()
	fmt.Println("Phone number : ", P.phone)
	fmt.Println("Address : ", P.address)
}
func main() {
	per := Person{}
	per.setPer()
	per.getPer()
}
