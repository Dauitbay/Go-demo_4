package main

import "fmt"

type account struct {
	login string
	password string
	url string
}

func main() {
	login := prompData("Enter your login: ")
	password := prompData("Enter your password")
	url := prompData("Enter your url")

	myAccount := account {
		password: password,
		url: url,
		login: login,
	}
	outputPassword(&myAccount)
}

func prompData(promp string) string {
	fmt.Print(promp + " ")
	var res string
	fmt.Scan(&res)
	return res
}

func outputPassword(acc *account) {
	fmt.Println(acc)
	fmt.Println(acc.login, acc.password, acc.url)
}
