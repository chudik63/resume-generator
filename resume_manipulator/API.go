package resume_manipulator

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

const ResumeFileName string = "resume.txt"

func GenerateResume() {
	message := map[string][]map[string]interface{}{
		"response": {},
	}
	resp := get()
	err := json.Unmarshal(resp, &message)
	if err != nil {
		log.Fatal("Json reading error")
	}

	data := message["response"][0]

	file, err := os.Create(ResumeFileName)
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}

	resume := []byte(fmt.Sprintf("<html>"+
		"<body>"+
		"<h1>Резюме</h1>"+
		"<b>Имя: </b>%s<br>"+
		"<b>Фамилия: </b>%s<br>"+
		"<b>Университет: </b>%s<br>"+
		"<b>Факультет: </b>%s<br>"+
		"<b>Дата рождения: </b>%s<br>"+
		"</body>"+
		"</html>", data["first_name"], data["last_name"], data["university_name"], data["faculty_name"], data["bdate"]))

	_, err = file.Write(resume)
	if err != nil {
		log.Fatal(err)
	}
}

func GetResume() []byte {
	data, err := os.ReadFile(ResumeFileName)
	if err != nil {
		log.Fatal(err)
	}
	return data
}
