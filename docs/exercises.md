# User Endpoints

## Get All Exercises
**URL**: `/exercises`  
**Method**: `GET`  

### Success Response
**Code**: `200 OK`  
**Example**:  
```
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


## Get All Exercise Names
**URL**: `/exercises/names`  
**Method**: `GET`  

### Success Response
**Code**: `200 OK`  
**Example**:  
```
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
Assumption: Request uses Exercise Name = 'ab roller` (urlencoded)
```
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


## Get All Exercise Categories
**URL**: `/exercises/categories`   
**Method**: `GET`   

### Success Response
**Code**: `200 OK`  
**Example**:  
```
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
Assumption: Request uses Exercise ID = 1
```
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
