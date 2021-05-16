package service

import (
	"Website_1/util"
	"bufio"
	"fmt"
	"net/http"
)

type Medications struct {
	Id            string
	Medication    string
	DoseInfo      string
	Reason        string
	PrescribeDate string
	FinishDate    string
	Note          string
	FileName      string
	FileData      []byte
}

func AddMedication(w http.ResponseWriter, r *http.Request) {
	user, err := GetUserSession(r)
	if err != nil || !user.Authenticated {
		http.Redirect(w, r, "/logout", 301)
	}
	err = util.Tpl.ExecuteTemplate(w, "addMedication.gohtml", user)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/", 301)
	}
}

func AddMedicationProcess(w http.ResponseWriter, r *http.Request) {
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

		medicationData := Medications{
			Medication:    r.FormValue("med"),
			DoseInfo:      r.FormValue("dose"),
			Reason:        r.FormValue("reason"),
			PrescribeDate: r.FormValue("datep"),
			FinishDate:    r.FormValue("datef"),
			Note:          r.FormValue("note"),
			FileName:      handler.Filename,
			FileData:      inputFile,
		}

		_, err = util.DB.Exec("INSERT INTO Medication(Medication, DoseInfo, Reason, PrescribeDate, FinishDate, Note, file_name, file_data, user_id) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9)",
			medicationData.Medication, medicationData.DoseInfo, medicationData.Reason, medicationData.PrescribeDate, medicationData.FinishDate, medicationData.Note, medicationData.FileName, medicationData.FileData, r.FormValue("user_id"))
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Server error, unable to create your account.", 500)
			return
		}

	}
	http.Redirect(w, r, "/", 301)
	w.Write([]byte("Done uploaded to your activity"))
}
