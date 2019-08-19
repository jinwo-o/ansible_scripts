package database

import (

	// SQL driver for mysql

	//Need this for the other one below
	_ "github.com/go-sql-driver/mysql"
	sqlx "github.com/jmoiron/sqlx"
)

var databaseName string
var connectionString string
var joinMap map[string]map[string]string

var err error

//DBSelect is the database reference for selects
var DBSelect *sqlx.DB

//DBUpdate is the database reference for updates
var DBUpdate *sqlx.DB

//Init creates a connection to the database
func Init(dbName, connectionstring string, DB *sqlx.DB) *sqlx.DB {
	databaseName = dbName
	connectionString = connectionstring
	DB, err = sqlx.Connect(databaseName, connectionString)
	if err != nil {
		panic(err)
	}
	DB.SetMaxOpenConns(100)
	makeJoinMap()
	return DB
}

func makeJoinMap() {
	joinMap = make(map[string]map[string]string, 5)
	//Patient
	patientMap := make(map[string]string, 1)
	patientMap["samples"] = "patient_id"
	joinMap["patients"] = patientMap
	//Sample
	sampleMap := make(map[string]string, 2)
	sampleMap["experiments"] = "sample_id"
	sampleMap["patients"] = "patient_id"
	joinMap["samples"] = sampleMap
	//Experiment
	experimentMap := make(map[string]string, 2)
	experimentMap["samples"] = "sample_id"
	experimentMap["results"] = "experiment_id"
	joinMap["experiments"] = experimentMap
	//Results
	resultMap := make(map[string]string, 2)
	resultMap["experiments"] = "experiment_id"
	resultMap["resultdetails"] = "results_id"
	joinMap["results"] = resultMap
	//ResultDetails
	resultdetailsMap := make(map[string]string, 1)
	resultdetailsMap["results"] = "results_id"
	joinMap["resultdetails"] = resultdetailsMap
}
