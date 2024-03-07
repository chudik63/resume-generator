package resume_manipulator

import (
	"encoding/json"
	"html/template"
	"log"
	"os"
	"time"
)

const TemplateFileName string = "templates/index.html"
const PublicFileName string = "public/index.html"

type Career struct {
	Company  string
	From     float64
	Until    float64
	Position string
}

type Content struct {
	FirstName  string
	LastName   string
	City       string
	University string
	Faculty    string
	Graduation float64
	Experience Career
	BirthDate  string
	DateTime   string
}

func GenerateResume() {
	message := map[string][]map[string]interface{}{
		"response": {},
	}
	resp := get()
	err := json.Unmarshal(resp, &message)
	if err != nil {
		log.Fatal(err)
	}

	data := message["response"][0]
	career := data["career"].([]interface{})[0].(map[string]interface{})
	info := Content{
		FirstName:  data["first_name"].(string),
		LastName:   data["last_name"].(string),
		City:       data["city"].(map[string]interface{})["title"].(string),
		University: data["university_name"].(string),
		Faculty:    data["faculty_name"].(string),
		Graduation: data["graduation"].(float64),
		Experience: Career{Company: career["company"].(string), From: career["from"].(float64), Until: career["until"].(float64), Position: career["position"].(string)},
		BirthDate:  data["bdate"].(string),
		DateTime:   time.Now().Format("02.01.2006 15:04"),
	}
	tmpl, _ := template.ParseFiles(TemplateFileName)
	publicFile, err := os.OpenFile(PublicFileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(publicFile, &info)
	if err != nil {
		log.Fatal(err)
	}
}
