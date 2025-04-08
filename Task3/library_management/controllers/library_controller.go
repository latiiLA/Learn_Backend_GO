package controllers

import (
	"bufio"
	"fmt"
	"library_management/models"
	"library_management/services"
	"os"
	"strconv"
	"strings"
)

type LibraryController struct {
	libraryService services.LibraryManager
}

func NewLibraryController(service services.LibraryManager) *LibraryController{
	return &LibraryController{libraryService: service}
}

func(lc *LibraryController) Run(){
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nLibrary Management System")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books by Member")
		fmt.Println("7. Add Member")
		fmt.Println("8. Exit")
		fmt.Print("Enter your choice: ")

		input, _ := reader.ReadString('\n')
		choice, err := strconv.Atoi(strings.TrimSpace(input))
		if err != nil{
			fmt.Println("Invalid input. Please enter a number")
			continue
		}

		switch choice{
		case 1:
			lc.addBook(reader)
		case 2:
			lc.removeBook(reader)
		case 3:
			lc.borrowBook(reader)
		case 4:
			lc.returnBook(reader)
		case 5:
			lc.listAvailableBooks()
		case 6: 
			lc.listBorrowedBooks(reader)
		case 7:
			lc.addMember(reader)
		case 8:
			fmt.Print("Exiting...")
			return
		default:
			fmt.Print("Invalid choice. Please try again.")
		}
	}

}

func (lc *LibraryController) addBook(reader *bufio.Reader) {
	fmt.Print("Enter Book ID: ")
	idInput, _ := reader.ReadString('\n')
	id, _ := strconv.Atoi(strings.TrimSpace(idInput))

	fmt.Print("Enter Book Title: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Print("Enter Book Author: ")
	author, _ := reader.ReadString('\n')
	author = strings.TrimSpace(author)

	book := models.Book{
		ID:     id,
		Title:  title,
		Author: author,
	}

	lc.libraryService.AddBook(book)
	fmt.Println("Book added successfully!")
}

func (lc *LibraryController) removeBook(reader *bufio.Reader) {
    fmt.Print("Enter Book ID to remove: ")
    idInput, _ := reader.ReadString('\n')
    id, _ := strconv.Atoi(strings.TrimSpace(idInput))

    err := lc.libraryService.RemoveBook(id)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Book removed successfully!")
}

func (lc *LibraryController) borrowBook(reader *bufio.Reader) {
	fmt.Print("Enter Book ID to borrow: ")
	bookIDInput, _ := reader.ReadString('\n')
	bookID, _ := strconv.Atoi(strings.TrimSpace(bookIDInput))

	fmt.Print("Enter Member ID: ")
	memberIDInput, _ := reader.ReadString('\n')
	memberID, _ := strconv.Atoi(strings.TrimSpace(memberIDInput))

	err := lc.libraryService.BorrowBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Book borrowed successfully!")
}

func (lc *LibraryController) returnBook(reader *bufio.Reader) {
	fmt.Print("Enter Book ID to return: ")
	bookIDInput, _ := reader.ReadString('\n')
	bookID, _ := strconv.Atoi(strings.TrimSpace(bookIDInput))

	fmt.Print("Enter Member ID: ")
	memberIDInput, _ := reader.ReadString('\n')
	memberID, _ := strconv.Atoi(strings.TrimSpace(memberIDInput))

	err := lc.libraryService.ReturnBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Book returned successfully!")
}

func (lc *LibraryController) listAvailableBooks() {
	books := lc.libraryService.ListAvailableBooks()
	if len(books) == 0 {
		fmt.Println("No available books in the library.")
		return
	}

	fmt.Println("\nAvailable Books:")
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}
}

func (lc *LibraryController) listBorrowedBooks(reader *bufio.Reader) {
	fmt.Print("Enter Member ID: ")
	memberIDInput, _ := reader.ReadString('\n')
	memberID, _ := strconv.Atoi(strings.TrimSpace(memberIDInput))

	books := lc.libraryService.ListBorrowedBooks(memberID)
	if len(books) == 0 {
		fmt.Println("No books borrowed by this member.")
		return
	}

	fmt.Println("\nBorrowed Books:")
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}
}

func (lc *LibraryController) addMember(reader *bufio.Reader) {
	fmt.Print("Enter Member ID: ")
	idInput, _ := reader.ReadString('\n')
	id, _ := strconv.Atoi(strings.TrimSpace(idInput))

	fmt.Print("Enter Member Name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	library, ok := lc.libraryService.(*services.Library)
	if !ok {
		fmt.Println("Error: cannot add member with current service implementation")
		return
	}

	member := models.Member{
		ID:   id,
		Name: name,
	}
	library.Members[id] = member
	fmt.Println("Member added successfully!")
}
