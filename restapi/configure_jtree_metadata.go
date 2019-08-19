package restapi

import (
	"crypto/tls"
	localerrors "errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	// "os/exec"
	"strconv"

	config "github.com/Bio-core/jtree/conf"
	database "github.com/Bio-core/jtree/database"
	"github.com/Bio-core/jtree/dummydata"
	"github.com/Bio-core/jtree/models"
	"github.com/Bio-core/jtree/repos"
	keycloak "github.com/Bio-core/keycloakgo"
	errors "github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	graceful "github.com/tylerb/graceful"

	"github.com/Bio-core/jtree/restapi/operations"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/rs/cors"
)

var c config.Conf

func newID() string {
	// below command does not work since JTree is not using a linux container
	// out, err := exec.Command("uuidgen").Output()
	out, err := uuid.NewV4()
	if err != nil {
		log.Fatal(err)
	}
	ID := fmt.Sprintf("%s", out)
	ID = ID[:len(ID)-1]
	return ID
}

func addPatient(patient *models.Patient) string {
	if patient == nil {
		return "error"
	}

	if patient.PatientID != nil {
		patientOLD := repos.GetPatientByID(*patient.PatientID)
		if patientOLD == nil {
			return "error"
		}
		repos.UpdatePatient(patient)
		return *patient.PatientID
	}
	NewID := newID()
	patient.PatientID = &NewID
	repos.InsertPatient(patient)
	return NewID
}

func addSample(sample *models.Sample) string {
	if sample == nil {
		return "error"
	}

	if sample.SampleID != nil {
		sampleOLD := repos.GetSampleByID(*sample.SampleID)
		if sampleOLD == nil {
			return "error"
		}
		repos.UpdateSample(sample)
		return *sample.SampleID
	}
	NewID := newID()
	sample.SampleID = &NewID
	repos.InsertSample(sample)
	return NewID
}

func addExperiment(experiment *models.Experiment) string {
	if experiment == nil {
		//return errors.New(500, "error")
		return "error"
	}

	if experiment.ExperimentID != nil {
		experimentOLD := repos.GetExperimentByID(*experiment.ExperimentID)
		if experimentOLD == nil {
			//return errors.New(500, "error")
			return "error"
		}
		repos.UpdateExperiment(experiment)
		return *experiment.ExperimentID
	}
	NewID := newID()
	experiment.ExperimentID = &NewID
	repos.InsertExperiment(experiment)
	return NewID
}

func addResult(result *models.Result) string {
	if result == nil {
		return "error"
	}

	if result.ResultsID != nil {
		resultOLD := repos.GetResultByID(*result.ResultsID)
		if resultOLD == nil {
			return "error"
		}
		repos.UpdateResult(result)
		return *result.ResultsID
	}
	NewID := newID()
	result.ResultsID = &NewID
	repos.InsertResult(result)
	return NewID
}

func addResultdetail(resultdetail *models.Resultdetails) string {
	if resultdetail == nil {
		return "error"
	}

	if resultdetail.ResultsDetailsID != nil {
		resultdetailOLD := repos.GetResultDetailByID(*resultdetail.ResultsDetailsID)
		if resultdetailOLD == nil {
			return "error"
		}
		repos.UpdateResultDetail(resultdetail)
		return *resultdetail.ResultsDetailsID
	}
	NewID := newID()
	resultdetail.ResultsDetailsID = &NewID
	repos.InsertResultDetail(resultdetail)
	return NewID
}

func allSamples(query string) (result []*models.Record) {
	if query == "search" || query == "" {
		query = "SELECT * FROM Samples"
	}
	list := repos.GetAllSamples(query)
	result = make([]*models.Record, 0)
	for _, item := range list {
		result = append(result, item)
	}
	return
}

func getSamplesByQuery(query *models.Query) []*models.Record {
	queryString := database.BuildQuery(*query)
	return allSamples(queryString)
}

