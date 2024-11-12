package main

import (
	"demo/password/account"
	"demo/password/files"
	"demo/password/output"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	fmt.Println("Manager of passwords")
	vault := account.NewVault(files.NewJsonDb("data.json"))
Menu:
	for {
		variant := getMenu()
		switch variant {
		case 1:
			createAccount(vault)
		case 2:
			findAccount(vault)
		case 3:
			deleteAccount(vault)
		default:
			break Menu
		}
	}
	createAccount(vault)
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

func findAccount(vault *account.VaultWithDb) {
	url := prompData("Enter url for search")
	accounts := vault.FindAccountsByUrl(url)
	if len(accounts) == 0 {
		color.Red("No accounts are found")
	}
	for _, account := range accounts {
		account.Output()
	}
}

func deleteAccount(vault *account.VaultWithDb) {
	url := prompData("Enter url for search")
	isDeleted := vault.DeleteAccountByUrl(url)
	if isDeleted {
		color.Green("Deleted")
	} else {
		output.PrintError("Not found")
	}
}
func createAccount(vault *account.VaultWithDb) {
	// files.WriteFile("Hello i am file", "file.txt")
	// files.ReadFile()
	login := prompData("Enter your login")
	password := prompData("Enter your password")
	url := prompData("Enter your url")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		output.PrintError("Invalid url or login")
		return
	}
	vault.AddAccount(*myAccount)
}

func prompData(promp string) string {
	fmt.Println(promp + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}
