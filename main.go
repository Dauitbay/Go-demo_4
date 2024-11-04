package main

import (
	"demo/password/account"
	"demo/password/files"
	"fmt"
)

func main() {
	fmt.Println("Manager of passwords")
Menu:
	for {
		variant := getMenu()
		switch variant {
		case 1:
			createAccount()
		case 2:
			findAccount()
		case 3:
			deleteAccount()
		default:
			break Menu
		}
	}
	createAccount()
}

func getMenu() int {
	var variant int
	fmt.Println("Choose variant")
	fmt.Println("1. Create an account")
	fmt.Println("2. Search an account")
	fmt.Println("3. Delete an account")
	fmt.Println("4. Exit")
	fmt.Scanln(&variant)
	return variant
}

func findAccount() {

}
func deleteAccount() {

}
func createAccount() {
	// files.WriteFile("Hello i am file", "file.txt")
	// files.ReadFile()
	login := prompData("Enter your login")
	password := prompData("Enter your password")
	url := prompData("Enter your url")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Invalid url or login")
		return
	}
	vault := account.NewVault()
	vault.AddAccount(*myAccount)
	data, err := vault.ToBytes()
	if err != nil {
		fmt.Println("Could not convert to JSON")
		return
	}
	files.WriteFile(data, "data.json")
}

func prompData(promp string) string {
	fmt.Println(promp + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}
