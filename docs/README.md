# API
Documentation of all supported endpoints, their methods, response, and sample request

# Endpoints
## `/users`
| Method | Endpoint | Sample Request | Successful Response |
| :--- | :--- | :---- | :--- |
| `GET` | `/users` | `curl -X GET http://localhost:8080/users` | [see here](./users.md#get-all-users) | 
| `POST` | `/users` | `curl -X POST -d '{"userid":1,"name":{"firstname":"John","lastname":"Doe"},"age":21,"weight":185.5}' http://localhost:8080/users` | [see here](./users.md#post-a-new-user) |
| `GET` | `/users/id/{userid}` | `curl -X GET http://localhost:8080/users/id/1` | [see here](./users.md#get-user-by-user-id) |
| `DELETE` | `/users/id/{userid}` | `curl -X DELETE http://localhost:8080/users/id/1` | [see here](./users.md#delete-user-by-user-id) |

## `/exercises`
| Method | Endpoint | Sample Request | Successful Response |
| :--- | :--- | :---- | :--- |
| `GET` | `/exercises` | `curl -X GET http://localhost:8080/exercises` | [see here](./exercises.md#get-all-exercises) | 
| `POST` | `/exercises` | `curl -X POST -d '{"exerciseid":1000,"name":"eccentric barbell curls","category":"biceps","description":"Curl Barbell up then slow the eccentric movement, fighting gravity on the way down"}' http://localhost:8080/exercises` | [see here](./exercises.md#post-a-new-exercise) |
| `GET` | `/exercises/names` | `curl -X GET http://localhost:8080/exercises/names` | [see here](./users.md#get-all-exercise-names) |
| `GET` | `/exercises/name/{name}` | `curl -X GET http://localhost:8080/exercises/name/ab%20roller` | [see here](./users.md#get-exercise-by-name) |
| `DELETE` | `/exercises/name/{name}` | `curl -X DELETE http://localhost:8080/exercises/name/ab%20roller` | [see here](./users.md#delete-exercise-by-name) |
| `GET` | `/exercises/categories` | `curl -X GET http://localhost:8080/exercises/categories` | [see here](./users.md#get-all-exercise-categories) |
| `GET` | `/exercises/category/{category}` | `curl -X GET http://localhost:8080/exercises/category/quadriceps` | [see here](./users.md#get-exercises-by-categories) |
| `GET` | `/exercises/id/{exerciseid}` | `curl -X GET http://localhost:8080/exercises/id/1` | [see here](./users.md#get-exercise-by-exercise-id) |
| `DELETE` | `/exercises/id/{exerciseid}` | `curl -X DELETE http://localhost:8080/exercises/id/1` | [see here](./users.md#delete-exercise-by-exercise-id) |