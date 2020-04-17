# Routine Endpoints

## Get All Routines
**URL**: `/routines`  
**Method**: `GET`  

### Success Response
**Code**: `200 OK`  
**Example**:  
```
$ curl -X GET localhost:8080/routines
{
    "routines": [
        {
            "day": 1,
            "description": "Some compound back movements with isolation bicep exercises",
            "name": "My Super Awesome Routine",
            "routineid": 1,
            "userid": 1
        },
        ... more routines ...
    ]
}
```

## Post A New Routine
**URL**: `/routines`  
**Method**: `POST`  

### Success Response
**Code**: `200 OK`  
```
$ curl -X POST -d '{"routineid":2,"userid":1,"name":"some name","day":2,"description":"some description"}' http://localhost:8080/routines
```

## Get Routine By Routine ID
**URL**: `/routines/routineid/{routineid}`  
**Method**: `GET`  

### Success Response
**Code**: `200 OK`  
**Example**:  
```
$ curl -X GET localhost:8080/routines/routineid/1
{
    "day": 1,
    "description": "Some compound back movements with isolation bicep exercises",
    "name": "My Super Awesome Routine",
    "routineid": 1,
    "userid": 1
}
```

## Delete Routine By Routine ID
**URL**: `/routines/routineid/{routineid}`  
**Method**: `DELETE`  

### Success Response
**Code**: `200 OK`  
**Example**:  
```
$ curl -X DELETE localhost:8080/routines/routineid/1
```

## Get Routine By User ID
**URL**: `/routines/userid/{userid}`  
**Method**: `GET`  

### Success Response
**Code**: `200 OK`  
**Example**:  
```
$ curl -X GET localhost:8080/routines/userid/1
{
    "routines": [
        {
            "day": 1,
            "description": "Some compound back movements with isolation bicep exercises",
            "name": "My Super Awesome Routine",
            "routineid": 1,
            "userid": 1
        },
        ... more routines associated with userid=1 ...
    ]
}
```

## Delete Routine By User ID
**URL**: `/routines/userid/{userid}`  
**Method**: `DELETE`  

### Success Response
**Code**: `200 OK`  
**Example**:  
```
$ curl -X DELETE localhost:8080/routines/userid/1
```
