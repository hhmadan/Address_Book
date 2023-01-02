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

// to delete contact from Contact List

func delete(ContactList *[]Contact) {
	var fName string
	fmt.Println("Enter First Name to delete")
	fmt.Scanln(&fName)
	for i, val := range *ContactList {
		if val.FirstName == fName {
			*ContactList = append((*ContactList)[:i], (*ContactList)[i+1:]...)
			fmt.Println("Contact Deleted Successfully..!")
			return
		}
	}
	fmt.Println("Contact Don't Exist..!")
}

/*
Logic: Func to find contacts of same City and display them.
Creating new Slice to store contacts of same City
*/
func searchByCity(ContactList *[]Contact) {
	var inputCityName string
	cityList := []Contact{}

	fmt.Println("Enter Name of City to search People: ")
	fmt.Scanln(&inputCityName)

	cityName := strings.TrimSpace(inputCityName)
	for _, contact := range *ContactList {
		if contact.City == cityName {
			cityList = append(cityList, contact)
		}
	}
	fmt.Println("City List of Persons: ", cityList)
}

/*
Logic: Func to find contacts of same State and display them.
Creating new Slice to store contacts of same State
*/
func searchByState(ContactList *[]Contact) {
	var inputStateName string
	stateList := []Contact{}

	fmt.Println("Enter Name of State to search People: ")
	fmt.Scanln(&inputStateName)

	stateName := strings.TrimSpace(inputStateName)
	for _, contact := range *ContactList {
		if contact.State == stateName {
			stateList = append(stateList, contact)
		}
	}
	fmt.Println("State List of Persons: ", stateList)
}

//To display all contacts in Contact List

func Display(ContactList *[]Contact) {
	fmt.Println("----Contact List----")
	for _, val := range *ContactList {
		fmt.Println(val)
	}
}

/*
	display to user for selecting required operations on contact list
*/

func menu(ContactList *[]Contact) {
	var menuOption int
	fmt.Println("----CONTACT MENU----\n1.Add Contact\n2.Delete Contact\n3.Search By City\n4.Search By State\n5.Display Contact List\n6.Exit")
	fmt.Scanln(&menuOption)
	switch menuOption {
	case 1:
		AddContact(ContactList)
	case 2:
		delete(ContactList)
	case 3:
		searchByCity(ContactList)
	case 4:
		searchByState(ContactList)
	case 5:
		Display(ContactList)
	case 6:
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
