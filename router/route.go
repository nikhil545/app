package router

import (
	"Website_1/service"
	"Website_1/util"
	"fmt"
	"net/http"
)

func HttpEndpoint() {
	http.Handle("/public/CSS/", http.StripPrefix("/public/CSS", http.FileServer(http.Dir("./public/CSS"))))
	http.Handle("/public/img/", http.StripPrefix("/public/img", http.FileServer(http.Dir("./public/img"))))
	http.HandleFunc("/", dashboard)
	http.HandleFunc("/index", index)
	http.HandleFunc("/addAllergy", service.AddAllergy)
	http.HandleFunc("/addAllergy/process", service.AddAllergyProcess)
	http.HandleFunc("/allDoc", service.AllDocs)
	http.HandleFunc("/addHospital", service.AddHospital)
	http.HandleFunc("/addHospital/process", service.AddHospitalProcess)
	http.HandleFunc("/addMedication", service.AddMedication)
	http.HandleFunc("/addMedication/process", service.AddMedicationProcess)
	http.HandleFunc("/addSymptom", service.AddSymptom)
	http.HandleFunc("/addSymptom/process", service.AddSymptomProcess)
	//http.HandleFunc("/symptom/download", service.DownloadImmuFile)
	http.HandleFunc("/addImmunity", addImmunity)
	http.HandleFunc("/addImmunity/process", service.AddImmuprocess)
	http.HandleFunc("/immu/download", service.DownloadImmuFile)
	http.HandleFunc("/contactus", contactus)
	http.HandleFunc("/contactus/process", service.MessageBrokerService)
	http.HandleFunc("/signup", service.Signup)
	http.HandleFunc("/signup/process", service.Signupprocess)
	http.HandleFunc("/login/process", service.Loginprocess)
	http.HandleFunc("/login", service.Login)
	http.HandleFunc("/logout", service.Logout)
	http.HandleFunc("/tutorial", tutorial)
	http.ListenAndServe(":5050", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	err := util.Tpl.ExecuteTemplate(w, "index.gohtml", nil)
	fmt.Println(err)
}

func dashboard(w http.ResponseWriter, r *http.Request) {
	user, err := service.GetUserSession(r)
	if err != nil || !user.Authenticated {
		http.Redirect(w, r, "/login", 301)
	}

	err = util.Tpl.ExecuteTemplate(w, "dashboard.gohtml", user)
	fmt.Println(err)
}

func contactus(w http.ResponseWriter, r *http.Request) {
	err := util.Tpl.ExecuteTemplate(w, "contactus.gohtml", nil)
	fmt.Println(err)
}

func addImmunity(w http.ResponseWriter, r *http.Request) {
	user, err := service.GetUserSession(r)
	if err != nil || !user.Authenticated {
		http.Redirect(w, r, "/login", 301)
	}
	err = util.Tpl.ExecuteTemplate(w, "addImmu.gohtml", user)
	fmt.Println(err)
}

func tutorial(w http.ResponseWriter, r *http.Request) {
	err := util.Tpl.ExecuteTemplate(w, "tutorial.gohtml", nil)
	fmt.Println(err)
}
