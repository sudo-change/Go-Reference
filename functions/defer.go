package main

import "fmt"

func queryDatabase(query string) string {
	var result string
	connectDatabase()
	// defer tells GO to run a function at the end of the current function
	defer disconnectDatabase()

	if query == "SELECT * FROM coolTable;" {
		result = "NAME|DOB\nVincent Van Gopher|March 30, 1853"
	}
	fmt.Println(result)
	return result
}

func connectDatabase() {
	fmt.Println("connecting to the database.")
}

func disconnectDatabase() {
	fmt.Println("disconnecting from the database")
}

func main() {
	queryDatabase("SELECT * FROM coolTable;")
}
