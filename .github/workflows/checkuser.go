package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: checkuser <user> <pr>")
		os.Exit(1)
	}

	user := os.Args[1]
	pr := os.Args[2]

	if err := checkUser(user, pr); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func checkUser(user string, pr string) error {
	// Open the CSV file
	file, err := os.OpenFile("./.github/workflows/usersprs.csv", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Read the CSV file
	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		return err
	}

	// Check if the user exists in the CSV file
	for _, record := range records {
		if record[0] == user {
			if record[1] == pr {
				fmt.Println("Pass: User exists with the same PR number")
				return nil
			} else {
				fmt.Println("Fail: User exists with different PR number")
				return fmt.Errorf("User %s already exists with PR number %s", user, record[1])
			}
		}
	}

	fmt.Printf("User %s has not submitted a gentx PR yet\n", user)
	return nil
}
