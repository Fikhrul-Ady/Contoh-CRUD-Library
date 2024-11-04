package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var books []string

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Choose an option:")
		fmt.Println("1. Add New Book")
		fmt.Println("2. View All Books")
		fmt.Println("3. Delete Book")
		fmt.Println("4. Exit")
		fmt.Print("Insert your choice: ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			addBook(scanner)
		case "2":
			viewBooks()
		case "3":
			deleteBook(scanner)
		case "4":
			fmt.Println("Exiting the program. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice! Please try again 🙏")
		}
	}
}

func addBook(scanner *bufio.Scanner) {
	fmt.Print("Enter a new book title: ")
	scanner.Scan()
	book := scanner.Text()
	books = append(books, book)
	fmt.Println("Book successfully added!")
	viewBooks() // Display updated book list
}

func viewBooks() {
	if len(books) == 0 {
		fmt.Println("No books in the list.")
		return
	}
	fmt.Println("Booklist:")
	for i, book := range books {
		fmt.Printf("%d. Title: `%s`\n", i+1, book)
	}
}

func deleteBook(scanner *bufio.Scanner) {
	viewBooks()
	if len(books) == 0 {
		return // No books to delete
	}
	fmt.Print("Enter the number of the book you want to delete: ")
	scanner.Scan()
	numStr := scanner.Text()
	num, err := strconv.Atoi(numStr)
	if err != nil || num < 1 || num > len(books) {
		fmt.Println("Invalid book number.")
		return
	}
	books = append(books[:num-1], books[num:]...)
	fmt.Println("The book has been deleted!")
	viewBooks() // Display updated book list
}
