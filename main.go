package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
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

func newAccount(login, password, urlString string) (*account, error) {
	if login == "" {
		return nil, errors.New("INVALID LOGIN")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID URL")
	}
	newAcc := &account{
		login: login,
		password: password,
		url: urlString,
	}

	if password == "" {
		var newP int = 12
		newAcc.generatePassword(newP)
	}
	return newAcc, nil
}

var letterRunes = []rune("abcdfghijklmnopqrstuvwxyzABCDFGHIJKLMNOPQRSTUVWXYZ1234567890-*!")

func main() {
	
	login := prompData("Enter your login: ")
	password := prompData("Enter your password")
	url := prompData("Enter your url")

	myAccount, err := newAccount(login, password, url)
	if err != nil {
		fmt.Println("Invalid url or login")
		return
	}
	myAccount.outputPassword()
}

func prompData(promp string) string {
	fmt.Print(promp + " ")
	var res string
	fmt.Scanln(&res)
	return res
}

