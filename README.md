# JTree

Database Access API for AMDL Results Database

[View On SwaggerHub](https://app.swaggerhub.com/apis/JTree/jtree-metadata_api/0.1.0)
</br>
[![Build Status](https://travis-ci.org/Bio-Core/JTree.svg?branch=master)](https://travis-ci.org/Bio-Core/JTree)
[![Go Report Card](https://goreportcard.com/badge/Bio-core/Jtree)](https://goreportcard.com/report/Bio-core/Jtree)

## Getting started

This project requires Go to be installed. On OS X with Homebrew you can just run `brew install go`.

Installing this code:
`$ go get github.com/Bio-core/JTree`

To get any missing dependencies run:
`$ bash go_get.sh`

Running it then should be as simple as:

```console
$ make database
$ make build
$ ./bin/jtree
```
To generate fake data, run `$ ./bin/jtree -g=100`, where 100 is the amount of dummy data requested
To run on a spesific port run `$ ./bin/jtree -p=8000`, where 8000 is the desired port



Endpoints:

```sh
# QUERIES
# This will return all of the columns in the database
$ curl http://127.0.0.1:8000/Jtree/metadata/0.1.0/columns

# This will return all of the searchable fields in the database
$ curl http://127.0.0.1:8000/Jtree/metadata/0.1.0/searchable

# This will return all of the uneditable fields in the database
$ curl http://127.0.0.1:8000/Jtree/metadata/0.1.0/uneditable

# This is an example query that will return all data from every table in the database
$ curl http://127.0.0.1:8000/Jtree/metadata/0.1.0/query -X POST -H "content-type:application/json" /
-d '{"selected_fields":["*"],"selected_tables":["samples", "patients","experiments", "results", "resultdetails"],"selected_conditions":[[]]}'

# This is an example query that will return all data where the date of birth is greater than 1950
$ curl http://127.0.0.1:8000/Jtree/metadata/0.1.0/query -X POST -H "content-type:application/json" -d /
'{"selected_fields":["samples.sample_id", "patients.dob"],"selected_tables":["samples", "patients","experiments", "results", "resultdetails"],"selected_conditions":[["AND", "patients.dob", "Greater than", "1950"]]}'


# INSERTS and UPDATES
# Note that inserts and updates operate the same way, the only difference is that the public key is not passed with the object structure for in insert
# samples
$ curl -X POST -H "Content-Type: application/json" /
 -d '{`# See models.samples for object structure`}' 127.0.0.1:8000/Jtree/metadata/0.1.0/sample

# patients
$ curl -X POST -H "Content-Type: application/json" /
 -d '{`# See models.patients for object structure`}' 127.0.0.1:8000/Jtree/metadata/0.1.0/patient

# experiments
$ curl -X POST -H "Content-Type: application/json" /
 -d '{`# See models.experiments for object structure`}' 127.0.0.1:8000/Jtree/metadata/0.1.0/experiment

 # results
$ curl -X POST -H "Content-Type: application/json" /
 -d '{`# See models.results for object structure`}' 127.0.0.1:8000/Jtree/metadata/0.1.0/result

 # Example:
 $ curl -X POST -H "Content-Type: application/json" -d '{"results.failed_regions":"ABC", "results.mean_depth_of_coveage":928.123, "results.mlpa_pcr":"ABCD", "results.mutation":"EFG", "results.overall_hotspots_threshold":419.668, "results.overall_quality_threshold":123.234, "results.uid":"Jin", "results.verification_pcr":"Hwang"}' localhost:8000/Jtree/metadata/0.1.0/result


 # resultdetails
$ curl -X POST -H "Content-Type: application/json" /
 -d '{`# See models.resultdetails for object structure`}' 127.0.0.1:8000/Jtree/metadata/0.1.0/resultdetails

```


# Docker
For golang code, may need to change the connection string with an updated IP address for the docker container
```bash
$ docker inspect <container name> | grep IPAddr
```
Then git commit and puch to master (because it does a goget) and run the docker commands below
```bash
$ docker network create -d bridge mysql-network
$ docker run --name mysqldb -p 3306:3306 -e MYSQL_ROOT_PASSWORD=waterloo -d --network=mysql-network mysql/mysql-server
$ docker exec -i mysqldb mysql -u root -pwaterloo -e "CREATE DATABASE JTree"
$ docker exec -i mysqldb mysql -u root -pwaterloo JTree < ./sql/jtree_backup.sql
$ docker exec -i mysqldb mysql -u root -pwaterloo -e "CREATE USER 'select'@'%' identified by 'passwords';grant SELECT on JTree.* to 'select'@'%';flush privileges;CREATE USER 'update'@'%' identified by 'passwordu';grant SELECT,INSERT, UPDATE on JTree.* to 'update'@'%';flush privileges;"
$ docker build -t docker/jtree .
$ docker run --network=mysql-network --name jtree -p 8000:8000 -d docker/jtree
```
