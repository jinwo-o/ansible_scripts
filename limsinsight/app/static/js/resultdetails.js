var queryall = '{"selected_fields":["resultdetails.VAF", "resultdetails.c_nomenclature", "resultdetails.coverage", "resultdetails.exon", "resultdetails.gene", "resultdetails.p_nomenclature", "resultdetails.pcr", "resultdetails.quality_score", "resultdetails.result", "resultdetails.results_details_id", "resultdetails.results_id", "resultdetails.risk_score", "resultdetails.uid"], "selected_tables":["resultdetails"], "selected_conditions":[[]]}'
var queryurl = apihost + "/query";
$(document).ready(function() {
    var resultDetailsTable;
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
                resultsTable = $("#resultdetailstable").append(body);
            });
            $('#resultdetailstable tfoot th').each( function () {
                var title = $(this).text();
                $(this).html( '<input type="text" placeholder="Search '+title+'" />' );
            });
            resultDetailsTable = $( "#resultdetailstable").DataTable();
            resultDetailsTable.columns().every( function () {
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
    $('#resultdetailstable').on('click', 'tr', function(){
        if ($(this).hasClass('selected')) {
            $(this).removeClass('selected');
        } else {
            resultDetailsTable.$('tr.selected').removeClass('selected');
            $(this).addClass('selected');
        }
    })
});
