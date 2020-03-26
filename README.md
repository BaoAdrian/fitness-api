# Fitness API
RESTful API Developed in Go

# Deploying
Spin up services
```
docker-compose up -d --build
```

Populate mysql database using provided script
```
bash populate_db.sh
```

Run testing suite
```
go test -v
```
Current supported tests
- API Endpoints
- DB Queries


# Requests
GET - Listing all existing exercises
```
curl -X GET -H "Content-Type: application/json" localhost:8080/exercises
```

POST - Adding a new Exercise
```
curl -X POST -H "Content-Type: application/json" -d '{"id": 12345, "name": "some exercise", "category": "some category", "description": null}' localhost:8080/exercises
```