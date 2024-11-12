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
		variant := prompData([]string{
			"1. Create an account",
			"2. Search an account",
			"3. Delete an account",
			"4. Exit",
			"Choose the variant",
		})
		switch variant {
		case "1":
			createAccount(vault)
		case "2":
			findAccount(vault)
		case "3":
			deleteAccount(vault)
		default:
			break Menu
		}
	}
	createAccount(vault)
}


func findAccount(vault *account.VaultWithDb) {
	url := prompData([]string{"Enter url for search"})
	accounts := vault.FindAccountsByUrl(url)
	if len(accounts) == 0 {
		color.Red("No accounts are found")
	}
	for _, account := range accounts {
		account.Output()
	}
}

func deleteAccount(vault *account.VaultWithDb) {
	url := prompData([]string{"Enter url for search"})
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
	login := prompData([]string{"Enter your login"})
	password := prompData([]string{"Enter your password"})
	url := prompData([]string{"Enter your url"})

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		output.PrintError("Invalid url or login")
		return
	}
	vault.AddAccount(*myAccount)
}

func prompData[T any](promp []T) string {
	for i, line := range promp {
		if i == len(promp)-1 {
			fmt.Printf("%v: ", line)
		} else {
			fmt.Println(line)
		}
	}
	var res string
	fmt.Scanln(&res)
	return res
}
