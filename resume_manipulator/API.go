package resume_manipulator

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const RESUME_FILE_NAME string = "resume.txt"

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}
}

func GenerateResume() {

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
	data := message["response"][0]

	file, err := os.Create(RESUME_FILE_NAME)
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}

	resume := fmt.Sprintf("<html>"+
		"<body>"+
		"<h1>Резюме</h1>"+
		"<b>Имя: </b>%s<br>"+
		"<b>Фамилия: </b>%s<br>"+
		"<b>Университет: </b>%s<br>"+
		"<b>Факультет: </b>%s<br>"+
		"<b>Дата рождения: </b>%s<br>"+
		"</body>"+
		"</html>", data["first_name"], data["last_name"], data["university_name"], data["faculty_name"], data["bdate"])

	_, err = file.Write([]byte(resume))
	if err != nil {
		log.Fatal(err)
	}
}

func GetResume() []byte {
	data, err := os.ReadFile(RESUME_FILE_NAME)
	if err != nil {
		log.Fatal(err)
	}
	return data
}
