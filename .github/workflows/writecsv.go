package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: writecsv.go <user> <pr>")
		os.Exit(1)
	}

	user := os.Args[1]
	pr := os.Args[2]

	if err := writecsv(user, pr); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func writecsv(user string, pr string) error {
	// Open the CSV file
	file, err := os.OpenFile("./.github/workflows/usersprs.csv", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a new CSV writer
	writer := csv.NewWriter(file)

	// Write the user and PR to the CSV file
	err = writer.Write([]string{user, pr})
	if err != nil {
		return err
	}

	// Flush the writer to ensure all data is written to the file
	writer.Flush()

	fmt.Printf("User %s's PR %s has been written to the CSV file\n", user, pr)
	return nil
}
