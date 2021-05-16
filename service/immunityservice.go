package service

import (
	"Website_1/util"
	"bufio"

	"fmt"

	"net/http"
	"os"
)

type Immu struct {
	Id         string
	Vaccine    string
	Protection string
	Date       string
	Note       string
	FileName   string
	file_data  []byte
}

func AddImmuprocess(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		file, handler, err := r.FormFile("file_name")
		if err != nil {
			fmt.Println("Error Retrieving the File")
			fmt.Println(err)
			return
		}
		defer file.Close()

		var size int64 = handler.Size
		bytes := make([]byte, size)
		buffer := bufio.NewReader(file)
		_, err = buffer.Read(bytes)

		imm := Immu{
			Vaccine:    r.FormValue("vaccine"),
			Protection: r.FormValue("protection"),
			Date:       r.FormValue("date"),
			Note:       r.FormValue("note"),
			FileName:   handler.Filename,
		}

		_, err = util.DB.Exec("INSERT INTO Immu(Vaccine, Protection,Date,Note,file_name,file_data,user_id) VALUES($1,$2,$3,$4,$5,$6,$7)", imm.Vaccine, imm.Protection, imm.Date, imm.Note, imm.FileName, bytes, r.FormValue("user_id"))
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Server error, unable to create your account.", 500)
			return
		}

	}
	http.Redirect(w, r, "/", 301)
	w.Write([]byte("Done uploaded to your activity"))
}

func DownloadImmuFile(w http.ResponseWriter, r *http.Request) {
	data := r.URL.Query()
	id := data["id"][0] //format -> immu/download?id=1 

	var fileName string
	var fileData []byte

	err := util.DB.QueryRow("SELECT file_name, file_data FROM immu where id = $1", id).Scan(&fileName, &fileData)
	if err != nil {
		fmt.Println(err)
		return
	}

	if fileName != "" && fileData != nil {
		tmpfile, err := os.Create(fileName)
		defer tmpfile.Close()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		l, err := tmpfile.Write(fileData)
		if err != nil {
			fmt.Println(err)
			tmpfile.Close()
			return
		}
		fmt.Println(l, "bytes written successfully")
		err = tmpfile.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	http.Redirect(w, r, "/alldocs", 301)
	w.Write([]byte("Done uploaded to your activity"))
}
