#!/bin/sh

docker cp ./db/create-exercises.sql sql-container:/
docker exec sql-container /bin/sh -c 'mysql -uroot -ppassword </create-exercises.sql'