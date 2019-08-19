var queryall = '{"selected_fields":["patients.first_name", "patients.last_name", "samples.sample_id", "patients.mrn", "samples.sample_id", "samples.facility", "samples.test_requested", "samples.se_num", "samples.date_collected", "samples.date_received", "samples.sample_type", "samples.material_received", "samples.material_received_num", "samples.material_received_other", "samples.volume_of_blood_marrow", "samples.surgical_num", "samples.tumor_site", "samples.historical_diagnosis", "samples.tumor_percnt_of_total", "samples.tumor_percnt_of_circled", "samples.reviewed_by", "samples.h_e_slide_location", "samples.non_uhn_id", "samples.name_of_requestor", "samples.dna_concentration", "samples.dna_volume", "samples.dna_location", "samples.rna_concentration", "samples.rna_volume", "samples.rna_location", "samples.wbc_location", "samples.plasma_location", "samples.cf_plasma_location", "samples.pb_bm_location", "samples.rna_lysate_location", "samples.sample_size", "samples.study_id", "samples.sample_name", "samples.date_submitted", "samples.container_type", "samples.container_name", "samples.container_id", "samples.container_well", "samples.copath_num", "samples.other_identifier", "samples.has_sample_files", "samples.dna_sample_barcode", "samples.dna_extraction_date", "samples.dna_quality", "samples.ffpe_qc_date", "samples.delta_ct_value", "samples.comments", "samples.rnase_p_date", "samples.dna_quality_by_rnase_p", "samples.rna_quality", "samples.rna_extraction_date", "samples.patient_id"],"selected_tables":["patients", "samples"],"selected_conditions":[[]]}';
var queryurl = apihost + "/query";
$(document).ready(function() {
    var sampleTable;
    var sampledata = $.ajax({
        type: 'POST',
        url: queryurl,
        contentType: 'application/json; charset=utf-8',
        dataType: 'json',
        data: queryall,
        dataSrc: "",
        success: function(data) {
            $.each(data, function(i, data) {
                var body = "<tr>";
                body += "<td>" + data["patients.first_name"] + "</td>";
                body += "<td>" + data["patients.last_name"] + "</td>";
                body += "<td>" + data["patients.mrn"] + "</td>";
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
                $("#sampletable tbody").append(body);
            });
            $('#sampletable tfoot th').each( function () {
                var title = $(this).text();
                $(this).html( '<input type="text" placeholder="Search '+title+'" />' );
            });
            sampleTable = $("#sampletable").DataTable();
            sampleTable.columns().every( function () {
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
    $('#sampletable').on('click', 'tr', function(){
        if ($(this).hasClass('selected')) {
            $(this).removeClass('selected');
        } else {
            sampleTable.$('tr.selected').removeClass('selected');
            $(this).addClass('selected');
        }
    })
});