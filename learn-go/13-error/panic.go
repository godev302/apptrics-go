package main

import (
	"database/sql"
	"fmt"
	"runtime/debug"
)

func main() {

	//var i []int
	//
	//i[10] = 100
	startApp(nil)
	fmt.Println("end of main")

}
func recoverFunc() {
	r := recover() // nil means no panic // otherwise r would be the msg of the panic
	if r != nil {
		fmt.Println("recovered from the panic", r)
		fmt.Println(string(debug.Stack()))
	}
}

func startApp(db *sql.DB) {
	defer recoverFunc()
	fmt.Println("set up all the required things like db connection")
	//let say db conn failed at the start app level
	if db == nil {
		panic("db connection is not working")
	}

}

// create a func that accepts one number, and if number is less than 5 use panic
