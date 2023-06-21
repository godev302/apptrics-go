package main

import "fmt"

type expense interface {
	calc() float64
}
type perm struct {
	empId   int
	empName string
	salary  float64
	bonus   float64
}

func (p perm) calc() float64 {
	return p.salary + p.bonus
}

type contract struct {
	emailId string
	pay     float64
}

func (c contract) calc() float64 {
	return c.pay
}

func TotalExpense(e ...expense) {
	var sum float64
	for _, emp := range e {
		sum = sum + emp.calc()
	}
	fmt.Println(sum)
}

func main() {
	p := perm{
		salary: 10,
		bonus:  10,
	}

	p1 := perm{
		salary: 10,
		bonus:  10,
	}

	c := contract{
		pay: 10,
	}
	c1 := contract{
		pay: 10,
	}

	TotalExpense(p, p1, c, c1)
}
