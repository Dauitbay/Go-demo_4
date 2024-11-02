package main

import (
	"demo/password/account"
	"demo/password/files"
	"fmt"
)

func main() {
	createAccount()
}

func createAccount() {
	// files.WriteFile("Hello i am file", "file.txt")
	// files.ReadFile()
	login := prompData("Enter your login: ")
	password := prompData("Enter your password")
	url := prompData("Enter your url")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Invalid url or login")
		return
	}
	file, err := myAccount.ToBytes()
	if err != nil {
		fmt.Println("Could not convert to JSON")
		return
	}
	files.WriteFile(file, "data.json")
}

func prompData(promp string) string {
	fmt.Print(promp + " ")
	var res string
	fmt.Scanln(&res)
	return res
}
