var queryall = '{"selected_fields":["patients.first_name", "patients.last_name", "patients.initials", "patients.gender", "patients.mrn", "patients.dob", "patients.on_hcn", "patients.clinical_history", "patients.patient_type", "patients.se_num", "patients.patient_id", "patients.date_received", "patients.referring_physician", "patients.date_reported", "patients.surgical_date"],"selected_tables":["patients"],"selected_conditions":[[]]}';
var queryurl = apihost + "/query";
$(document).ready(function() {
    var patientTable;
    var patientdata = $.ajax({
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
                body += "</tr>";
                $("#patienttable").append(body);
            });
            $('#patienttable tfoot th').each( function () {
                var title = $(this).text();
               $(this).html( '<input type="text" placeholder="Search '+title+'" />' );
            });
            patientTable = $("#patienttable").DataTable();
            patientTable.columns().every( function () {
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
    $('#patienttable').on('click', 'tr', function(){
        if ($(this).hasClass('selected')) {
            $(this).removeClass('selected');
        } else {
            patientTable.$('tr.selected').removeClass('selected');
            $(this).addClass('selected');
        }
    })
});
