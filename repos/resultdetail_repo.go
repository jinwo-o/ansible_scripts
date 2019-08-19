package repos

import (
	"fmt"
	"log"

	database "github.com/Bio-core/jtree/database"
	"github.com/Bio-core/jtree/models"
)

//InsertResultDetail allows users to add generic objects to a collection in the database
func InsertResultDetail(result *models.Resultdetails) bool {
	stmt, err := database.DBUpdate.Prepare("INSERT INTO `resultdetails`(`VAF`,`c_nomenclature`,`coverage`,`exon`,`gene`,`p_nomenclature`,`pcr`,`quality_score`,`result`,`results_details_id`,`results_id`,`risk_score`,`uid`)VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?);")
	if err != nil {
		log.Fatal(err)
	}
	outcome, err := stmt.Exec(
		result.VAF,
		result.CNomenclature,
		result.Coverage,
		result.Exon,
		result.Gene,
		result.PNomenclature,
		result.Pcr,
		result.QualityScore,
		result.Result,
		result.ResultsDetailsID,
		result.ResultsID,
		result.RiskScore,
		result.UID)
	stmt.Close()
	if err != nil {
		log.Fatal(err, outcome)
	}
	return true
}

//GetResultDetailByID gets all and results a list of samples
func GetResultDetailByID(ID string) *models.Resultdetails {
	resultdetails := []*models.Resultdetails{}
	query := models.Query{}
	query.SelectedFields = make([]string, 0)
	query.SelectedFields = append(query.SelectedFields, "*")
	query.SelectedTables = make([]string, 0)
	query.SelectedTables = append(query.SelectedTables, "resultdetails")
	query.SelectedCondition = make([][]string, 0)
	//query.SelectedCondition = append(query.SelectedCondition, make([]string, 0))
	conditions := []string{"AND", "resultdetails.results_details_id", "Equal to", ID}
	query.SelectedCondition = append(query.SelectedCondition, conditions)

	queryString := database.BuildQuery(query)
	err := database.DBSelect.Select(&resultdetails, queryString)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if len(resultdetails) == 0 {
		return nil
	}
	return resultdetails[0]
}

//UpdateResultDetail allows users to add generic objects to a collection in the database
func UpdateResultDetail(result *models.Resultdetails) bool {
	stmt, err := database.DBUpdate.Prepare("UPDATE `resultdetails` SET `VAF` = ?,`c_nomenclature` = ?,`coverage` = ?,`exon` = ?,`gene` = ?,`p_nomenclature` = ?,`pcr` = ?,`quality_score` = ?,`result` = ?,`results_details_id` = ?,`results_id` = ?,`risk_score` = ?,`uid` = ? WHERE `results_details_id` = ?;")
	if err != nil {
		log.Fatal(err)
	}
	outcome, err := stmt.Exec(
		result.VAF,
		result.CNomenclature,
		result.Coverage,
		result.Exon,
		result.Gene,
		result.PNomenclature,
		result.Pcr,
		result.QualityScore,
		result.Result,
		result.ResultsDetailsID,
		result.ResultsID,
		result.RiskScore,
		result.UID,
		result.ResultsDetailsID)
	stmt.Close()
	if err != nil {
		log.Fatal(err, outcome)
	}
	return true
}
