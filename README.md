# SAD-ToDo-List-Application

This application allows the user to:
- POST activities that need to be done
- GET a list of all the activities that are on the to-do-list
- GET an unique activity on the to-do-list based on its id
- UPDATE an existing activity using the PUT command
- DELETE an existing activity using the DELETE command

These actions can be performed using two different methods:
1. using curl commands in the command line interface (CLI)
2. using the service Postman

## Approach number 1: Using curl commands

**1. List All Items (GET Request)**
```
curl -X GET http://localhost:8080/items
```

**2. Add a New Item (POST Request)**
```
curl -X POST http://localhost:8080/items -H "Content-Type: application/json" -d '{"title":"New Task","description":"Description of task"}'
```

**3. Retrieve an Item by ID (GET Request)**
```
curl -X GET http://localhost:8080/items/{id}
```

**4. Update an Item by ID (PUT Request)**
```
curl -X PUT http://localhost:8080/items/{id} -H "Content-Type: application/json" -d '{"title":"Updated Task","description":"Updated description"}'
```

**5. Delete an Item by ID (DELETE Request)**
```
curl -X DELETE http://localhost:8080/items/{id}
```