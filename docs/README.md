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
| `GET` | `/exercises/names` | `curl -X GET http://localhost:8080/exercises/names` | [see here](./exercises.md#get-all-exercise-names) |
| `GET` | `/exercises/name/{name}` | `curl -X GET http://localhost:8080/exercises/name/ab%20roller` | [see here](./exercises.md#get-exercise-by-name) |
| `DELETE` | `/exercises/name/{name}` | `curl -X DELETE http://localhost:8080/exercises/name/ab%20roller` | [see here](./exercises.md#delete-exercise-by-name) |
| `GET` | `/exercises/categories` | `curl -X GET http://localhost:8080/exercises/categories` | [see here](./exercises.md#get-all-exercise-categories) |
| `GET` | `/exercises/category/{category}` | `curl -X GET http://localhost:8080/exercises/category/quadriceps` | [see here](./exercises.md#get-exercises-by-categories) |
| `GET` | `/exercises/id/{exerciseid}` | `curl -X GET http://localhost:8080/exercises/id/1` | [see here](./exercises.md#get-exercise-by-exercise-id) |
| `DELETE` | `/exercises/id/{exerciseid}` | `curl -X DELETE http://localhost:8080/exercises/id/1` | [see here](./exercises.md#delete-exercise-by-exercise-id) |

## `/workouts`
| Method | Endpoint | Sample Request | Successful Response |
| :--- | :--- | :---- | :--- |
| `GET` | `/workouts` | `curl -X GET http://localhost:8080/workouts` | [see here](./workouts.md#get-all-exercises) | 
| `POST` | `/workouts` | `curl -X POST -d '{"exerciseid":1,"routineid":1,"setcount":3,"repcount":10}' http://localhost:8080/workouts` | [see here](./workouts.md#post-a-new-workout) |
| `GET` | `/workouts/exerciseid/{exerciseid}` | `curl -X GET http://localhost:8080/workouts/exerciseid/1` | [see here](./workouts.md#get-workout-by-exercise-id) |
| `DELETE` | `/workouts/exerciseid/{exerciseid}` | `curl -X DELETE http://localhost:8080/workouts/exerciseid/1` | [see here](./workouts.md#delete-workout-by-exercise-id) |
| `GET` | `/workouts/routineid/{routineid}` | `curl -X GET http://localhost:8080/workouts/routineid/1` | [see here](./workouts.md#get-workout-by-routine-id) |
| `DELETE` | `/workouts/routineid/{routineid}` | `curl -X DELETE http://localhost:8080/workouts/routineid/1` | [see here](./workouts.md#delete-workout-by-routine-id) |
| `GET` | `/workouts/ids` | `curl -X GET http://localhost:8080/workouts/ids?routineid=1&exerciseid=1` | [see here](./workouts.md#get-workout-by-primary-ids) |
| `DELETE` | `/workouts/ids` | `curl -X DELETE http://localhost:8080/workouts/ids?routineid=1&exerciseid=1` | [see here](./workouts.md#delete-workout-by-primary-ids) |

## `/routines`
| Method | Endpoint | Sample Request | Successful Response |
| :--- | :--- | :---- | :--- |
| `GET` | `/routines` | `curl -X GET http://localhost:8080/routines` | [see here](./routines.md#get-all-routines) | 
| `POST` | `/routines` | `curl -X POST -d '{"exerciseid":1,"routineid":1,"setcount":3,"repcount":10}' http://localhost:8080/routines` | [see here](./routines.md#post-a-new-routine) |
| `GET` | `/routines/routineid/{routineid}` | `curl -X GET http://localhost:8080/routines/routineid/1` | [see here](./routines.md#get-routine-by-routine-id) |
| `DELETE` | `/routines/routineid/{routineid}` | `curl -X DELETE http://localhost:8080/routines/routineid/1` | [see here](./routines.md#delete-routine-by-routine-id) |
| `GET` | `/routines/userid/{userid}` | `curl -X GET http://localhost:8080/routines/userid/1` | [see here](./routines.md#get-routine-by-user-id) |
| `DELETE` | `/routines/userid/{userid}` | `curl -X DELETE http://localhost:8080/routines/userid/1` | [see here](./routines.md#delete-routine-by-user-id) |
