package dummydata

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	models "github.com/Bio-core/jtree/models"
	repos "github.com/Bio-core/jtree/repos"
	yaml "gopkg.in/yaml.v2"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var r *rand.Rand
var id1 int
var id2 int
var id3 int
var id4 int
var id5 int
var genes *GeneArray
var random []RandomPerson

const shortForm = "2006-01-02"

//GeneArray is an object with an array of gene types
type GeneArray struct {
	Genes []string
}

//RandomPerson is a scruct to help make real names
type RandomPerson struct {
	Name    string `json:"name,omitempty"`
	Surname string `json:"surname,omitempty"`
	Gender  string `json:"gender,omitempty"`
	Region  string `json:"region,omitempty"`
}

//MakeData makes dummy data and puts it into the db
func MakeData(number int) {
	r = rand.New(rand.NewSource(99))
	random = getManyRandomPeople(number)
	genes = &GeneArray{}
	genes = genes.GetGenes()
	num1 := createPatients(number)
	num2 := createSamples(num1)
	num3 := createExperiments(num2)
	num4 := createResults(num3)
	createResultDetails(num4)
}

func makeRandomString() string {
	num := rand.Intn(50)
	value := randSeq(num)
	return value
}
func makeRandomName(id int, last bool) string {
	if id < 0 {
		id = rand.Intn(99) + 1
	}
	if last {
		return random[id-1].Surname
	}
	// return random[1].Name
	return random[id-1].Name
}

func makeRandomDate() string {
	year := strconv.Itoa(rand.Intn(118) + 1900)
	monthint := rand.Intn(11) + 1
	month := strconv.Itoa(monthint)
	if monthint < 10 {
		month = "0" + month
	}
	dayint := rand.Intn(27) + 1
	day := strconv.Itoa(dayint)
	if dayint < 10 {
		day = "0" + day
	}
	date := year + "-" + month + "-" + day
	return date
}

func makeRandomFloat() float32 {
	num := rand.Float32()
	num += float32(rand.Intn(600))
	return num
}

func makeRandomBool() bool {
	num := rand.Intn(1)
	if num == 1 {
		return true
	}
	return false
}

func makeRandomGene() string {
	num := genrand(0, 568, 0, 568, 5)
	return genes.Genes[num]
}

func getManyRandomPeople(num int) []RandomPerson {
	url := fmt.Sprintf("https://uinames.com/api/?region=england&amount=%v", num)

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return nil
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return nil
	}
	defer resp.Body.Close()
	var persons []RandomPerson

	if err := json.NewDecoder(resp.Body).Decode(&persons); err != nil {
		log.Println(err)
	}
	return persons

}

func genrand(bmin, bmax, rmin, rmax, n int) int {
	const randMax = 32767
	// Generalized random number generator;
	// sum of n random variables (usually 3).
	// Bell curve spans bmin<=x<bmax; then,
	// values outside rmin<=x<rmax are rejected.
	var sum, i, u int
	sum = 0
	for {
		for i = 0; i < n; i++ {
			sum += bmin + (rand.Intn(randMax) % (bmax - bmin))
		}
		if sum < 0 {
			sum -= n - 1
		}
		u = sum / n

		if rmin <= u && u < rmax {
			break
		}
	}
	return u
}

func createPatients(number int) int {
	for i := 0; i < number; i++ {
		id1++
		tempPatient := MakePatient(id1)
		repos.InsertPatient(&tempPatient)
	}
	return id1
}
func createResults(number int) int {
	id1 = 0
	for i := 0; i < number; i++ {
		id1++
		c := rand.Intn(2) + 1
		for j := 0; j < c; j++ {
			id4++
			tempResult := MakeResult(id1, id4)
			repos.InsertResult(&tempResult)
		}
	}
	return id4
}
func createResultDetails(number int) int {
	id1 = 0
	for i := 0; i < number; i++ {
		id1++
		c := rand.Intn(2) + 1
		for j := 0; j < c; j++ {
			id5++
			tempResultDetail := MakeResultDetail(id1, id5)
			repos.InsertResultDetail(&tempResultDetail)
		}
	}
	return id5
}

