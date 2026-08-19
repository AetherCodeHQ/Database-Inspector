package main

import (
	"fmt"
	"os"
)

// database_inspector - Inspect database structures
func database_inspector(path string) {
	fmt.Println("========================================")
	fmt.Println("  Database-Inspector")
	fmt.Println("  Inspect database structures")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Target:", path)
	fmt.Println("Processing...")
	fmt.Println("Done!")
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	database_inspector(path)
}
