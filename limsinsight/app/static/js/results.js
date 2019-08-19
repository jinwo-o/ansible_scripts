var queryall = '{"selected_fields":["results.failed_regions", "results.mean_depth_of_coveage", "results.mlpa_pcr", "results.mutation", "results.overall_hotspots_threshold", "results.overall_quality_threshold", "results.results_id", "results.uid", "results.verification_pcr", "results.experiment_id"], "selected_tables":["results"], "selected_conditions":[[]]}'
// var queryall = '{"selected_fields":["results.failed_regions", "results.mean_depth_of_coveage", "results.mlpa_pcr", "results.mutation", "results.overall_hotspots_threshold", "results.overall_quality_threshold", "results.results_id", "results.uid", "results.verification_pcr"], "selected_tables":["results"], "selected_conditions":[[]]}'
var queryurl = apihost + "/query";
$(document).ready(function() {
    var resultsTable;
    var experimentdata = $.ajax({
        type: 'POST',
        url: queryurl,
        contentType: 'application/json; charset=utf-8',
        dataType: 'json',
        data: queryall,
        dataSrc: "",
        success: function(data) {
            $.each(data, function(i, data) {
                var body = "<tr>";
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
                body += "</tr>";
                $("#resultstable").append(body);
            });
            $('#resultstable tfoot th').each( function () {
                var title = $(this).text();
                $(this).html( '<input type="text" placeholder="Search '+title+'" />' );
            });
            resultsTable = $("#resultstable").DataTable();
            resultsTable.columns().every( function () {
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
    $('#resultstable').on('click', 'tr', function(){
        if ($(this).hasClass('selected')) {
            $(this).removeClass('selected');
        } else {
            resultsTable.$('tr.selected').removeClass('selected');
            $(this).addClass('selected');
        }
    })
});