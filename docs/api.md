# API
Documentation of all supported endpoints, their methods, response, and sample request

# Endpoints
## `/users`
| Method | Endpoint | Sample Request | Successful Response |
| :--- | :--- | :---- | :--- |
| `GET` | `/users` | `curl -X GET http://localhost:8080/users` | [see here](./users/get_users.md) | 
| `POST` | `/users` | `curl -X POST -d '{"userid":1,"name":{"firstname":"John","lastname":"Doe"},"age":21,"weight":185.5}' http://localhost:8080/users` | [see here](./users/post_user.md)