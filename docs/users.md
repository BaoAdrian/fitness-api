# User Endpoints

## Get All Users
**URL**: `/users`  
**Method**: `GET`  

### Success Response
**Code**: `200 OK`  
**Example**:  
```
$ curl -x GET localhost:8080/users
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
```
$ curl -X POST -d '{"userid":1,"name":{"firstname":"John","lastname":"Doe"},"age":21,"weight":185.5}' http://localhost:8080/users
```


## Get User By User ID
**URL**: `/users/id/{userid}`  
**Method**: `GET`  

### Success Response
**Code**: `200 OK`  
**Example**:  
```
$ curl -X GET localhost:8080/users/id/1
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

## Delete User By User ID
**URL**: `/users/id/{userid}`  
**Method**: `DELETE`  

### Success Response
**Code**: `200 OK`  
```
$ curl -X DELETE http://localhost:8080/users/id/1
```