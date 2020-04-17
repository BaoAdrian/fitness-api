# Workout Endpoints

## Get All Workouts
**URL**: `/workouts`  
**Method**: `GET`  

### Success Response
**Code**: `200 OK`  
**Example**:  
```
$ curl -X GET localhost:8080/workouts
{
    "workouts": [
        {
            "exerciseid": 7,
            "repcount": 10,
            "routineid": 1,
            "setcount": 4
        },
        {
            "exerciseid": 27,
            "repcount": 8,
            "routineid": 1,
            "setcount": 4
        },
        ... more workouts ...
    ]
}
```

## Post A New Workout
**URL**: `/workouts`  
**Method**: `DELETE`  

### Success Response
**Code**: `200 OK`  
```
curl -X POST -d '{"exerciseid":1,"routineid":1,"setcount":3,"repcount":10}' http://localhost:8080/workouts
```

## Get Workout By Exercise ID
**URL**: `/workouts/exerciseid/{exerciseid}`  
**Method**: `GET`  

### Success Response
**Code**: `200 OK`  
**Example**:  
```
$ curl -X GET localhost:8080/workouts/exerciseid/7
{
    "workouts": [
        {
            "exerciseid": 7,
            "repcount": 10,
            "routineid": 1,
            "setcount": 4
        }
    ]
}
```

## Delete Workout By Exercise ID
**URL**: `/workouts/exerciseid/{exerciseid}`  
**Method**: `DELETE`  

### Success Response
**Code**: `200 OK`  
**Example**:  
```
$ curl -X DELETE localhost:8080/workouts/exerciseid/7
```

## Get Workout By Routine ID
**URL**: `/workouts/routineid/{routineid}`  
**Method**: `GET`  

### Success Response
**Code**: `200 OK`  
**Example**:  
```
$ curl -X GET localhost:8080/workouts/routineid/1
{
    "workouts": [
        {
            "exerciseid": 7,
            "repcount": 10,
            "routineid": 1,
            "setcount": 4
        },
        {
            "exerciseid": 27,
            "repcount": 8,
            "routineid": 1,
            "setcount": 4
        },
        {
            "exerciseid": 70,
            "repcount": 8,
            "routineid": 1,
            "setcount": 4
        },
        ... more exercise with routineid=1 ...
    ]
}
```

## Delete Workout By Routine ID
**URL**: `/workouts/routineid/{routineid}`  
**Method**: `DELETE`  

### Success Response
**Code**: `200 OK`  
**Example**:  
```
$ curl -X DELETE localhost:8080/workouts/routineid/1
```

## Get Workout By Primary IDs
This endpoint uses query parameters to provide the (exerciseid, routineid) Primary Key for the database to retrieve the workout corresponding to those ids.
**URL**: `/workouts/ids`    
**Method**: `GET`  

### Success Response
**Code**: `200 OK`  
**Example**:  
```
$ curl -X GET "localhost:8080/workouts/ids?exerciseid=27&routineid=1"
{
    "exerciseid": 27,
    "repcount": 8,
    "routineid": 1,
    "setcount": 4
}
```

## Delete Workout By Primary IDs
This endpoint uses query parameters to provide the (exerciseid, routineid) Primary Key for the database to retrieve the workout corresponding to those ids.
**URL**: `/workouts/ids`    
**Method**: `DELETE`  

### Success Response
**Code**: `200 OK`  
**Example**:  
```
$ curl -X DELETE "localhost:8080/workouts/ids?exerciseid=27&routineid=1"
```
