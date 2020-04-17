# Fitness API
RESTful Fitness API Developed in Go

API Endpoints include interactions with the following base URIs:
- `/exercises`
- `/workouts`
- `/routines`
- `/users`

Each base URI is associated with a particular resource stored in a MySQL Database and all queries are formatted and returned as JSON. See [`API Documentation`](./docs/README.md) for more info

# Components
The project consists of 
- [Golang Docker container](https://hub.docker.com/_/golang) to host the API served on port `8080`
- [MySQL Docker container](https://hub.docker.com/_/mysql) to host the resources accessed/inserted/updated by the API served on port `3306`

# Deploying 
Spin up the services
```
docker-compose up --build -d
```

An sql dump is provided for you if you wish to populate the database with the 873-count exercise dataset & associated tables   
```
bash populate_db.sh
```

Services should be up and running!   

If you wish to run the unittests to verify everything is working properly: 
```
go test -v ./...
```

See [`API Documentation`](./docs/README.md) for more information on requests & endpoint definitions


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
... more exercises ... 
```

## Workouts
```
mysql> SELECT * FROM Workouts;
+-------------+------------+-----------+-----------+
| exercise_id | routine_id | set_count | rep_count |
+-------------+------------+-----------+-----------+
|           7 |          1 |         4 |        10 |
|          27 |          1 |         4 |         8 |
|          70 |          1 |         4 |         8 |
|         102 |          1 |         4 |        10 |
|         365 |          1 |         4 |        10 |
|         651 |          1 |         4 |        10 |
... more workouts ...
```

## Routines
```
mysql> SELECT * FROM Routines;
+------------+---------+--------------------------+-------------------------------------------------------------+------+
| routine_id | user_id | routine_name             | description                                                 | day  |
+------------+---------+--------------------------+-------------------------------------------------------------+------+
|          1 |       1 | My Super Awesome Routine | Some compound back movements with isolation bicep exercises |    1 |
... more routines ...
```

## Users
```
mysql> SELECT * FROM Users;
+---------+------------+-----------+------+--------+
| user_id | first_name | last_name | age  | weight |
+---------+------------+-----------+------+--------+
|       1 | John       | Doe       |   21 | 185.50 |
|       2 | Jane       | Doe       |   20 | 135.0  |
... more users ...
```


# Tips
These are some helpful commands you may wish to use & their usecase

Need to rebuild the Backend Container without affecting the MySQL Container?
```
docker-compose up --build -d --no-deps api
```

Need to open a shell to the db container?
```
docker exec -it db /bin/sh -c 'mysql -uroot -ppassword'
```

Need to run an SQL Query from outside the MySQL Container?
```
docker exec -i db /bin/sh -c 'mysql -uroot -ppassword -e "CREATE DATABASE fitnessdb"'
```

Need to create database dump ([source](https://gist.github.com/spalladino/6d981f7b33f6e0afe6bb))?
```
docker exec db /usr/bin/mysqldump -u root --password=password fitnessdb > backup.sql
```

Need to restore database from dump ([source](https://gist.github.com/spalladino/6d981f7b33f6e0afe6bb))?
```
cat backup.sql | docker exec -i db /usr/bin/mysql -u root --password=password fitnessdb
```

Want to run a specific Test by name?
```
go test -v -run [TEST_NAME]
go test -v -run TestGetWorkoutsByWorkoutID
```
