let editorServer = require('datatables.net-editor-server');
var express = require("express");
var app = express();
var router = express.Router()
var path = __dirname + "/";

// Setup the static location for images, code and stylesheets
app.use(express.static(path));

router.use(function(req, res, next) {
    console.log("/" + req.method);
    next();
});

router.get("/", function(req, res) {
    res.sendFile(path + "app/view/main.html");
});

router.get("/experiments", function(req, res) {
    res.sendFile(path + "app/view/experiments.html");
})

router.get("/patients", function(req, res) {
    res.sendFile(path + "app/view/patients.html");
})

router.get("/results", function(req, res) {
    res.sendFile(path + "app/view/results.html");
})

router.get("/resultdetails", function(req, res) {
    res.sendFile(path + "app/view/resultdetails.html");
})

router.get("/samples", function(req, res) {
    res.sendFile(path + "app/view/samples.html");
});

router.get("/all", function(req, res) {
    res.sendFile(path + "app/view/all.html");
});

app.use("/", router);
var listener = app.listen(8003, function() {
    console.log("Listening on port: " + listener.address().port);
});