func createSamples(number int) int {
	id1 = 0
	for i := 0; i < number; i++ {
		id1++
		c := rand.Intn(5) + 1
		for j := 0; j < c; j++ {
			id2++
			tempSample := MakeSample(id1, id2)
			repos.InsertSample(&tempSample)
		}
	}
	return id2
}
func createExperiments(number int) int {
	id1 = 0
	for i := 0; i < number; i++ {
		id1++
		c := rand.Intn(5) + 1
		for j := 0; j < c; j++ {
			id3++
			tempExperiment := MakeExperiment(id1, id3)
			repos.InsertExperiment(&tempExperiment)
		}
	}
	return id3
}

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

//MakePatient makes a patient provided a patient ID
func MakePatient(patientID int) models.Patient {
	patient := models.Patient{}
	ClinicalHistory := makeRandomString()
	patient.ClinicalHistory = &ClinicalHistory
	DateReceived, _ := time.Parse(shortForm, makeRandomDate())
	patient.DateReceived = &DateReceived
	DateReported, _ := time.Parse(shortForm, makeRandomDate())
	patient.DateReported = &DateReported
	Dob, _ := time.Parse(shortForm, makeRandomDate())
	patient.Dob = &Dob
	FirstName := makeRandomName(patientID, false)
	patient.FirstName = &FirstName
	Gender := makeRandomString()
	patient.Gender = &Gender
	Initials := makeRandomString()
	patient.Initials = &Initials
	LastName := makeRandomName(patientID, true)
	patient.LastName = &LastName
	Mrn := makeRandomString()
	patient.Mrn = &Mrn
	OnHcn := makeRandomString()
	patient.OnHcn = &OnHcn
	PatientID := strconv.Itoa(patientID)
	if patientID > 0 {
		patient.PatientID = &PatientID
	}
	PatientType := makeRandomString()
	patient.PatientType = &PatientType
	ReferringPhysician := makeRandomString()
	//FIXTHIS
	patient.ReferringPhysician = &ReferringPhysician
	SeNum := makeRandomString()
	patient.SeNum = &SeNum
	SurgicalDate, _ := time.Parse(shortForm, makeRandomDate())
	patient.SurgicalDate = &SurgicalDate

	return patient
}

