//var queryall = '{"selected_fields":["*""], "selected_tables":["samples", "patients", "experiments", "results", "resultdetails"], "selected_conditions":[[]]}'
var queryall = '{"selected_fields":["*"], "selected_tables":["patients", "samples", "experiments", "results", "resultdetails"], "selected_conditions":[[]]}'
var queryurl = apihost + "/query";
$(document).ready(function() {
    var allTable;
    var alldata = $.ajax({
        type: 'POST',
        url: queryurl,
        contentType: 'application/json; charset=utf-8',
        dataType: 'json',
        dataSrc: "",
        data: queryall,
        success: function(data) {
            $.each(data, function(i, data) {
                var body = "<tr>";
                body += "<td>" + data["patients.mrn"] + "</td>";
                body += "<td>" + data["patients.se_num"] + "</td>";
                body += "<td>" + data["patients.first_name"] + "</td>";
                body += "<td>" + data["patients.last_name"] + "</td>";
                body += "<td>" + data["patients.initials"] + "</td>";
                body += "<td>" + data["patients.gender"] + "</td>";
                body += "<td>" + data["patients.dob"] + "</td>";
                body += "<td>" + data["patients.on_hcn"] + "</td>";
                body += "<td>" + data["patients.clinical_history"] + "</td>";
                body += "<td>" + data["patients.patient_type"] + "</td>";
                body += "<td>" + data["patients.patient_id"] + "</td>";
                body += "<td>" + data["patients.date_received"] + "</td>";
                body += "<td>" + data["patients.referring_physician"] + "</td>";
                body += "<td>" + data["patients.date_reported"] + "</td>";
                body += "<td>" + data["patients.surgical_date"] + "</td>";
                body += "<td>" + data["samples.sample_id"] + "</td>";
                body += "<td>" + data["samples.facility"] + "</td>";
                body += "<td>" + data["samples.test_requested"] + "</td>";
                body += "<td>" + data["samples.se_num"] + "</td>";
                body += "<td>" + data["samples.date_collected"] + "</td>";
                body += "<td>" + data["samples.date_received"] + "</td>";
                body += "<td>" + data["samples.sample_type"] + "</td>";
                body += "<td>" + data["samples.material_received"] + "</td>";
                body += "<td>" + data["samples.material_received_num"] + "</td>";
                body += "<td>" + data["samples.material_received_other"] + "</td>";
                body += "<td>" + data["samples.volume_of_blood_marrow"] + "</td>";
                body += "<td>" + data["samples.surgical_num"] + "</td>";
                body += "<td>" + data["samples.tumor_site"] + "</td>";
                body += "<td>" + data["samples.historical_diagnosis"] + "</td>";
                body += "<td>" + data["samples.tumor_percnt_of_total"] + "</td>";
                body += "<td>" + data["samples.tumor_percnt_of_circled"] + "</td>";
                body += "<td>" + data["samples.reviewed_by"] + "</td>";
                body += "<td>" + data["samples.h_e_slide_location"] + "</td>";
                body += "<td>" + data["samples.non_uhn_id"] + "</td>";
                body += "<td>" + data["samples.name_of_requestor"] + "</td>";
                body += "<td>" + data["samples.dna_concentration"] + "</td>";
                body += "<td>" + data["samples.dna_volume"] + "</td>";
                body += "<td>" + data["samples.dna_location"] + "</td>";
                body += "<td>" + data["samples.rna_concentration"] + "</td>";
                body += "<td>" + data["samples.rna_volume"] + "</td>";
                body += "<td>" + data["samples.rna_location"] + "</td>";
                body += "<td>" + data["samples.wbc_location"] + "</td>";
                body += "<td>" + data["samples.plasma_location"] + "</td>";
                body += "<td>" + data["samples.cf_plasma_location"] + "</td>";
                body += "<td>" + data["samples.pb_bm_location"] + "</td>";
                body += "<td>" + data["samples.rna_lysate_location"] + "</td>";
                body += "<td>" + data["samples.sample_size"] + "</td>";
                body += "<td>" + data["samples.study_id"] + "</td>";
                body += "<td>" + data["samples.sample_name"] + "</td>";
                body += "<td>" + data["samples.date_submitted"] + "</td>";
                body += "<td>" + data["samples.container_type"] + "</td>";
                body += "<td>" + data["samples.container_name"] + "</td>";
                body += "<td>" + data["samples.container_id"] + "</td>";
                body += "<td>" + data["samples.container_well"] + "</td>";
                body += "<td>" + data["samples.copath_num"] + "</td>";
                body += "<td>" + data["samples.other_identifier"] + "</td>";
                body += "<td>" + data["samples.has_sample_files"] + "</td>";
                body += "<td>" + data["samples.dna_sample_barcode"] + "</td>";
                body += "<td>" + data["samples.dna_extraction_date"] + "</td>";
                body += "<td>" + data["samples.dna_quality"] + "</td>";
                body += "<td>" + data["samples.ffpe_qc_date"] + "</td>";
                body += "<td>" + data["samples.delta_ct_value"] + "</td>";
                body += "<td>" + data["samples.comments"] + "</td>";
                body += "<td>" + data["samples.rnase_p_date"] + "</td>";
                body += "<td>" + data["samples.dna_quality_by_rnase_p"] + "</td>";
                body += "<td>" + data["samples.rna_quality"] + "</td>";
                body += "<td>" + data["samples.rna_extraction_date"] + "</td>";
                body += "<td>" + data["samples.patient_id"] + "</td>";
                body += "<td>" + data["experiments.experiment_id"] + "</td>";
                body += "<td>" + data["experiments.study_id"] + "</td>";
                body += "<td>" + data["experiments.panel_assay_screened"] + "</td>";
                body += "<td>" + data["experiments.test_date"] + "</td>";
                body += "<td>" + data["experiments.chip_cartridge_barcode"] + "</td>";
                body += "<td>" + data["experiments.complete_date"] + "</td>";
                body += "<td>" + data["experiments.pcr"] + "</td>";
                body += "<td>" + data["experiments.sample_id"] + "</td>";
                body += "<td>" + data["experiments.project_name"] + "</td>";
                body += "<td>" + data["experiments.priority"] + "</td>";
                body += "<td>" + data["experiments.opened_date"] + "</td>";
                body += "<td>" + data["experiments.project_id"] + "</td>";
                body += "<td>" + data["experiments.has_project_files"] + "</td>";
                body += "<td>" + data["experiments.procedure_order_datetime"] + "</td>";
                body += "<td>" + data["results.failed_regions"] + "</td>";
                body += "<td>" + data["results.mean_depth_of_coveage"] + "</td>";
                body += "<td>" + data["results.mlpa_pcr"] + "</td>";
                body += "<td>" + data["results.mutation"] + "</td>";
                body += "<td>" + data["results.overall_hotspots_threshold"] + "</td>";
                body += "<td>" + data["results.overall_quality_threshold"] + "</td>";
                body += "<td>" + data["results.results_id"] + "</td>";
                body += "<td>" + data["results.uid"] + "</td>";
                body += "<td>" + data["results.verification_pcr"] + "</td>";
                body += "<td>" + data["results.experiment_id"] + "</td>";
                body += "<td>" + data["resultdetails.VAF"] + "</td>";
                body += "<td>" + data["resultdetails.c_nomenclature"] + "</td>";
                body += "<td>" + data["resultdetails.coverage"] + "</td>";
                body += "<td>" + data["resultdetails.exon"] + "</td>";
                body += "<td>" + data["resultdetails.gene"] + "</td>";
                body += "<td>" + data["resultdetails.p_nomenclature"] + "</td>";
                body += "<td>" + data["resultdetails.pcr"] + "</td>";
                body += "<td>" + data["resultdetails.quality_score"] + "</td>";
                body += "<td>" + data["resultdetails.result"] + "</td>";
                body += "<td>" + data["resultdetails.results_details_id"] + "</td>";
                body += "<td>" + data["resultdetails.results_id"] + "</td>";
                body += "<td>" + data["resultdetails.risk_score"] + "</td>";
                body += "<td>" + data["resultdetails.uid"] + "</td>";
                body += "</tr>";
                $("#alltable").append(body);
            });
            $('#alltable tfoot th').each( function () {
                var title = $(this).text();
               $(this).html( '<input type="text" placeholder="Search '+title+'" />' );
            });
            allTable = $("#alltable").DataTable();
            allTable.columns().every( function () {
                var that = this;
 
                $( 'input', this.footer() ).on( 'keyup change', function () {
                    if ( that.search() !== this.value ) {
                        that
                            .search( this.value )
                            .draw();
                    }
                });
            });
        }
    });
    $('#alltable').on('click', 'tr', function(){
        if ($(this).hasClass('selected')) {
            $(this).removeClass('selected');
        } else {
            allTable.$('tr.selected').removeClass('selected');
            $(this).addClass('selected');
        }
    })
});
