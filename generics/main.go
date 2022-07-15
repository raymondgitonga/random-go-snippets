package main

import "fmt"

func main() {
	//print([]string{"Hello, ", "playground"})
	//print([]int{1, 2, 3})

	//var testAge int64 = 23
	//var testAge2 float64 = 24.5
	//
	//newGenericFunc(testAge)
	//newGenericFunc(testAge2)

	engineer := Engineer{Salary: 10}
	manager := Manager{Salary: 100}

	getSalary(engineer)
	getSalary(manager)

}
func print[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

type Age interface {
	int64 | int32 | float32 | float64
}

func newGenericFunc[age Age](myAge age) {
	val := int(myAge) + 1
	fmt.Println(val)
}

type Employee interface {
	PrintSalary()
}

func getSalary[E Employee](e E) {
	e.PrintSalary()
}

type Engineer struct {
	Salary int32
}

func (e Engineer) PrintSalary() {
	fmt.Println(e.Salary)
}

type Manager struct {
	Salary int64
}

func (m Manager) PrintSalary() {
	fmt.Println(m.Salary)
}
