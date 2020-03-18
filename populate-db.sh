#!/bin/sh

docker cp ./db/create-exercises.sql db:/
docker exec db /bin/sh -c 'mysql -uroot -ppassword </create-exercises.sql'