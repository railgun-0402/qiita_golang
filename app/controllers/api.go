package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Article struct {
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}

// APIの実行メソッド
func ApiFetch() {
	resp, err := http.Get("https://qiita.com/api/v2/items")
	if err != nil {
		log.Fatalln(err)
	}
	// Close
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(body)
	}

	var data []Article

	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatalln(err)
	}

	for _, item := range data {
		fmt.Printf("%s %s\n", item.CreatedAt, item.Title)
	}

}
