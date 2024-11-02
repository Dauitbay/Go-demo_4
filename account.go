package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)

type account struct {
	login    string
	password string
	url      string
}

type accountWithTimeStamp struct {
	createdAt time.Time
	updatedAt time.Time
	acc       account
}

func (acc *account) outputPassword() {
	fmt.Println(acc)
	fmt.Println(acc.login, acc.password, acc.url)
}

func (acc *account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.password = string(res)
}

// func newAccount(login, password, urlString string) (*account, error) {
// 	if login == "" {
// 		return nil, errors.New("INVALID LOGIN")
// 	}
// 	_, err := url.ParseRequestURI(urlString)
// 	if err != nil {
// 		return nil, errors.New("INVALID URL")
// 	}
// 	newAcc := &account{
// 		login:    login,
// 		password: password,
// 		url:      urlString,
// 	}

// 	if password == "" {
// 		var newP int = 12
// 		newAcc.generatePassword(newP)
// 	}
// 	return newAcc, nil
// }

func newAccountWithTimeStamp(login, password, urlString string) (*accountWithTimeStamp, error) {
	if login == "" {
		return nil, errors.New("INVALID LOGIN")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID URL")
	}
	newAcc := &accountWithTimeStamp{
		createdAt: time.Now(),
		updatedAt: time.Now(),
		acc: account{
			login:    login,
			password: password,
			url:      urlString,
		},
	}

	if password == "" {
		var newP int = 12
		newAcc.acc.generatePassword(newP)
	}
	return newAcc, nil
}

var letterRunes = []rune("abcdfghijklmnopqrstuvwxyzABCDFGHIJKLMNOPQRSTUVWXYZ1234567890-*!")
