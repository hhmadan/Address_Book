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
	Filter Function to get Duplicate Contacts
*/

func duplicateCheckFilter(ContactList *[]Contact, f func(string) bool) []Contact {
	filteredDuplicateNames := []Contact{}
	for _, fName := range *ContactList {
		if f(fName.FirstName) {
			filteredDuplicateNames = append(filteredDuplicateNames, fName)
		}
	}
	return filteredDuplicateNames
}

/*
	@paramters: contact list of Contact type
	logic: creating and adding contacts to contact list
*/

func AddContact(ContactList *[]Contact) {

	var option string
	var person Contact
	var firstName string

	fmt.Println("Enter First Name: ")
	fmt.Scanln(&firstName)

	checkDuplicate := duplicateCheckFilter(ContactList, func(fName string) bool {
		if fName == firstName {
			return true
		}
		return false
	})
	if len(checkDuplicate) > 0 {
		fmt.Println("Contact Already Exists..!")

	} else {
		person.FirstName = firstName

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
@Logic: Func to find contacts of same State and display them.
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

//Function to Edit an existing Contact

func editContact(ContactList *[]Contact, nameToChange string) {

	var newFName, newLName, newAddress, newCity, newState, newPhoneNum, newEmail string
	var editChoice int

	for i, contact := range *ContactList {
		if contact.FirstName == nameToChange {
			fmt.Println("1.First Name\n2.Last Name\n3.Address\n4.City\n5.State\n6.Phone Number\n7.Email-Id")
			fmt.Println("Enter choice to Edit : ")
			fmt.Scanln(&editChoice)
			switch editChoice {
			case 1:
				fmt.Println("New First Name: ")
				fmt.Scanln(&newFName)
				(*ContactList)[i].FirstName = newFName
			case 2:
				fmt.Println("New Last Name: ")
				fmt.Scanln(&newLName)
				(*ContactList)[i].LastName = newLName
			case 3:
				fmt.Println("New Address: ")
				fmt.Scanln(&newAddress)
				(*ContactList)[i].Address = newAddress
			case 4:
				fmt.Println("New City: ")
				fmt.Scanln(&newCity)
				(*ContactList)[i].City = newCity
			case 5:
				fmt.Println("New State: ")
				fmt.Scanln(&newState)
				(*ContactList)[i].State = newState
			case 6:
				fmt.Println("New Phone Number: ")
				fmt.Scanln(&newPhoneNum)
				(*ContactList)[i].PhoneNumber = newPhoneNum
			case 7:
				fmt.Println("New Email-Id: ")
				fmt.Scanln(&newEmail)
				(*ContactList)[i].Email = newEmail
			}
		}
	}
}

//To display all contacts in Contact List

func Display(ContactList *[]Contact) {
	fmt.Println("\n----Contact List----")
	for _, val := range *ContactList {
		fmt.Println(val)
	}
}

/*
	display to user for selecting required operations on contact list
*/

func menu(ContactList *[]Contact) {
	var menuOption int
	fmt.Println("----CONTACT MENU----\n1.Add Contact\n2.Delete Contact\n3.Edit Contact\n4.Search By City\n5.Search By State\n6.Display Contact List\n7.Exit")
	fmt.Scanln(&menuOption)
	switch menuOption {
	case 1:
		AddContact(ContactList)
	case 2:
		delete(ContactList)
	case 3:
		var nameToChange string
		fmt.Println("Enter First Name to edit :")
		fmt.Scanln(&nameToChange)

		for _, v := range *ContactList {
			if v.FirstName == nameToChange {
				editContact(ContactList, nameToChange)
			} else {
				fmt.Println("No Contact Found..!")
			}
		}
	case 4:
		searchByCity(ContactList)
	case 5:
		searchByState(ContactList)
	case 6:
		Display(ContactList)
	case 7:
		fmt.Println("Exiting...")
		os.Exit(0)
	}
	menu(ContactList)
}

func main() {
	fmt.Println("\n******WELCOME TO ADDRESS BOOK******")
	var ContactList = []Contact{}
	menu(&ContactList)
}
