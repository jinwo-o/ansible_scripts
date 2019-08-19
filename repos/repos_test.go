package repos

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

var host = "http://127.0.0.1:8000"

func TestUpdatePatientPOST(t *testing.T) {

	patient := GetPatientByID("1")
	first := "Mitchell"
	last := "Strong"
	patient.FirstName = &first
	patient.LastName = &last
	person1Bytes, err := json.Marshal(patient)

	if err != nil {
		t.Fail()
		return
	}

	body := bytes.NewReader(person1Bytes)

	req, err := http.NewRequest("POST", host+"/Jtree/metadata/0.1.0/patient", body)

	if err != nil {
		t.Fail()
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)

	if resp.Status != "201 Created" {
		t.Fail()
		return
	}

	if err != nil {
		t.Fail()
		return
	}

	defer resp.Body.Close()

	patientNew := GetPatientByID("1")

	if *patientNew.FirstName != first || *patientNew.LastName != last {
		t.Fail()
		return
	}

	return
}