//MakeSample makes a sample provided a sample ID
func MakeSample(patientID int, sampleID int) models.Sample {
	sample := models.Sample{}
	SampleID := strconv.Itoa(sampleID)
	if sampleID > 0 {
		sample.SampleID = &SampleID
	}
	Facility := makeRandomString()
	sample.Facility = &Facility
	TestRequested := makeRandomString()
	sample.TestRequested = &TestRequested
	SeNum := makeRandomString()
	sample.SeNum = &SeNum
	DateCollected, _ := time.Parse(shortForm, makeRandomDate())
	sample.DateCollected = &DateCollected
	DateReceived, _ := time.Parse(shortForm, makeRandomDate())
	sample.DateReceived = &DateReceived
	SampleType := makeRandomString()
	sample.SampleType = &SampleType
	MaterialReceived := makeRandomString()
	sample.MaterialReceived = &MaterialReceived
	MaterialReceivedNum := makeRandomString()
	sample.MaterialReceivedNum = &MaterialReceivedNum
	MaterialReceivedOther := makeRandomString()
	sample.MaterialReceivedOther = &MaterialReceivedOther
	VolumeOfBloodMarrow := makeRandomFloat()
	sample.VolumeOfBloodMarrow = &VolumeOfBloodMarrow
	SurgicalNum := makeRandomString()
	sample.SurgicalNum = &SurgicalNum
	TumorSite := makeRandomString()
	sample.TumorSite = &TumorSite
	HistoricalDiagnosis := makeRandomString()
	sample.HistoricalDiagnosis = &HistoricalDiagnosis
	TumorPercntOfTotal := makeRandomFloat()
	sample.TumorPercntOfTotal = &TumorPercntOfTotal
	TumorPercntOfCircled := makeRandomFloat()
	sample.TumorPercntOfCircled = &TumorPercntOfCircled
	ReviewedBy := makeRandomString()
	sample.ReviewedBy = &ReviewedBy
	HESlideLocation := makeRandomString()
	sample.HESlideLocation = &HESlideLocation
	NonUhnID := makeRandomString()
	sample.NonUhnID = &NonUhnID
	NameOfRequestor := makeRandomString()
	sample.NameOfRequestor = &NameOfRequestor
	DnaConcentration := makeRandomFloat()
	sample.DnaConcentration = &DnaConcentration
	DnaVolume := makeRandomFloat()
	sample.DnaVolume = &DnaVolume
	DnaLocation := makeRandomString()
	sample.DnaLocation = &DnaLocation
	RnaConcentration := makeRandomFloat()
	sample.RnaConcentration = &RnaConcentration
	RnaVolume := makeRandomFloat()
	sample.RnaVolume = &RnaVolume
	RnaLocation := makeRandomString()
	sample.RnaLocation = &RnaLocation
	WbcLocation := makeRandomString()
	sample.WbcLocation = &WbcLocation
	PlasmaLocation := makeRandomString()
	sample.PlasmaLocation = &PlasmaLocation
	CfPlasmaLocation := makeRandomString()
	sample.CfPlasmaLocation = &CfPlasmaLocation
	PbBmLocation := makeRandomString()
	sample.PbBmLocation = &PbBmLocation
	RnaLysateLocation := makeRandomString()
	sample.RnaLysateLocation = &RnaLysateLocation
	SampleSize := makeRandomString()
	sample.SampleSize = &SampleSize
	StudyID := makeRandomString()
	sample.StudyID = &StudyID
	SampleName := makeRandomString()
	sample.SampleName = &SampleName
	DateSubmitted, _ := time.Parse(shortForm, makeRandomDate())
	sample.DateSubmitted = &DateSubmitted
	ContainerType := makeRandomString()
	sample.ContainerType = &ContainerType
	ContainerID := makeRandomString()
	sample.ContainerID = &ContainerID
	ContainerWell := makeRandomString()
	sample.ContainerWell = &ContainerWell
	CopathNum := makeRandomString()
	sample.CopathNum = &CopathNum
	OtherIdentifier := makeRandomString()
	sample.OtherIdentifier = &OtherIdentifier
	HasSampleFiles := makeRandomBool()
	sample.HasSampleFiles = &HasSampleFiles
	DnaSampleBarcode := makeRandomString()
	sample.DnaSampleBarcode = &DnaSampleBarcode
	DnaExtractionDate, _ := time.Parse(shortForm, makeRandomDate())
	sample.DnaExtractionDate = &DnaExtractionDate
	DnaQuality := makeRandomString()
	sample.DnaQuality = &DnaQuality
	FfpeQcDate, _ := time.Parse(shortForm, makeRandomDate())
	sample.FfpeQcDate = &FfpeQcDate
	DeltaCtValue := makeRandomFloat()
	sample.DeltaCtValue = &DeltaCtValue
	Comments := makeRandomString()
	sample.Comments = &Comments
	RnasePDate, _ := time.Parse(shortForm, makeRandomDate())
	sample.RnasePDate = &RnasePDate
	DnaQualityByRnaseP := makeRandomFloat()
	sample.DnaQualityByRnaseP = &DnaQualityByRnaseP
	RnaQuality := makeRandomFloat()
	sample.RnaQuality = &RnaQuality
	RnaExtractionDate, _ := time.Parse(shortForm, makeRandomDate())
	sample.RnaExtractionDate = &RnaExtractionDate
	PatientID := strconv.Itoa(patientID)
	sample.PatientID = &PatientID

	return sample
}

