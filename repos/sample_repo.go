package repos

import (
	"fmt"
	"log"

	database "github.com/Bio-core/jtree/database"
	"github.com/Bio-core/jtree/models"
)

//GetAllSamples gets all and results a list of samples
func GetAllSamples(query string) []*models.Record {
	samples := []*models.Record{}
	err := database.DBSelect.Select(&samples, query)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return samples
}

//InsertSample allows users to add generic objects to a collection in the database
func InsertSample(sample *models.Sample) bool {
	stmt, err := database.DBUpdate.Prepare("INSERT INTO `samples` (`sample_id`, `facility`, `test_requested`, `se_num`, `date_collected`, `date_received`, `sample_type`, `material_received`, `material_received_num`, `material_received_other`, `volume_of_blood_marrow`, `surgical_num`, `tumor_site`, `historical_diagnosis`, `tumor_percnt_of_total`, `tumor_percnt_of_circled`, `reviewed_by`, `h_e_slide_location`, `non_uhn_id`, `name_of_requestor`, `dna_concentration`, `dna_volume`, `dna_location`, `rna_concentration`, `rna_volume`, `rna_location`, `wbc_location`, `plasma_location`, `cf_plasma_location`, `pb_bm_location`, `rna_lysate_location`, `sample_size`, `study_id`, `sample_name`, `date_submitted`, `container_type`, `container_name`, `container_id`, `container_well`, `copath_num`, `other_identifier`, `has_sample_files`, `dna_sample_barcode`, `dna_extraction_date`, `dna_quality`, `ffpe_qc_date`, `delta_ct_value`, `comments`, `rnase_p_date`, `dna_quality_by_rnase_p`, `rna_quality`, `rna_extraction_date`, `patient_id`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?); ")

	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(
		sample.SampleID,
		sample.Facility,
		sample.TestRequested,
		sample.SeNum,
		sample.DateCollected.Format(shortForm),
		sample.DateReceived.Format(shortForm),
		sample.SampleType,
		sample.MaterialReceived,
		sample.MaterialReceivedNum,
		sample.MaterialReceivedOther,
		sample.VolumeOfBloodMarrow,
		sample.SurgicalNum,
		sample.TumorSite,
		sample.HistoricalDiagnosis,
		sample.TumorPercntOfTotal,
		sample.TumorPercntOfCircled,
		sample.ReviewedBy,
		sample.HESlideLocation,
		sample.NonUhnID,
		sample.NameOfRequestor,
		sample.DnaConcentration,
		sample.DnaVolume,
		sample.DnaLocation,
		sample.RnaConcentration,
		sample.RnaVolume,
		sample.RnaLocation,
		sample.WbcLocation,
		sample.PlasmaLocation,
		sample.CfPlasmaLocation,
		sample.PbBmLocation,
		sample.RnaLysateLocation,
		sample.SampleSize,
		sample.StudyID,
		sample.SampleName,
		sample.DateSubmitted.Format(shortForm),
		sample.ContainerType,
		sample.ContainerName,
		sample.ContainerID,
		sample.ContainerWell,
		sample.CopathNum,
		sample.OtherIdentifier,
		sample.HasSampleFiles,
		sample.DnaSampleBarcode,
		sample.DnaExtractionDate,
		sample.DnaQuality,
		sample.FfpeQcDate,
		sample.DeltaCtValue,
		sample.Comments,
		sample.RnasePDate.Format(shortForm),
		sample.DnaQualityByRnaseP,
		sample.RnaQuality,
		sample.RnaExtractionDate.Format(shortForm),
		sample.PatientID)
	stmt.Close()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

//GetSampleByID gets all and results a list of samples
func GetSampleByID(ID string) *models.Sample {
	samples := []*models.Sample{}
	query := models.Query{}
	query.SelectedFields = make([]string, 0)
	query.SelectedFields = append(query.SelectedFields, "*")
	query.SelectedTables = make([]string, 0)
	query.SelectedTables = append(query.SelectedTables, "samples")
	query.SelectedCondition = make([][]string, 0)
	//query.SelectedCondition = append(query.SelectedCondition, make([]string, 0))
	conditions := []string{"AND", "samples.sample_id", "Equal to", ID}
	query.SelectedCondition = append(query.SelectedCondition, conditions)

	queryString := database.BuildQuery(query)
	err := database.DBSelect.Select(&samples, queryString)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if len(samples) == 0 {
		return nil
	}
	return samples[0]
}

//UpdateSample allows users to update generic objects to a collection in the database
func UpdateSample(sample *models.Sample) bool {
	stmt, err := database.DBUpdate.Prepare("UPDATE `samples` SET `sample_id` = ?,`facility` = ?,`test_requested` = ?,`se_num` = ?,`date_collected` = ?,`date_received` = ?,`sample_type` = ?,`material_received` = ?,`material_received_num` = ?,`material_received_other` = ?,`volume_of_blood_marrow` = ?,`surgical_num` = ?,`tumor_site` = ?,`historical_diagnosis` = ?,`tumor_percnt_of_total` = ?,`tumor_percnt_of_circled` = ?,`reviewed_by` = ?,`h_e_slide_location` = ?,`non_uhn_id` = ?,`name_of_requestor` = ?,`dna_concentration` = ?,`dna_volume` = ?,`dna_location` = ?,`rna_concentration` = ?,`rna_volume` = ?,`rna_location` = ?,`wbc_location` = ?,`plasma_location` = ?,`cf_plasma_location` = ?,`pb_bm_location` = ?,`rna_lysate_location` = ?,`sample_size` = ?,`study_id` = ?,`sample_name` = ?,`date_submitted` = ?,`container_type` = ?,`container_name` = ?,`container_id` = ?,`container_well` = ?,`copath_num` = ?,`other_identifier` = ?,`has_sample_files` = ?,`dna_sample_barcode` = ?,`dna_extraction_date` = ?,`dna_quality` = ?,`ffpe_qc_date` = ?,`delta_ct_value` = ?,`comments` = ?,`rnase_p_date` = ?,`dna_quality_by_rnase_p` = ?,`rna_quality` = ?,`rna_extraction_date` = ?,`patient_id` = ? WHERE `sample_id` = ?;")

	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(
		sample.SampleID,
		sample.Facility,
		sample.TestRequested,
		sample.SeNum,
		sample.DateCollected.Format(shortForm),
		sample.DateReceived.Format(shortForm),
		sample.SampleType,
		sample.MaterialReceived,
		sample.MaterialReceivedNum,
		sample.MaterialReceivedOther,
		sample.VolumeOfBloodMarrow,
		sample.SurgicalNum,
		sample.TumorSite,
		sample.HistoricalDiagnosis,
		sample.TumorPercntOfTotal,
		sample.TumorPercntOfCircled,
		sample.ReviewedBy,
		sample.HESlideLocation,
		sample.NonUhnID,
		sample.NameOfRequestor,
		sample.DnaConcentration,
		sample.DnaVolume,
		sample.DnaLocation,
		sample.RnaConcentration,
		sample.RnaVolume,
		sample.RnaLocation,
		sample.WbcLocation,
		sample.PlasmaLocation,
		sample.CfPlasmaLocation,
		sample.PbBmLocation,
		sample.RnaLysateLocation,
		sample.SampleSize,
		sample.StudyID,
		sample.SampleName,
		sample.DateSubmitted.Format(shortForm),
		sample.ContainerType,
		sample.ContainerName,
		sample.ContainerID,
		sample.ContainerWell,
		sample.CopathNum,
		sample.OtherIdentifier,
		sample.HasSampleFiles,
		sample.DnaSampleBarcode,
		sample.DnaExtractionDate.Format(shortForm),
		sample.DnaQuality,
		sample.FfpeQcDate.Format(shortForm),
		sample.DeltaCtValue,
		sample.Comments,
		sample.RnasePDate.Format(shortForm),
		sample.DnaQualityByRnaseP,
		sample.RnaQuality,
		sample.RnaExtractionDate.Format(shortForm),
		sample.PatientID,
		sample.SampleID)
	stmt.Close()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
