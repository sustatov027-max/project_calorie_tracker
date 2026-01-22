# Calorie_Tracker RESTful-API
A Golang RESTful API application for calculating the daily calorie intake of a specific user

A Go service using the **Gin** framework and the **Gorm** ORM. The application and the **PostgreSQL** database are run in **Docker** containers using **Docker Compose**. **Makefile** commands are used for building and managing the service.

## Demonstration

## Installation & Run
### Prerequisites
1. [Go (version 1.20+)](https://go.dev/doc/install)
2. [Docker & Docker Compose](https://docs.docker.com/engine/install/)

### Clone repository
```bash
# Clone this project
git clone https://github.com/sustatov027-max/project_calorie_tracker.git
```

### Configuration
Before running the service, you need to configure the database connection. The configuration is set via environment variables in the **.env** file. Create a **.env** file in the project root and specify your values for the PostgreSQL connection and other parameters. Example file:
```.env
DB_CONFIG="host=localhost user=your_user password=your_password dbname=tracker_calories port=5433 sslmode=disable"
COST=14
SECRET="your_secret_key_here"
PORT=8080
```


### Quick Start (Development)
1. Start the data-base:
    ```bash
    make db-up
    ```
    This launches PostgreSQL on port 5433 with database tracker_calories.
2. Run the application:
    ```bash
    make dev
    ```
    This sets up environment variables and starts the Go server.

### Available Commands
* ```make dev``` - Start both database and application
* ```make db-up``` - Start only PostgreSQL database
* ```make db-down``` - Stop the database
* ```make test``` - Run Go tests

### API Endpoint : http://localhost:8080

## Structure
```
├── app
│   ├── app.go
│   ├── handler          // Our API core handlers
│   │   ├── common.go    // Common response functions
│   │   ├── projects.go  // APIs for Project model
│   │   └── tasks.go     // APIs for Task model
│   └── model
│       └── model.go     // Models for our application
├── config
│   └── config.go        // Configuration
└── main.go
```

## API

#### /projects
* `GET` : Get all projects
* `POST` : Create a new project

#### /projects/:title
* `GET` : Get a project
* `PUT` : Update a project
* `DELETE` : Delete a project

#### /projects/:title/archive
* `PUT` : Archive a project
* `DELETE` : Restore a project 

#### /projects/:title/tasks
* `GET` : Get all tasks of a project
* `POST` : Create a new task in a project

#### /projects/:title/tasks/:id
* `GET` : Get a task of a project
* `PUT` : Update a task of a project
* `DELETE` : Delete a task of a project

#### /projects/:title/tasks/:id/complete
* `PUT` : Complete a task of a project
* `DELETE` : Undo a task of a project