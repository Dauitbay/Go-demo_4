package main

import (
	"demo/password/account"
	"demo/password/files"
	"fmt"
)

func main() {
	files.WriteFile("Hello i am file", "file.txt")
	login := prompData("Enter your login: ")
	password := prompData("Enter your password")
	url := prompData("Enter your url")

	myAccount, err := account.NewAccountWithTimeStamp(login, password, url)
	if err != nil {
		fmt.Println("Invalid url or login")
		return
	}
	myAccount.OutputPassword()
	
	fmt.Println(myAccount)
}

func prompData(promp string) string {
	fmt.Print(promp + " ")
	var res string
	fmt.Scanln(&res)
	return res
}
