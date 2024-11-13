package main

import (
	"demo/password/account"
	"demo/password/encrypter"
	"demo/password/files"
	"demo/password/output"
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

var menu = map[string]func(*account.VaultWithDb){
	"1":createAccount,
	"2":findAccountByUrl,
	"3":findAccountByLogin,
	"4":deleteAccount,
}

func main() {
	fmt.Println("Manager of passwords")
	err := godotenv.Load()
	if err != nil{
		output.PrintError("Could not find env file")
	}
	vault := account.NewVault(files.NewJsonDb("data.json"), *encrypter.NewEncrypter())
Menu:
	for {
		variant := prompData(
			"1. Create an account",
			"2. Search an account by Url",
			"3. Search an account by Login",
			"4. Delete an account",
			"5. Exit",
			"Choose the variant",
		)
		menuFunc := menu[variant]
		if menuFunc == nil{
			break Menu
		}
		menuFunc(vault)
		// switch variant {
		// case "1":
		// 	createAccount(vault)
		// case "2":
		// 	findAccount(vault)
		// case "3":
		// 	deleteAccount(vault)
		// default:
		// 	break Menu
		// }
	}
}


func findAccountByUrl(vault *account.VaultWithDb) {
	url := prompData("Enter url for search")
	accounts := vault.FindAccounts(url, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Url, str)
	})
	outPutResult(&accounts)
}

func findAccountByLogin(vault *account.VaultWithDb) {
	login := prompData("Enter url for search")
	accounts := vault.FindAccounts(login, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Login, str)
	})
	outPutResult(&accounts)
}

func outPutResult(accounts *[]account.Account){
	if len(*accounts) == 0 {
		color.Red("No accounts are found")
	}
	for _, account := range *accounts {
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

func prompData(promp ...any) string {
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
