package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"Assignment1_TuknazarovRakhat/Banking"
	"Assignment1_TuknazarovRakhat/Employees"
	"Assignment1_TuknazarovRakhat/Library"
	"Assignment1_TuknazarovRakhat/Shapes"
)

func main() {
	// ----------------------------
	// Exercise 1 Demonstration
	// ----------------------------
	lib := Library.NewLibrary()
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Exercise 1: Library Management")
	for {
		fmt.Println("Choose an option: Add | Borrow | Return | List | Exit")
		line, _ := reader.ReadString('\n')
		cmd := strings.TrimSpace(line)
		switch strings.ToLower(cmd) {
		case "add":
			fmt.Print("Enter Book ID: ")
			id, _ := reader.ReadString('\n')
			id = strings.TrimSpace(id)
			fmt.Print("Enter Title: ")
			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)
			fmt.Print("Enter Author: ")
			author, _ := reader.ReadString('\n')
			author = strings.TrimSpace(author)
			b := Library.Book{ID: id, Title: title, Author: author, IsBorrowed: false}
			lib.AddBook(b)
		case "borrow":
			fmt.Print("Enter Book ID to borrow: ")
			id, _ := reader.ReadString('\n')
			id = strings.TrimSpace(id)
			lib.BorrowBook(id)
		case "return":
			fmt.Print("Enter Book ID to return: ")
			id, _ := reader.ReadString('\n')
			id = strings.TrimSpace(id)
			lib.ReturnBook(id)
		case "list":
			lib.ListBooks()
		case "exit":
			goto AFTER_LIB
		default:
			fmt.Println("Unknown command.")
		}
	}

AFTER_LIB:
	// ----------------------------
	// Exercise 2 Demonstration
	// ----------------------------
	fmt.Println("Exercise 2: Shapes")
	shapes := []Shapes.Shape{
		Shapes.Rectangle{Length: 5, Width: 3},
		Shapes.Circle{Radius: 4},
		Shapes.Square{Length: 2.5},
		Shapes.Triangle{SideA: 3, SideB: 4, SideC: 5},
	}

	for _, s := range shapes {
		Shapes.PrintShapeDetails(s)
	}

	// ----------------------------
	// Exercise 3 Demonstration
	// ----------------------------
	fmt.Println("Exercise 3: Employee Management")
	comp := Employees.NewCompany()
	fte := Employees.FullTimeEmployee{ID: 1, Name: "John Doe", Salary: 300000}
	pte := Employees.PartTimeEmployee{ID: 2, Name: "Jane Smith", HourlyRate: 5000, HoursWorked: 20.5}

	comp.AddEmployee("emp1", fte)
	comp.AddEmployee("emp2", pte)
	comp.ListEmployees()

	// ----------------------------
	// Exercise 4 Demonstration
	// ----------------------------
	fmt.Println("Exercise 4: Bank Account System")
	acc := &Banking.BankAccount{AccountNumber: "1234567890", HolderName: "John Doe", Balance: 10000}
	for {
		fmt.Println("Choose an option: Deposit | Withdraw | Get balance | Exit")
		line, _ := reader.ReadString('\n')
		cmd := strings.TrimSpace(line)
		switch strings.ToLower(cmd) {
		case "deposit":
			fmt.Print("Enter amount: ")
			amtStr, _ := reader.ReadString('\n')
			amtStr = strings.TrimSpace(amtStr)
			amt, _ := strconv.ParseFloat(amtStr, 64)
			acc.Deposit(amt)
		case "withdraw":
			fmt.Print("Enter amount: ")
			amtStr, _ := reader.ReadString('\n')
			amtStr = strings.TrimSpace(amtStr)
			amt, _ := strconv.ParseFloat(amtStr, 64)
			acc.Withdraw(amt)
		case "get balance":
			acc.GetBalance()
		case "exit":
			goto END
		default:
			fmt.Println("Unknown command.")
		}
	}

END:
	fmt.Println("Program finished.")
}
