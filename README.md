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

## Approach number 2: Using Postman
It is also possible to use the API service Postman to make HTTP requests like GET, POST, PUT and UPDATE.\\
To be able to use Postman, the user needs to log into Postman using the link: https://www.postman.com/. \\
Next, the user needs to sign up and make an user, afterwards they are able to log in to Postman. \\
Here the user is able to make HTTP requests by entering the link for the to-do-app: http://localhost:8080, \\
and then select the wanted HTTP request.

# How to run the application
In order to run the application the user must have Docker running on their computer.
To build the program the user needs to write 
```
docker-compose up --build
```
in the terminal. The application will build and map the containers port 8080 to hosts port 8080, \\
thus the application can be accessed through the link http://localhost:8080. \\
If the user decides that they want to quit, they can gracefully close the application by writing
```
docker-compose down
```
which will stop and remove the containers.