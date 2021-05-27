package service

import (
	"Website_1/util"
	"fmt"
	"net/http"
)

type AllDocument struct {
	Hospitals   []Hospitals
	Immus       []Immu
	Medications []Medications
	Symptoms    []Symptoms
	Allegry     []Allegry
}

func AllDocs(w http.ResponseWriter, r *http.Request) {
	user, err := GetUserSession(r)
	if err != nil || !user.Authenticated {
		http.Redirect(w, r, "/dashboard", 301)
	}

	hospitalData, err := GetHospitalData("SELECT id, hospital, reason, admissiondate, dischargedate, note, file_name from hospital where user_id = $1", user.Id)
	if err != nil {
		fmt.Println("hospital data fetching error :", err)
	}

	immuData, err := GetImmuData("SELECT id, vaccine, protection, date, note, file_name from immu where user_id = $1", user.Id)
	if err != nil {
		fmt.Println("vaccine data fetching error :", err)
	}

	medicationData, err := GetMedicationData("SELECT id, medication, doseinfo, reason, prescribedate, finishdate, note, file_name from medication where user_id = $1", user.Id)
	if err != nil {
		fmt.Println("medication data fetching error :", err)
	}

	symptomData, err := GetSymptomData("SELECT id, symptom, severity, context, symptom_duration, date, note, file_name from symptom where user_id = $1", user.Id)
	if err != nil {
		fmt.Println("symptom data fetching error :", err)
	}

	allegryData, err := GetAllegryData("SELECT id, allergen, reaction, severity, date, note, file_name from allegry where user_id = $1", user.Id)
	if err != nil {
		fmt.Println("allergen data fetching error :", err)
	}

	allDocsData := AllDocument{
		Hospitals:   hospitalData,
		Immus:       immuData,
		Medications: medicationData,
		Symptoms:    symptomData,
		Allegry:     allegryData,
	}
	if err := util.Tpl.ExecuteTemplate(w, "alldoc.gohtml", allDocsData); err != nil {
		http.Redirect(w, r, "/dashboard", 301)
	}
}

func GetHospitalData(query string, userId string) ([]Hospitals, error) {
	var hospitalDatas []Hospitals

	results, err := util.DB.Query(query, userId)
	if err != nil {
		return nil, err
	}

	for results.Next() {
		var hospitalData Hospitals
		err = results.Scan(&hospitalData.Id, &hospitalData.Hospital, &hospitalData.Reason, &hospitalData.AdmissionDate, &hospitalData.DischargedDate, &hospitalData.Note, &hospitalData.FileName)
		if err != nil {
			return nil, err
		}
		hospitalDatas = append(hospitalDatas, hospitalData)
	}

	return hospitalDatas, nil
}

func GetImmuData(query string, userId string) ([]Immu, error) {

	var immuDatas []Immu

	results, err := util.DB.Query(query, userId)
	if err != nil {
		return nil, err
	}

	for results.Next() {
		var immuData Immu
		err = results.Scan(&immuData.Id, &immuData.Vaccine, &immuData.Protection, &immuData.Date, &immuData.Note, &immuData.FileName)
		if err != nil {
			return nil, err
		}
		immuDatas = append(immuDatas, immuData)
	}

	return immuDatas, nil
}

func GetMedicationData(query string, userId string) ([]Medications, error) {
	var medicationDatas []Medications
	results, err := util.DB.Query(query, userId)
	if err != nil {
		return nil, err
	}

	for results.Next() {
		var medicationData Medications
		err = results.Scan(&medicationData.Id, &medicationData.Medication, &medicationData.DoseInfo, &medicationData.Reason, &medicationData.PrescribeDate, &medicationData.FinishDate, &medicationData.Note, &medicationData.FileName)
		if err != nil {
			return nil, err
		}
		medicationDatas = append(medicationDatas, medicationData)
	}

	return medicationDatas, nil
}

func GetSymptomData(query string, userId string) ([]Symptoms, error) {
	var symptomDatas []Symptoms

	results, err := util.DB.Query(query, userId)
	if err != nil {
		return nil, err
	}

	for results.Next() {
		var symptomData Symptoms
		err = results.Scan(&symptomData.Id, &symptomData.Symptom, &symptomData.Severity, &symptomData.Context, &symptomData.SymptomDuration, &symptomData.Date, &symptomData.Note, &symptomData.FileName)
		if err != nil {
		}
		symptomDatas = append(symptomDatas, symptomData)
	}

	return symptomDatas, nil
}

func GetAllegryData(query string, userId string) ([]Allegry, error) {
	var allegryDatas []Allegry
	results, err := util.DB.Query(query, userId)
	if err != nil {
		return nil, err
	}

	for results.Next() {
		var allegryData Allegry
		err = results.Scan(&allegryData.Id, &allegryData.Allergen, &allegryData.Reaction, &allegryData.Severity, &allegryData.Date, &allegryData.Note, &allegryData.FileName)
		if err != nil {
		}
		allegryDatas = append(allegryDatas, allegryData)
	}

	return allegryDatas, nil
}
