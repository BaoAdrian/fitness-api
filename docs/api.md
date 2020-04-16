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