package main

import (
	"flag"
	"fmt"
	"os"
)

type Book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Price    string `json:"price"`
	ImageUrl string `json:"imageurl"`
}

func main() {

	/*
		get books --id or --all
	*/

	getCmd := flag.NewFlagSet("get", flag.ExitOnError)
	getall := getCmd.Bool("all", false, "List all the books")
	getid := getCmd.String("id", "", "List book by particular Id")

	/*

		add books --id --title --author --price --imageurl
	*/

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addid := addCmd.String("id", "", "Unique ID of the book")
	addTitle := addCmd.String("title", "", "Title of the book")
	addAuthor := addCmd.String("author", "", "Author name of the book")
	addPrice := addCmd.String("price", "", "Price of the book")
	addimageurl := addCmd.String("imageurl", "", "imageurl name of the book")

	/*
		delete books --id
	*/

	delCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	delid := delCmd.String("id", "", "id of book to be deleted")

	/*

		update books --id --title --author --price --imageurl
	*/

	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
	updateid := updateCmd.String("id", "", "Unique ID of the book")
	updateTitle := updateCmd.String("title", "", "Title of the book")
	updateAuthor := updateCmd.String("author", "", "Author name of the book")
	updatePrice := updateCmd.String("price", "", "Price of the book")
	updateimageurl := updateCmd.String("imageurl", "", "imageurl name of the book")

	if len(os.Args) < 2 {
		fmt.Println("expected 'get' subcommand")
		os.Exit(1)
	}

	switch os.Args[1] {

	case "get":
		HandleallBooks(getCmd, getall, getid)
	case "add":
		AddBook(addCmd, addid, addTitle, addAuthor, addPrice, addimageurl)
	case "delete":
		Delbooks(delCmd, delid)
	case "update":
		UpdateBooks(updateCmd, updateid, updateTitle, updateAuthor, updatePrice, updateimageurl)

	default:
		fmt.Println("Please provide get, update, update, delete commands")
		os.Exit(1)
	}

}
