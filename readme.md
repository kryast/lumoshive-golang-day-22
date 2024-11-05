<!-- method create user -->
curl -X POST http://localhost:8080/create -d '{"name": "Ahmad 5", "username" : "user5", "password" : "pass"}'

<!-- method create task -->
curl -X POST http://localhost:8080/users/create-task -d '{"description" : "description 3"}'