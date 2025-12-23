package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

func checkerr(err error) {

	if err != nil {
		fmt.Println("Error happened: ", err)
		os.Exit(1)
	}
}

// func to get all the books.
func getbooks() []Book {

	booksbyte, err := os.ReadFile("./books.json")

	checkerr(err)

	var books []Book

	err = json.Unmarshal(booksbyte, &books)

	return books
}

func HandleallBooks(getCmd *flag.FlagSet, all *bool, id *string) {

	getCmd.Parse(os.Args[2:])

	// Checking for id or string

	if !*all && *id == "" {

		fmt.Println("subcommand --all or --id needed")
		getCmd.PrintDefaults()
		os.Exit(1)
	}

	// if for all books return all books

	if *all {

		books := getbooks()

		fmt.Printf("Id \t Title \t Author \t Price \t ImageURL \n")

		for _, book := range books {

			fmt.Printf("%v \t %v \t %v \t %v \t %v \n", book.ID, book.Title, book.Author, book.Price, book.ImageUrl)
		}
	}

	if *id != "" {

		books := getbooks()
		fmt.Printf("Id \t Title \t Author \t Price \t ImageURL \n")

		// to check a book exist or not

		var foundBook bool
		for _, book := range books {
			foundBook = true
			if *id == book.ID {
				fmt.Printf("%v \t %v \t %v \t %v \t %v \n", book.ID, book.Title, book.Author, book.Price, book.ImageUrl)
			}
		}

		if !foundBook {
			fmt.Println("Book not found")
			os.Exit(1)
		}
	}

}
