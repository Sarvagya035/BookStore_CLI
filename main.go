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

	if len(os.Args) < 2 {
		fmt.Println("expected 'get' subcommand")
		os.Exit(1)
	}

	switch os.Args[1] {

	case "get":
		HandleallBooks(getCmd, getall, getid)
	default:
		fmt.Println("Please provide get, update, update, delete commands")
		os.Exit(1)
	}

}
