package main

import (
	"fmt"
	"golang/database"
	"golang/helper"
	_ "golang/internal"
)

var badWord = []string{"anjing", "babi", "bangsat", "bajingan"}

func main() {

	registeredUser("wahyu", "bajingan")
	fmt.Println(database.GetDatabase())
	// internal.init()
}

func registeredUser(username string, keyword string) {
	if helper.BacklistWord(keyword, badWord) {
		fmt.Println("maaf " + username + " kamu sudah di banned karena menggunakan kata kunci " + keyword)
	} else {
		fmt.Println("hallo selamat datang " + username)

	}

}
