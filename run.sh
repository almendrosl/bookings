#!/bin/bash

go build -o booking cmd/web/*.go
./booking -dbname=bookings -dbuser=root -cache=false -production=false -dbpass=toor
# host=localhost port=5432 dbname=bookings user=root password=toor