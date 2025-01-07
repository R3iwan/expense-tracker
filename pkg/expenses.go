package pkg

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/r3iwan/expense-tracker/pkg/models"
)

var Expenses []models.Expense
var nextID = 1

func TrackerExpense() {
	for {
		var command string
		reader := bufio.NewReader(os.Stdin)
		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("error reading a command")
		}
		command = strings.TrimSpace(command)

		parts := strings.Split(command, " ")
		if len(parts) == 0 {
			fmt.Println("Invalid command")
			continue
		}

		cmd := parts[0]
		args := parts[1:]

		switch cmd {
		case "add":
			addExpenses(args)
		case "list":
			listExpenses()
		case "delete":
			deleteExpense(args)
		case "summary":
			summaryExpenses()
		case "exit":
			fmt.Println("Exitting command")
			return
		default:
			fmt.Println("Unknown command. Supported commands: add")
		}
	}
}

func addExpenses(args []string) {
	var expense models.Expense

	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--description":
			if i+1 < len(args) {
				expense.Description = args[i+1]
				i++
			}
		case "--amount":
			if i+1 < len(args) {
				amount, err := strconv.ParseFloat(args[i+1], 64)
				if err != nil {
					fmt.Println("Invalid amount")
					return
				}
				expense.Amount = amount
				i++
			}
		}
	}

	expense.ID = nextID
	expense.Date = time.Now().Format("2006-01-02")
	Expenses = append(Expenses, expense)
	nextID++

	fmt.Println("Expense added sucessfully")

	err := saveFileToJSON("expenses.json")
	if err != nil {
		fmt.Println("Error saving to expenses.json file")
	}
}

func listExpenses() {
	if len(Expenses) == 0 {
		fmt.Println("Expenses is empty")
	}

	fmt.Println("ID\tDate\tDescription\tAmount")
	for _, expense := range Expenses {
		fmt.Printf("%d\t%s\t%s\t%.2f\n", expense.ID, expense.Date, expense.Description, expense.Amount)
	}
}

func deleteExpense(args []string) {
	var idToDelete int
	foundID := false

	for i := 0; i < len(args); i++ {
		if args[i] == "--id" {
			if i+1 < len(args) {
				id, err := strconv.Atoi(args[i+1])
				if err != nil {
					fmt.Println("Couldn't convert into int")

				}
				idToDelete = id
				foundID = true
				break
			}
		}
	}

	if !foundID {
		fmt.Println("Error: Missing --id flag.")
		return
	}

	for index, expense := range Expenses {
		if expense.ID == idToDelete {
			Expenses = append(Expenses[:index], Expenses[index+1:]...)

			err := saveFileToJSON("expenses.json")
			if err != nil {
				fmt.Println("Error saving to expenses.json file:", err)
			}
			return
		}
	}
}

func summaryExpenses() {
	var sum float64 = 0

	for _, expense := range Expenses {
		sum = sum + expense.Amount
	}

	fmt.Printf("Total expenses: %.2f", sum)
}
