package main

import (
	"fmt"
)

var badWord = []string{"anjing", "babi", "bangsat", "bajingan"}

func main() {

	registeredUser("wahyu", "bajingan")
}

func blacklistWord(word string, keyword []string) bool {
	for _, result := range keyword {
		if result == word {
			return true
		}
	}
	return false
}

func registeredUser(username string, keyword string) {
	if blacklistWord(keyword, badWord) {
		fmt.Println("maaf " + username + " kamu sudah di banned karena menggunakan kata kunci " + keyword)
	} else {
		fmt.Println("hallo selamat datang " + username)

	}

}
