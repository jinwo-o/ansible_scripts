var queryall = '{"selected_fields":["experiments.experiment_id", "experiments.study_id", "experiments.panel_assay_screened", "experiments.test_date", "experiments.chip_cartridge_barcode", "experiments.complete_date", "experiments.pcr", "experiments.sample_id", "experiments.project_name", "experiments.priority", "experiments.opened_date", "experiments.project_id", "experiments.has_project_files", "experiments.procedure_order_datetime"], "selected_tables":["experiments"], "selected_conditions":[[]]}'
var queryurl = apihost + "/query";
$(document).ready(function() {
    var experimentTable;
    var experimentdata = $.ajax({
        type: 'POST',
        url: queryurl,
        contentType: 'application/json; charset=utf-8',
        dataType: 'json',
        data: queryall,
        dataSrc: "",
        success: function(data) {
            $.each(data, function(i, data) {
                // console.log(data);
                var body = "<tr>";
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
                body += "</tr>";
                $("#experimenttable").append(body);
            });
            $('#experimenttable tfoot th').each( function () {
                var title = $(this).text();
                $(this).html( '<input type="text" placeholder="Search '+title+'" />' );
            });
            experimentTable = $("#experimenttable").DataTable();
            experimentTable.columns().every( function () {
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
    $('#experimenttable').on('click', 'tr', function(){
        if ($(this).hasClass('selected')) {
            $(this).removeClass('selected');
        } else {
            experimentTable.$('tr.selected').removeClass('selected');
            $(this).addClass('selected');
        }
    })
});