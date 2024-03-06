package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	message := map[string][]map[string]interface{}{
		"response": {},
	}
	method, exists := os.LookupEnv("method")
	if !exists {
		log.Fatal("No method found")
	}

	id, exists := os.LookupEnv("vk_id")
	if !exists {
		log.Fatal("No vk_id found")
	}
	token, exists := os.LookupEnv("token")
	if !exists {
		log.Fatal("No token found")
	}
	resp, _ := http.Get(fmt.Sprintf("https://api.vk.com/method/%s?user_ids=%s&fields=bdate,education&access_token=%s&v=5.131", method, id, token))
	body, _ := ioutil.ReadAll(resp.Body)
	err := json.Unmarshal(body, &message)
	if err != nil {
		log.Fatal("Json reading error")
	}
	fmt.Print(message["response"][0]["faculty_name"])
}
