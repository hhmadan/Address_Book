package main

import (
	"fmt"
	"os"
	"strings"
)

type Contact struct {
	FirstName   string
	LastName    string
	Address     string
	City        string
	State       string
	PhoneNumber string
	Email       string
}

/*
	@paramters: contact list of Contact type
	logic: creating and adding contacts to contact list
*/

func AddContact(ContactList *[]Contact) {

	var option string
	var person Contact

	fmt.Println("Enter First Name: ")
	fmt.Scanln(&person.FirstName)

	fmt.Println("Enter Last Name: ")
	fmt.Scanln(&person.LastName)

	fmt.Println("Enter Address: ")
	fmt.Scanln(&person.Address)

	fmt.Println("Enter City: ")
	fmt.Scanln(&person.City)

	fmt.Println("Enter State: ")
	fmt.Scanln(&person.State)

	fmt.Println("Enter Phone Number: ")
	fmt.Scanln(&person.PhoneNumber)

	fmt.Println("Enter Email-Id: ")
	fmt.Scanln(&person.Email)

	*ContactList = append(*ContactList, person)

	fmt.Println("Want to Add More Contact? Press Y or N")
	fmt.Scanln(&option)

	if strings.ToUpper(option) == "Y" {
		AddContact(ContactList)
	} else {
		return
	}
}

/*
	display to user for selecting required operations on contact list
*/

func menu(ContactList *[]Contact) {
	var menuOption int
	fmt.Println("----CONTACT MENU----\n1.Add Contact\n2.Exit")
	fmt.Scanln(&menuOption)
	switch menuOption {
	case 1:
		fmt.Println("Enter First Name:")
		AddContact(ContactList)
	case 2:
		fmt.Println("Exiting...")
		os.Exit(0)
	}
	menu(ContactList)
}

func main() {
	fmt.Println("******WELCOME TO ADDRESS BOOK******")
	var ContactList = []Contact{}
	menu(&ContactList)
}
