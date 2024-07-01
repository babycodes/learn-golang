package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// Struct dengan json struct tag
type FactResponse struct {
	Teks string `json:"text"`
	Tipe string `json:"type"`
}

func main() {
	// 1. buat request
	req, err := http.NewRequest("GET", "https://cat-fact.herokuapp.com/facts/random", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// 2. buat vlient
	client := http.Client{}

	// 3. panggil request dengan client
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// tutup response body
	defer res.Body.Close()

	// 4. baca response body
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	//  5. convert ke tipe data FactResponse
	var factResponse FactResponse
	err = json.Unmarshal(resBody, &factResponse)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fmt.Println("teks", factResponse.Teks)
	fmt.Println("tipe", factResponse.Tipe)
}
