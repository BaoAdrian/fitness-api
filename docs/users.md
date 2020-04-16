# User Endpoints

## Get All Users
**URL**: `/users`  
**Method**: `GET`  

### Success Response
**Code**: `200 OK`  
**Example**:  
Assumption: Two users, John Doe (userid=1) and Jane Doe (userid=2) exist in the database
```
{
    "users": [
        {
            "age": 21,
            "name": {
                "firstname": "John",
                "lastname": "Doe"
            },
            "userid": 1,
            "weight": 185.5
        },
        {
            "age": 20,
            "name": {
                "firstname": "Jane",
                "lastname": "Doe"
            },
            "userid": 2,
            "weight": 130.5
        }
    ]
}
```

## Post A New User
**URL**: `/users`  
**Method**: `POST`  

### Success Response
**Code**: `200 OK`  


## Get User By User ID
**URL**: `/users/id/{userid}`  
**Method**: `GET`  

### Success Response
**Code**: `200 OK`  
**Example**:
Assumption: Requests uses User ID associated with John Doe (userid=1)
```
{
    "age": 21,
    "name": {
        "firstname": "John",
        "lastname": "Doe"
    },
    "userid": 1,
    "weight": 185.5
}
```

## Post A New User
**URL**: `/users/id/{userid}`  
**Method**: `DELETE`  

### Success Response
**Code**: `200 OK`  
