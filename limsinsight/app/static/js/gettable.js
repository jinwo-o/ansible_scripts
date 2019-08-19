var isHidden = false;
$(document).ready(function() {

    var html = '<table id="basictable" class="display" style="width:100%"><thead><tr><th>Element</th></tr></thead><tbody>';
    var data = $.getJSON("http://localhost:8000/Jtree/metadata/0.1.0/searchable", function(data) {
        for(var i in data) {
            html += '<tr>' + '<td>' + data[i] + '</td>' + '</tr>';
        }
        html += '</tbody></table>';
        document.getElementById("studytable").innerHTML = html;
    });
    $("#basictable").DataTable();
});
