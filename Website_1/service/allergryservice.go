package service

import (
	"Website_1/util"
	"bufio"
	"fmt"
	"net/http"
)

type Allegry struct {
	Id       string
	Allergen string
	Reaction string
	Severity string
	Date     string
	Note     string
	FileName string
	FileData []byte
}

func AddAllergy(w http.ResponseWriter, r *http.Request) {
	user, err := GetUserSession(r)
	if err != nil || !user.Authenticated {
		http.Redirect(w, r, "/logout", 301)
	}

	err = util.Tpl.ExecuteTemplate(w, "addAllergy.gohtml", user)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/", 301)
	}
}

func AddAllergyProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

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

		allergenData := Allegry{
			Allergen: r.FormValue("allergen"),
			Reaction: r.FormValue("reactions"),
			Severity: r.FormValue("severity"),
			Date:     r.FormValue("dateidenty"),
			Note:     r.FormValue("note"),
			FileName: handler.Filename,
			FileData: inputFile,
		}

		_, err = util.DB.Exec("INSERT INTO Allegry(Allergen, Reaction, Severity, Date, Note, file_name, file_data,user_id) VALUES($1,$2,$3,$4,$5,$6, $7, $8)",
			allergenData.Allergen, allergenData.Reaction, allergenData.Severity, allergenData.Date, allergenData.Note, allergenData.FileName, allergenData.FileData, r.FormValue("user_id"))
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Server error, unable to create allergen data.", 500)
			return
		}

	}
	http.Redirect(w, r, "/", 301)
	w.Write([]byte("Done uploaded to your activity"))
}