//MakeExperiment makes a experiemnt provided a experiment ID
func MakeExperiment(sampleID int, experimentID int) models.Experiment {
	experiment := models.Experiment{}
	ChipCartridgeBarcode := makeRandomString()
	experiment.ChipCartridgeBarcode = &ChipCartridgeBarcode
	CompleteDate, _ := time.Parse(shortForm, makeRandomDate())
	experiment.CompleteDate = &CompleteDate
	ExperimentID := strconv.Itoa(experimentID)
	if experimentID > 0 {
		experiment.ExperimentID = &ExperimentID
	}
	HasProjectFiles := makeRandomBool()
	experiment.HasProjectFiles = &HasProjectFiles
	OpenedDate, _ := time.Parse(shortForm, makeRandomDate())
	experiment.OpenedDate = &OpenedDate
	PanelAssayScreened := makeRandomString()
	experiment.PanelAssayScreened = &PanelAssayScreened
	Pcr := makeRandomString()
	experiment.Pcr = &Pcr
	Priority := makeRandomString()
	experiment.Priority = &Priority
	ProcedureOrderDatetime, _ := time.Parse(shortForm, makeRandomDate())
	experiment.ProcedureOrderDatetime = &ProcedureOrderDatetime
	ProjectID := makeRandomString()
	experiment.ProjectID = &ProjectID
	ProjectName := makeRandomString()
	experiment.ProjectName = &ProjectName
	SampleID := strconv.Itoa(sampleID)
	experiment.SampleID = &SampleID
	StudyID := makeRandomString()
	experiment.StudyID = &StudyID
	TestDate, _ := time.Parse(shortForm, makeRandomDate())
	experiment.TestDate = &TestDate

	return experiment
}

//MakeResult makes a result provided a result ID
func MakeResult(experimentID int, resultID int) models.Result {
	result := models.Result{}
	FailedRegions := makeRandomString()
	result.FailedRegions = &FailedRegions
	MeanDepthOfCoveage := makeRandomFloat()
	result.MeanDepthOfCoveage = &MeanDepthOfCoveage
	MlpaPcr := makeRandomString()
	result.MlpaPcr = &MlpaPcr
	Mutation := makeRandomString()
	result.Mutation = &Mutation
	OverallHotspotsThreshold := makeRandomFloat()
	result.OverallHotspotsThreshold = &OverallHotspotsThreshold
	OverallQualityThreshold := makeRandomFloat()
	result.OverallQualityThreshold = &OverallQualityThreshold
	ResultsID := strconv.Itoa(resultID)
	if resultID > 0 {
		result.ResultsID = &ResultsID
	}
	ExperimentID := strconv.Itoa(experimentID)
	result.ExperimentID = &ExperimentID
	UID := makeRandomString()
	result.UID = &UID
	VerificationPcr := makeRandomString()
	result.VerificationPcr = &VerificationPcr

	return result
}

//MakeResultDetail makes a result detail provided a resiltdetail ID
func MakeResultDetail(resultID int, resultdetailID int) models.Resultdetails {
	resultdetail := models.Resultdetails{}
	CNomenclature := makeRandomString()
	resultdetail.CNomenclature = &CNomenclature
	Coverage := int64(rand.Intn(1000))
	resultdetail.Coverage = &Coverage
	Exon := int64(rand.Intn(1000))
	resultdetail.Exon = &Exon
	Gene := makeRandomGene()
	resultdetail.Gene = &Gene
	Pcr := makeRandomString()
	resultdetail.Pcr = &Pcr
	PNomenclature := makeRandomString()
	resultdetail.PNomenclature = &PNomenclature
	QualityScore := makeRandomFloat()
	resultdetail.QualityScore = &QualityScore
	Result := makeRandomString()
	resultdetail.Result = &Result
	ResultsDetailsID := strconv.Itoa(resultdetailID)
	if resultdetailID > 0 {
		resultdetail.ResultsDetailsID = &ResultsDetailsID
	}
	ResultsID := strconv.Itoa(resultID)
	resultdetail.ResultsID = &ResultsID
	RiskScore := makeRandomFloat()
	resultdetail.RiskScore = &RiskScore
	UID := makeRandomString()
	resultdetail.UID = &UID
	VAF := makeRandomFloat()
	resultdetail.VAF = &VAF

	return resultdetail
}

//GetGenes fills the gene array struct
func (g *GeneArray) GetGenes() *GeneArray {
	path, _ := filepath.Abs("./models/genes.yaml")
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, g)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return g
}
