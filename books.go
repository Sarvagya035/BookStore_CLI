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

func savebooks(books []Book) error {

	booksbyte, err := json.Marshal(books)
	checkerr(err)

	err = os.WriteFile("./books.json", booksbyte, 0644)
	checkerr(err)

	return err
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

func AddBook(addCmd *flag.FlagSet, id *string, title *string, author *string, price *string, imageurl *string) {

	addCmd.Parse(os.Args[2:])

	if *id == "" || *title == "" || *author == "" || *price == "" || *imageurl == "" {

		fmt.Println("subcommand --all or --id needed")
		addCmd.PrintDefaults()
		os.Exit(1)
	}

	books := getbooks()

	var newBook Book

	newBook = Book{*id, *title, *author, *price, *imageurl}
	books = append(books, newBook)

	err := savebooks(books)
	checkerr(err)
	fmt.Println("New book added sucessfully")
}
