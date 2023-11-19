package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func insertItem(itemName string) {
	// formattedItem := "=> " + itemName + "--------------------------------------------\n"
	toDoListFile, err := os.OpenFile("./todolist.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	checkNilError(err)
	defer toDoListFile.Close()

	_, writeError := io.WriteString(toDoListFile, itemName)
	checkNilError(writeError)

	fmt.Println("To-Do list has been updated!")
}

func deleteItem(itemNumber int) {

	toDoListFileRead, fileReadError := os.Open("./todolist.txt")
	checkNilError(fileReadError)

	scanner := bufio.NewScanner(toDoListFileRead)
	scanner.Split(bufio.ScanLines)

	var counter int = 0
	var newItems []string
	for scanner.Scan() {
		counter++
		if counter != itemNumber {
			newItems = append(newItems, scanner.Text())
		}
	}

	toDoListFileRead.Close()

	toDoListFile, fileError := os.OpenFile("./todolist.txt", os.O_WRONLY|os.O_TRUNC, 0644)
	checkNilError(fileError)
	defer toDoListFile.Close()

	for _, newItem := range newItems {
		_, writeError := toDoListFile.WriteString(newItem + "\n")
		checkNilError(writeError)
	}

}

func displayToDoList() {
	toDoListFile, fileError := os.Open("./todolist.txt")
	checkNilError(fileError)
	defer toDoListFile.Close()

	scanner := bufio.NewScanner(toDoListFile)
	scanner.Split(bufio.ScanLines)

	var toDoItems []string

	for scanner.Scan() {
		toDoItems = append(toDoItems, scanner.Text())
	}

	fmt.Println("To-Do List")
	for index, item := range toDoItems {
		fmt.Println("--------------------------------------------")
		formattedItems := strconv.Itoa(index+1) + ". " + item
		fmt.Println(formattedItems)
	}
	fmt.Println("--------------------------------------------")

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("======================= Welcome to To-Do List Application =======================")

	fmt.Println("Please select an option below: ")
	fmt.Println("1. Display List")
	fmt.Println("2. Insert Item")
	fmt.Println("3. Delete Item")

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter option: ")
	userSelectionStr, err := reader.ReadString('\n')
	checkNilError(err)

	userSelection, _ := strconv.Atoi(strings.TrimSpace(userSelectionStr))

	switch userSelection {
	case 1:
		displayToDoList()

	case 2:
		fmt.Printf("\nEnter the item to insert: ")
		userItem, err := reader.ReadString('\n')
		checkNilError(err)

		insertItem(userItem)
	case 3:
		fmt.Printf("\nEnter the item number to delete: ")
		userItemStr, err := reader.ReadString('\n')
		checkNilError(err)

		userItem, _ := strconv.Atoi(strings.TrimSpace(userItemStr))

		deleteItem(userItem)

	default:
		fmt.Println("Wrong Option selected. Please try again")
	}

}
