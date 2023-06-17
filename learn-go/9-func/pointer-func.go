package main

import "fmt"

func main() {
	y := 10
	update(&y)
	fmt.Println(y)

}
func update(p *int) {
	*p = 20
}
