# Fitness API
RESTful Fitness API Developed in Go

Designed to provide developers uniform access to structures and datasets ready to support their Fitness applications. 

API Endpoints include interactions with the following base URIs:
- `/exercises`
- `/workouts`
- `/routines`
- `/users`

Each base URI is associated with a particular resource stored in a MySQL Database and all queries are formatted and returned as JSON. See `[Sample Requests](#sample-requests)` for more info

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
| exerciseid | name | category | description | 
| :--- | :--- | :--- | :--- |
| 1 | ab crunch machine | abdominals | NULL |
| 2 | ab roller | abdominals | NULL |
| 3 | adductor | abdominals | NULL |

## Workouts
```
mysql> SELECT * FROM Workouts;
+-----------+-----------------+------------+----------+----------+
| workoutid | name            | exerciseid | setcount | repcount |
+-----------+-----------------+------------+----------+----------+
|         1 | Back & Biceps   |          7 |        4 |       10 |
|         1 | Back & Biceps   |         70 |        4 |        8 |
|         1 | Back & Biceps   |        365 |        4 |       10 |
|         1 | Back & Biceps   |        651 |        4 |       10 |
|         1 | Back & Biceps   |         27 |        4 |        8 |
|         1 | Back & Biceps   |        102 |        4 |       10 |
|         2 | Chest & Triceps |        247 |        4 |       10 |
|         2 | Chest & Triceps |         69 |        5 |        5 |
|         2 | Chest & Triceps |        384 |        4 |        8 |
|         2 | Chest & Triceps |        147 |        4 |       15 |
|         2 | Chest & Triceps |        850 |        4 |       15 |
|         2 | Chest & Triceps |        550 |        4 |       15 |
+-----------+-----------------+------------+----------+----------+
```
| workoutid | name | exerciseid | setcount | repcount |
| :--- | :--- | :--- | :--- | :--- |
| 1 | Back & Biceps | 7 | 4 | 10 |
| 1 | Back & Biceps | 70 | 4 | 8 |
| 1 | Back & Biceps | 365 | 4 | 10 |

## Routines
WIP

## Users
WIP