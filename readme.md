Log Converter
==========================================

Log converter program for parsing log files and write to the database mongo.


Running
===============================

- ```docker-compose up -d``` - for running docker mongo
- ```go install``` - for installing
- ```log-converter /tmp/logs.txt 1``` - running process for monitoring file /tmp/logs.txt


Testing
================================
- ```docker-compose up -d```
- ```go test . ```
- 

Lookup stored data
===============================

go to http://localhost:8080/logs/