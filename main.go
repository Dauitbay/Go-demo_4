package main

import (
	"fmt"
)

func main() {

	login := prompData("Enter your login: ")
	password := prompData("Enter your password")
	url := prompData("Enter your url")

	myAccount, err := newAccountWithTimeStamp(login, password, url)
	if err != nil {
		fmt.Println("Invalid url or login")
		return
	}
	myAccount.acc.outputPassword()
}

func prompData(promp string) string {
	fmt.Print(promp + " ")
	var res string
	fmt.Scanln(&res)
	return res
}
