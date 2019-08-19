package jin

import (
	// "bytes"
	// "encoding/json"
	// "fmt"
	// "io/ioutil"
	// "net/http"
	// "github.com/Bio-core/jtree/dummydata"

	// "log"
	"os"
	"testing"

	// "github.com/Bio-core/jtree/restapi/operations"
	// "github.com/go-openapi/loads"
	// "log"
	// "github.com/Bio-core/jtree/restapi"
	// "github.com/Bio-core/jtree/restapi/operations"
	// "github.com/go-openapi/loads"
	// "github.com/Bio-core/jtree/restapi/operations"
	// "github.com/go-openapi/loads"
	// "github.com/Bio-core/jtree/database"
	// "github.com/Bio-core/jtree/models"
	// "github.com/Bio-core/jtree/repos"
)

var host = "http://127.0.0.1:8000"

func TestMain(m *testing.M) {
	testResults := m.Run()
	os.Exit(testResults)
}

// func TestSetupSever(t *testing.T) {
// 	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
// 	if err != nil {
// 		t.Errorf("%v", err)
// 		log.Fatalln(err)
// 	}

// 	api := operations.NewJtreeMetadataAPI(swaggerSpec)
// 	server := restapi.NewServer(api)

// 	server.ConfigureAPI()

// 	go server.Serve()
// }

func TestUrls(t *testing.T) {
	result := true
	result = result && CheckPageResponse(host+"/Jtree/metadata/0.1.0/columns")
	result = result && CheckPageResponse(host+"/Jtree/metadata/0.1.0/uneditable")
	result = result && CheckPageResponse(host+"/Jtree/metadata/0.1.0/searchable")
	result = result && !CheckPageResponse(host+"/x")
	result = result && !CheckNoPageResponse(host+"/Jtree/metadata/0.1.0/searchable")
	result = result && CheckNoPageResponse(host+"/x")

	if result != true {
		t.Error("Web Pages Not Successful")
	}
}

// func TestGenerateDummyData(t *testing.T) {
// 	dummydata.MakeData(100)

// 	query := models.Query{}
// 	query.SelectedFields = make([]string, 0)
// 	query.SelectedFields = append(query.SelectedFields, "*")
// 	query.SelectedTables = make([]string, 0)
// 	query.SelectedTables = append(query.SelectedTables, "patients")
// 	query.SelectedCondition = make([][]string, 0)
// 	querystring := database.BuildQuery(query)
// 	if len(repos.GetAllSamples(querystring)) != 100 {
// 		t.Fail()
// 	}
// 	query.SelectedTables[0] = "samples"
// 	querystring = database.BuildQuery(query)
// 	if len(repos.GetAllSamples(querystring)) != 287 {
// 		t.Fail()
// 	}
// 	query.SelectedTables[0] = "experiments"
// 	querystring = database.BuildQuery(query)
// 	if len(repos.GetAllSamples(querystring)) != 866 {
// 		t.Fail()
// 	}
// 	query.SelectedTables[0] = "results"
// 	querystring = database.BuildQuery(query)
// 	if len(repos.GetAllSamples(querystring)) != 1282 {
// 		t.Fail()
// 	}
// 	query.SelectedTables[0] = "resultdetails"
// 	querystring = database.BuildQuery(query)
// 	if len(repos.GetAllSamples(querystring)) != 1899 {
// 		t.Fail()
// 	}
// 	return
// }

// func TestAddPatientPOST(t *testing.T) {

// 	dummydata.MakeData(100)
// 	fmt.Println("Adding Patient Post")
// 	patient := dummydata.MakePatient(-1)
// 	person1Bytes, err := json.Marshal(patient)

// 	if err != nil {
// 		t.Fail()
// 		return
// 	}

// 	body := bytes.NewReader(person1Bytes)
// 	fmt.Print(body)

// 	req, err := http.NewRequest("POST", host+"/Jtree/metadata/0.1.0/patient", body)

// 	if err != nil {
// 		t.Fail()
// 		return
// 	}

// 	req.Header.Set("Content-Type", "application/json")

// 	resp, err := http.DefaultClient.Do(req)

// 	content, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		t.Fail()
// 		return
// 	}
// 	if resp.Status != "200 OK" && string(content) != "error" {
// 		t.Fail()
// 		return
// 	}

// 	if err != nil {
// 		t.Fail()
// 		return
// 	}

// 	defer resp.Body.Close()

// }
