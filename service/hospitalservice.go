package service

import (
	"Website_1/util"
	"bufio"
	"fmt"
	"net/http"
)

type Hospitals struct {
	Id             string
	Hospital       string
	Reason         string
	AdmissionDate  string
	DischargedDate string
	Note           string
	FileName       string
	FileData       []byte
}

func AddHospital(w http.ResponseWriter, r *http.Request) {
	user, err := GetUserSession(r)
	if err != nil || !user.Authenticated {
		http.Redirect(w, r, "/logout", 301)
	}
	err = util.Tpl.ExecuteTemplate(w, "addHospital.gohtml", user)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/", 301)
	}
}

func AddHospitalProcess(w http.ResponseWriter, r *http.Request) {
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

		hospitalData := Hospitals{
			Hospital:       r.FormValue("hospital"),
			Reason:         r.FormValue("reason"),
			AdmissionDate:  r.FormValue("dateadmis"),
			DischargedDate: r.FormValue("datedischarge"),
			Note:           r.FormValue("note"),
			FileName:       handler.Filename,
			FileData:       inputFile,
		}

		_, err = util.DB.Exec("INSERT INTO Hospital(Hospital, Reason, AdmissionDate, DischargeDate, Note, file_name, file_data,user_id) VALUES($1,$2,$3,$4,$5,$6, $7, $8)",
			hospitalData.Hospital, hospitalData.Reason, hospitalData.AdmissionDate, hospitalData.DischargedDate, hospitalData.Note, hospitalData.FileName, hospitalData.FileData, r.FormValue("user_id"))
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Server error, unable to create your account.", 500)
			return
		}

	}
	http.Redirect(w, r, "/", 301)
	w.Write([]byte("Done uploaded to your activity"))
}
