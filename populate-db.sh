#!/bin/sh

# Create Database
docker exec -i db /bin/sh -c 'mysql -uroot -ppassword -e "CREATE DATABASE IF NOT EXISTS fitnessdb"'

# Populate with sql dump (if exists, otherwise, database will be empty)
if [ -f dump.sql ]
then
	cat dump.sql | docker exec -i db /usr/bin/mysql -uroot -ppassword fitnessdb		
fi