func getColumns() [][]string {
	columns := database.GetColumns(database.GetTables())
	columnArray := make([][]string, len(columns))
	for i, column := range columns {
		columnArray[i] = make([]string, 2)
		columnArray[i][0] = column
		columnArray[i][1] = database.Map[column]
	}
	return columnArray
}

func getSearchable() []string {
	return models.Sefields.Searchable
}

func getUneditable() []string {
	return models.Sefields.Uneditable
}

func logout() bool {
	return true
}

func upload(file operations.PostUploadParams) error {
	if _, err := os.Stat("./uploads/" + file.Filename); !os.IsNotExist(err) {
		return localerrors.New("File already exists")
	}
	f, err := os.OpenFile("./uploads/"+file.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = io.Copy(f, file.Upfile)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

var databaseFlags = struct {
	Host       string `long:"databaseHost" description:"Database Host" required:"false"`
	Name       string `long:"databaseName" description:"Database Name" required:"false"`
	SelectUser string `long:"dbUsernameSelect" description:"Database Username for Select" required:"false"`
	SelectPass string `long:"dbPasswordSelect" description:"Database Password for Select" required:"false"`
	UpdateUser string `long:"dbUsernameUpdate" description:"Database Username for Update" required:"false"`
	UpdatePass string `long:"dbPasswordUpdate" description:"Database Password for Update" required:"false"`
}{}
var keycloakFlags = struct {
	Active bool   `short:"s" description:"Use Security Bool" required:"false"`
	Host   string `long:"keycloakHost" description:"Keycloak Host" required:"false"`
}{}
var dataGenFlags = struct {
	Generate int `short:"g" description:"generate data" required:"false"`
}{}

func configureFlags(api *operations.JtreeMetadataAPI) {
	api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{
		swag.CommandLineOptionsGroup{
			ShortDescription: "Database Flags",
			LongDescription:  "",
			Options:          &databaseFlags,
		},
		swag.CommandLineOptionsGroup{
			ShortDescription: "Keycloak Flags",
			LongDescription:  "",
			Options:          &keycloakFlags,
		},
		swag.CommandLineOptionsGroup{
			ShortDescription: "Data Generation Flags",
			LongDescription:  "",
			Options:          &dataGenFlags,
		},
	}
}

func configureAPI(api *operations.JtreeMetadataAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError
	c.GetConf()
	setupOptions()
	models.Enums = models.GetEnums(models.Enums)
	models.Sefields = &models.SEFields{}
	models.Sefields = models.Sefields.GetSEFields()

	database.Map = database.MapSuper()

	//database.DBSelect = database.Init(c.Database.Host, c.Database.Selectuser+":"+c.Database.Selectpass+"@/"+c.Database.Name+"?parseTime=true", database.DBSelect)
	//database.DBUpdate = database.Init(c.Database.Host, c.Database.Updateuser+":"+c.Database.Updatepass+"@/"+c.Database.Name+"?parseTime=true", database.DBUpdate)
	// database.DBSelect = database.Init(c.Database.Host, c.Database.Selectuser+":"+c.Database.Selectpass+"@tcp(172.23.0.2:3306)/"+c.Database.Name+"?parseTime=true", database.DBSelect)
	// database.DBUpdate = database.Init(c.Database.Host, c.Database.Updateuser+":"+c.Database.Updatepass+"@tcp(172.23.0.2:3306)/"+c.Database.Name+"?parseTime=true", database.DBUpdate)
	
	database.DBSelect = database.Init(c.Database.Host, c.Database.Selectuser+":"+c.Database.Selectpass+"@tcp(mysql:3306)/"+c.Database.Name+"?parseTime=true", database.DBSelect)
	database.DBUpdate = database.Init(c.Database.Host, c.Database.Updateuser+":"+c.Database.Updatepass+"@tcp(mysql:3306)/"+c.Database.Name+"?parseTime=true", database.DBUpdate)
	ServerName := c.App.Host + ":" + strconv.Itoa(c.App.Port)
	KeycloakserverName := c.Keycloak.Host

	if keycloakFlags.Active {
		keycloak.Init(KeycloakserverName, "http://"+ServerName, "/Jtree/metadata/0.1.0/columns", "/Jtree/metadata/0.1.0/logout")
	}
	if dataGenFlags.Generate != 0 {
		dummydata.MakeData(dataGenFlags.Generate)
	}
	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})

	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()
	api.PostUploadHandler = operations.PostUploadHandlerFunc(func(params operations.PostUploadParams) middleware.Responder {
		if err := upload(params); err != nil {
			return operations.NewPostUploadConflict()
		}
		return operations.NewPostUploadOK().WithPayload(true)
	})
	api.AddExperimentHandler = operations.AddExperimentHandlerFunc(func(params operations.AddExperimentParams) middleware.Responder {
		return operations.NewAddExperimentOK().WithPayload(addExperiment(params.Experiment))
	})
	api.AddPatientHandler = operations.AddPatientHandlerFunc(func(params operations.AddPatientParams) middleware.Responder {
		return operations.NewAddPatientOK().WithPayload(addPatient(params.Patient))
	})
	api.AddSampleHandler = operations.AddSampleHandlerFunc(func(params operations.AddSampleParams) middleware.Responder {
		return operations.NewAddSampleOK().WithPayload(addSample(params.Sample))
	})
	api.AddResultHandler = operations.AddResultHandlerFunc(func(params operations.AddResultParams) middleware.Responder {
		return operations.NewAddResultOK().WithPayload(addResult(params.Result))
	})
	api.AddResultdetailsHandler = operations.AddResultdetailsHandlerFunc(func(params operations.AddResultdetailsParams) middleware.Responder {
		return operations.NewAddResultdetailsOK().WithPayload(addResultdetail(params.Resultdetails))
	})
	api.GetSamplesByQueryHandler = operations.GetSamplesByQueryHandlerFunc(func(params operations.GetSamplesByQueryParams) middleware.Responder {
		return operations.NewGetSamplesByQueryOK().WithPayload(getSamplesByQuery(params.Query))
	})
	api.LogoutHandler = operations.LogoutHandlerFunc(func(params operations.LogoutParams) middleware.Responder {
		return operations.NewLogoutOK().WithPayload(logout())
	})
	api.GetSampleColumnsHandler = operations.GetSampleColumnsHandlerFunc(func(params operations.GetSampleColumnsParams) middleware.Responder {
		return operations.NewGetSampleColumnsOK().WithPayload(getColumns())
	})
	api.GetSearchableHandler = operations.GetSearchableHandlerFunc(func(params operations.GetSearchableParams) middleware.Responder {
		return operations.NewGetSearchableOK().WithPayload(getSearchable())
	})
	api.GetUneditableHandler = operations.GetUneditableHandlerFunc(func(params operations.GetUneditableParams) middleware.Responder {
		return operations.NewGetUneditableOK().WithPayload(getUneditable())
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *graceful.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	x := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT"},
		AllowedHeaders:   []string{"*"},
	})
	handler = x.Handler(handler)
	if keycloakFlags.Active {
		return keycloak.AuthMiddlewareHandler(handler)
	}
	return handler
}

func setupOptions() {
	if databaseFlags.Host != "" {
		c.Database.Host = databaseFlags.Host
	}
	if databaseFlags.Name != "" {
		c.Database.Name = databaseFlags.Name
	}
	if databaseFlags.SelectUser != "" {
		c.Database.Selectuser = databaseFlags.SelectUser
	}
	if databaseFlags.SelectPass != "" {
		c.Database.Selectpass = databaseFlags.SelectPass
	}
	if databaseFlags.UpdateUser != "" {
		c.Database.Updateuser = databaseFlags.UpdateUser
	}
	if databaseFlags.UpdatePass != "" {
		c.Database.Updatepass = databaseFlags.UpdatePass
	}
	if keycloakFlags.Host != "" {
		c.Keycloak.Host = keycloakFlags.Host
	}
}
