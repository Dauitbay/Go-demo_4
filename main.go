package main

import (
	"fmt"
	"math/rand/v2"
)

type account struct {
	login string
	password string
	url string
}

func (acc *account) outputPassword() {
	fmt.Println(acc)
	fmt.Println(acc.login, acc.password, acc.url)
}

func (acc *account) generatePassword(n int) {
	res := make([]rune , n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.password = string(res)
}

var letterRunes = []rune("abcdfghijklmnopqrstuvwxyzABCDFGHIJKLMNOPQRSTUVWXYZ1234567890-*!")

func main() {
	
	login := prompData("Enter your login: ")
	// password := prompData("Enter your password")
	url := prompData("Enter your url")

	myAccount := account {
		// password: password,
		url: url,
		login: login,
	}
	var newP int = 12
	myAccount.generatePassword(newP)
	myAccount.outputPassword()
}

func prompData(promp string) string {
	fmt.Print(promp + " ")
	var res string
	fmt.Scan(&res)
	return res
}

