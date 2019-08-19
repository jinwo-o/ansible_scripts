package jin

import (
	"net/http"
	// Does not get imported in TRAVIS FIX THIS
	// "github.com/Bio-core/jtree/models"
)

//CheckPageResponse checks if a page that should respond is found correctly
func CheckPageResponse(url string) bool {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false
	}
	response, err := client.Do(req)
	if err != nil {
		return false
	}
	if response == nil {
		return false
	}
	if response.Status == "404 Not Found" {
		return false
	}
	return true
}

//CheckNoPageResponse checks if a page that does not exist responds with a 404 Error
func CheckNoPageResponse(url string) bool {
	response, err := http.Get(url)
	if err != nil {
		return true
	}
	if response == nil {
		return true
	}
	if response.Status == "404 Not Found" {
		return true
	}
	return false
}

// func returnQuery(fields, tables []string, conditions [][]string) models.Query {
// 	query := models.Query{
// 		SelectedFields:    fields,
// 		SelectedTables:    tables,
// 		SelectedCondition: conditions,
// 	}
// 	return query
// }
