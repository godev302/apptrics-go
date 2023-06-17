package database

import "fmt"

// Config struct stores a database connection, we would use this struct to safely store our connections
type Config struct {
	db string // unexported
}

func (c Config) ReadData() {
	fmt.Println("reading data from", c.db)
}

// NewConfig initialise the config struct with the values provided by the user
func NewConfig(userName string, password string, dbName string) (Config, bool) {
	if userName == "" || password == "" || dbName == "" {
		return Config{}, false // error happened, so we would return an empty struct
	}

	dbConnection := userName + password + dbName // assuming that this is how we open connection

	c := Config{db: dbConnection} // assigning database connection in the struct field

	return c, true // returning config struct

}
