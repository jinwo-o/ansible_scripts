package repos

import (
	"fmt"
	"log"

	database "github.com/Bio-core/jtree/database"
	"github.com/Bio-core/jtree/models"
)

const shortForm = "2006-01-02"

//InsertPatient allows users to add generic objects to a collection in the database
func InsertPatient(person *models.Patient) bool {
	stmt, err := database.DBUpdate.Prepare("INSERT INTO `patients`(`first_name`,`last_name`,`initials`,`gender`,`mrn`,`dob`,`on_hcn`,`clinical_history`,`patient_type`,`se_num`,`patient_id`,`date_received`,`referring_physician`,`date_reported`,`surgical_date`)VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);")
	if err != nil {
		log.Fatal(err)
	}
	result, err := stmt.Exec(
		person.FirstName,
		person.LastName,
		person.Initials,
		person.Gender,
		person.Mrn,
		person.Dob.Format(shortForm),
		person.OnHcn,
		person.ClinicalHistory,
		person.PatientType,
		person.SeNum,
		person.PatientID,
		person.DateReceived.Format(shortForm),
		//FIXTHIS
		person.ReferringPhysician,
		person.DateReported.Format(shortForm),
		person.SurgicalDate.Format(shortForm))
	stmt.Close()
	if err != nil {
		log.Fatal(err, result)
	}
	return true
}

//UpdatePatient remove a patient by id
func UpdatePatient(person *models.Patient) bool {
	stmt, err := database.DBUpdate.Prepare("UPDATE `patients` SET`first_name` = ?,`last_name` = ?,`initials` = ?,`gender` = ?,`mrn` = ?,`dob` = ?,`on_hcn` = ?,`clinical_history` = ?,`patient_type` = ?,`se_num` = ?,`patient_id` = ?,`date_received` = ?,`referring_physician` = ?,`date_reported` = ?,`surgical_date` = ? WHERE `patient_id` = ?;")
	if err != nil {
		log.Fatal(err)
		return false
	}
	result, err := stmt.Exec(
		person.FirstName,
		person.LastName,
		person.Initials,
		person.Gender,
		person.Mrn,
		person.Dob.Format(shortForm),
		person.OnHcn,
		person.ClinicalHistory,
		person.PatientType,
		person.SeNum,
		person.PatientID,
		person.DateReceived.Format(shortForm),
		//FIXTHIS
		person.ReferringPhysician,
		person.DateReported.Format(shortForm),
		person.SurgicalDate.Format(shortForm),
		person.PatientID)
	stmt.Close()
	if err != nil {
		log.Fatal(err, result)
		return false
	}
	return true
}

//GetPatientByID gets all and results a list of samples
func GetPatientByID(ID string) *models.Patient {
	patients := []*models.Patient{}
	query := models.Query{}
	query.SelectedFields = make([]string, 0)
	query.SelectedFields = append(query.SelectedFields, "*")
	query.SelectedTables = make([]string, 0)
	query.SelectedTables = append(query.SelectedTables, "patients")
	query.SelectedCondition = make([][]string, 0)
	//query.SelectedCondition = append(query.SelectedCondition, make([]string, 0))
	conditions := []string{"AND", "patients.patient_id", "Equal to", ID}
	query.SelectedCondition = append(query.SelectedCondition, conditions)

	queryString := database.BuildQuery(query)
	err := database.DBSelect.Select(&patients, queryString)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if len(patients) == 0 {
		return nil
	}
	return patients[0]
}
