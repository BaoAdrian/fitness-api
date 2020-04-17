# User Endpoints

## Get All Exercises
**URL**: `/exercises`  
**Method**: `GET`  

### Success Response
**Code**: `200 OK`  
**Example**:  
```
$ curl -X GET localhost:8080/exercises
{
    "exercises": [
        {
            "category": "abdominals",
            "description": null,
            "exerciseid": 0,
            "name": "ab crunch machine"
        },
        {
            "category": "abdominals",
            "description": null,
            "exerciseid": 1,
            "name": "ab roller"
        },
        {
            "category": "adductors",
            "description": null,
            "exerciseid": 2,
            "name": "adductor"
        }, 
        ... more exercises ...
    ]
}
```

## Post A New Exercise
**URL**: `/exercises`  
**Method**: `POST`  

### Success Response
**Code**: `200 OK`  
```
$ curl -X POST -d '{"exerciseid":1000,"name":"eccentric barbell curls","category":"biceps","description":"Curl Barbell up then slow the eccentric movement, fighting gravity on the way down"}' http://localhost:8080/exercises
```


## Get All Exercise Names
**URL**: `/exercises/names`  
**Method**: `GET`  

### Success Response
**Code**: `200 OK`  
**Example**:  
```
$ curl -X GET http://localhost:8080/exercises/names
{
    "names": [
        "ab crunch machine",
        "ab roller",
        "adductor",
        "adductor/groin",
        "advanced kettlebell windmill",
        "air bike",
        "all fours quad stretch",
        "alternate hammer curl",
        ... more exercise names ...
    ]
}
```

## Get Exercise By Name
**URL**: `/exercises/name/{name}`  
**Method**: `GET`   

### Success Response
**Code**: `200 OK`  
**Example**:  
```
$ curl -X GET http://localhost:8080/exercises/name/ab%20roller
{
    "category": "abdominals",
    "description": null,
    "exerciseid": 1,
    "name": "ab roller"
}
```

## Delete Exercise By Name
**URL**: `/exercises/name/{name}`  
**Method**: `DELETE`   

### Success Response
**Code**: `200 OK`  
```
$ curl -X DELETE http://localhost:8080/exercises/name/ab%20roller
```


## Get All Exercise Categories
**URL**: `/exercises/categories`   
**Method**: `GET`   

### Success Response
**Code**: `200 OK`  
**Example**:  
```
$ curl -X GET http://localhost:8080/exercises/categories
{
    "categories": [
        {
            "category": "abdominals",
            "count": 90
        },
        {
            "category": "adductors",
            "count": 13
        },
        {
            "category": "quadriceps",
            "count": 144
        },
        ... more categories ...
    ]
}
```

## Get Exercises By Categories
**URL**: `/exercises/category/quadriceps`   
**Method**: `GET`   

### Success Response
**Code**: `200 OK`  
**Example**:  
```
$ curl -X GET http://localhost:8080/exercises/category/quadriceps
{
    "exercises": [
        {
            "category": "quadriceps",
            "description": null,
            "exerciseid": 6,
            "name": "all fours quad stretch"
        },
        {
            "category": "quadriceps",
            "description": null,
            "exerciseid": 10,
            "name": "alternate leg diagonal bound"
        },
        {
            "category": "quadriceps",
            "description": null,
            "exerciseid": 57,
            "name": "backward drag"
        },
        ... more exercises ...
    ]
}
```

## Get Exercises By Exercise ID
**URL**: `/exercises/id/1`   
**Method**: `GET`   

### Success Response
**Code**: `200 OK`  
**Example**:  
```
$ curl -X GET http://localhost:8080/exercises/id/1
{
    "category": "abdominals",
    "description": null,
    "exerciseid": 1,
    "name": "ab roller"
}
```

## Delete Exercises By Exercise ID
**URL**: `/exercises/id/1`   
**Method**: `DELETE`   

### Success Response
**Code**: `200 OK`  
```
$ curl -X DELETE http://localhost:8080/exercises/id/1
```
