# Fitness API
RESTful Fitness API Developed in Go

Designed to provide developers uniform access to structures and datasets ready to support their Fitness applications. 

API Endpoints include interactions with the following base URIs:
- `/exercises`
- `/workouts`
- `/routines`
- `/users`

Each base URI is associated with a particular resource stored in a MySQL Database and all queries are formatted and returned as JSON. See [`Sample Requests`](#sample-requests) for more info

# Deploying API
This service consists of two containers
1. Golang Backend container routing requests through port `8080`
2. MySQL Container storing resources available through the API through port `3306`

Spin up both containers
```
docker-compose up --build -d
```

Create database & populate with provided `dump.sql`
```
bash populate_db.sh
```

Now you should be up and running! If you wish to run the unittests to verify everything is working properly: 
```
go test -v ./...
```

# Sample Requests
GET - Listing all existing exercises
```
curl -X GET localhost:8080/exercises
```

POST - Adding a new Exercise
```
curl -X POST -H "Content-Type: application/json" -d '{"id": 12345, "name": "some exercise", "category": "some category", "description": null}' localhost:8080/exercises
```

DELETE - Delete Workout records associated with workout name = `Push`
```
curl -X DELETE localhost:8080/workouts/name/Push
```

# Database
These are the following tables constructed by database dump (`dump.sql`) with some sample data for reference

## Exercises
```
mysql> SELECT * FROM Exercises;
+------------+------------------------------------------------------------+-------------+-------------+
| exerciseid | name                                                       | category    | description |
+------------+------------------------------------------------------------+-------------+-------------+
|          0 | ab crunch machine                                          | abdominals  | NULL        |
|          1 | ab roller                                                  | abdominals  | NULL        |
|          2 | adductor                                                   | adductors   | NULL        |
|          3 | adductor/groin                                             | adductors   | NULL        |
|          4 | advanced kettlebell windmill                               | abdominals  | NULL        |
|          5 | air bike                                                   | abdominals  | NULL        |
|          6 | all fours quad stretch                                     | quadriceps  | NULL        |
|          7 | alternate hammer curl                                      | biceps      | NULL        |
|          8 | alternate heel touchers                                    | abdominals  | NULL        |
|          9 | alternate incline dumbbell curl                            | biceps      | NULL        |
|         10 | alternate leg diagonal bound                               | quadriceps  | NULL        |
|         11 | alternating cable shoulder press                           | shoulders   | NULL        |
|         12 | alternating deltoid raise                                  | shoulders   | NULL        |
```

## Workouts
```
mysql> SELECT * FROM Workouts;
WIP
```

## Routines
```
mysql> SELECT * FROM Routines;
WIP
```

## Users
```
mysql> SELECT * FORM Users;
WIP
```


# Tips
These are some helpful commands & their usecase

Need to rebuild the Backend Container without affecting the MySQL Container?
```
docker-compose up --build -d --no-deps api
```

Need to run an SQL Query from outside the MySQL Container?
```
docker exec -i /bin/sh -c 'mysql -uroot -ppassword -e "CREATE DATABASE fitnessdb"'
```

Need to create database backup ([source](https://gist.github.com/spalladino/6d981f7b33f6e0afe6bb)?
```
docker exec db /usr/bin/mysqldump -u root --password=password fitnessdb > backup.sql
```

Need to restore database from backup ([source](https://gist.github.com/spalladino/6d981f7b33f6e0afe6bb)?
```
cat backup.sql | docker exec -i db /usr/bin/mysql -u root --password=password fitnessdb
```

Want to run a specific Test by name?
```
go test -v -run [TEST_NAME]
go test -v -run TestGetWorkoutsByWorkoutID
```
