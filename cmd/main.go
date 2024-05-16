package main

import (
	"YadroTestCase/pkg/logger"
	"fmt"
	"os"
)
func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Invalid input argument. No such <filepath> argument")
		os.Exit(1)
	} 
	filePath := os.Args[1]
	club, err := logger.ParseFile(filePath)
	if err != nil {
		fmt.Printf("Paser error\n%s", err)
		os.Exit(1)
	}
	club.HandleEvents()
	club.PrintResults()
}
