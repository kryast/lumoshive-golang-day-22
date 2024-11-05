<!-- method create user -->
curl -X POST http://localhost:8080/create -d '{"name": "Ahmad 5", "username" : "user5", "password" : "pass"}'

<!-- method create task -->
curl -X POST http://localhost:8080/users/create-task -d '{"description" : "description 3"}'

<!-- method get user detail -->
curl -X GET http://localhost:8080/users/1

<!-- method update task -->
curl -X POST http://localhost:8080/users/list-task/1

<!-- method login user -->
curl -X POST http://localhost:8080/login -d '{"username": "user1", "password": "pass"}'