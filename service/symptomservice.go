package service

import (
	"Website_1/util"
	"bufio"
	"fmt"
	"net/http"
)

type Symptoms struct {
	Id              string
	Symptom         string
	Context         string
	SymptomDuration string
	Date            string
	Note            string
	Severity        string
	FileName        string
	FileData        []byte
}

func AddSymptom(w http.ResponseWriter, r *http.Request) {
	user, err := GetUserSession(r)
	if err != nil || !user.Authenticated {
		http.Redirect(w, r, "/logout", 301)
	}
	err = util.Tpl.ExecuteTemplate(w, "addSymptom.gohtml", user)
	fmt.Println(err)
}

func AddSymptomProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		//r.ParseMultipartForm(10 << 20)

		file, handler, err := r.FormFile("file")
		if err != nil {
			fmt.Println("Error Retrieving the File: ", err)
			return
		}
		defer file.Close()

		var size int64 = handler.Size
		inputFile := make([]byte, size)
		buffer := bufio.NewReader(file)
		_, err = buffer.Read(inputFile)

		symptomsData := Symptoms{
			Symptom:         r.FormValue("symptom"),
			Context:         r.FormValue("context"),
			SymptomDuration: r.FormValue("duration"),
			Date:            r.FormValue("date"),
			Note:            r.FormValue("note"),
			Severity:        r.FormValue("severity"),
			FileName:        handler.Filename,
			FileData:        inputFile,
		}

		_, err = util.DB.Exec("INSERT INTO Symptom(Symptom, Context, Date, symptom_duration, Severity, Note, file_name, file_data, user_id) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9)",
			symptomsData.Symptom, symptomsData.Context, symptomsData.Date, symptomsData.SymptomDuration, symptomsData.Severity, symptomsData.Note, symptomsData.FileName, symptomsData.FileData, r.FormValue("user_id"))
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Server error, unable to create your account.", 500)
			return
		}

	}
	http.Redirect(w, r, "/", 301)
	w.Write([]byte("Done uploaded to your activity"))
}
