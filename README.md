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
