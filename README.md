# Calorie_Tracker RESTful-API
A Golang RESTful API application for calculating the daily calorie intake of a specific user

A Go service using the **Gin** framework and the **Gorm** ORM. The application and the **PostgreSQL** database are run in **Docker** containers using **Docker Compose**. **Makefile** commands are used for building and managing the service.

## Demonstration
<table>
  <tr>
    <td style="vertical-align: top; width: 70%; padding: 10px;">
      <img src=".github/img/Registration.bmp" alt="Registartion" style="width: 90%; height: 400px; object-fit: contain;">
    </td>
    <td style="vertical-align: top; width: 70%; padding: 10px;">
      <img src=".github/img/Login.bmp" alt="Login" style="width: 90%; height: 400px; object-fit: contain;">
    </td>
  </tr>
  <tr>
    <td style="vertical-align: top; padding: 10px;">
      <img src=".github/img/Create product.bmp" alt="Create product" style="width: 90%; height: 400px; object-fit: contain;">
    </td>
    <td style="vertical-align: top; padding: 10px;">
      <img src=".github/img/Create meal.bmp" alt="Create meal" style="width: 90%; height: 400px; object-fit: contain;">
    </td>
  </tr>
  <tr>
    <td style="vertical-align: top;" colspan="2" padding: 10px;>
      <img src=".github/img/Summary meals.bmp" alt="Summary meal" style="width: 50%; height: 400px; object-fit: contain;">
    </td>
  </tr>
</table>

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
* `make dev` - Start both database and application
* `make db-up` - Start only PostgreSQL database
* `make db-down` - Stop the database
* `make test` - Run Go tests

### API Endpoint : http://localhost:8080

## Structure
```
tracker-calories/
├── cmd/
│   └── main.go      # Application entry point
├── internal/
│   ├── handlers/
│   │   ├── diary_handler.go          # Diary handlers
│   │   ├── product_handler.go        # Product handlers
│   │   ├── product_handler_test.go   # Unit tests for product
│   │   ├── user_handler.go           # User handlers
│   │   ├── user_handler_test.go      # Unit tests for user
│   │   └── mock/
│   │       ├── mock_product_services.go      # Mocks for product_services
│   │       └── mock_user_service.go          # Mocks for user_services
│   ├── models/
│   │   ├── daySummary.go    # Summary model
│   │   ├── mealLog.go       # MealLog model
│   │   ├── product.go       # Product model
│   │   └── user.go          # User model
│   ├── repositories/
│   │   ├── diary_repository.go     # Working with the DB: diary
│   │   ├── product_repository.go   # Working with the DB: product
│   │   └── user_repository.go      # Working with the DB: user
│   ├── services/
│   │   ├── diary_service.go    # Business logic of diary
│   │   ├── user_service.go     # Business logic of users
│   │   └── product_service.go  # Business logic of products
│   ├── validation/
│   │   ├── validator.go        # Validation of input data
│   └── middlewares/
│       └── authMiddleware.go   # Middleware authentication
├── pkg/
│   ├── database/
│   │   └── db.go            # Initializing the connection to the database
│   └── utils/
│       ├── context.go        # Gin context util
│       └── hashPasword.go    # Password hash utilities
├── Dockerfile           # Dockerfile of the application
├── docker-compose.yml   # Docker Compose config
├── .env                 # Local environment variables (in .gitignore)
├── .gitignore
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

## API

### Authentication
- **POST** `/auth/register` - Register a new user
- **POST** `/auth/login` - Login user

### User Profile
- **GET** `/me` - Get current user profile (Auth required)

### Food Diary (Meals)
- **GET** `/diary` - Get all meals for current day (Auth required)
- **POST** `/diary` - Add a new meal to diary (Auth required)
- **PUT** `/diary/:id` - Update specific meal (Auth required)
- **DELETE** `/diary/:id` - Remove meal from diary (Auth required)

### Diary Analytics
- **GET** `/diary/summary` - Get daily calorie summary (Auth required)

### Food Products
- **GET** `/products` - Get all food products (Auth required)
- **POST** `/products` - Create new food product (Auth required)
- **PUT** `/products/:id` - Update food product (Auth required)
- **DELETE** `/products/:id` - Delete food product (Auth required)

---

**Note:** All endpoints except `/auth/*` require JWT token in the `Authorization` header as `Bearer <token>`.

## Example requests
1. Register User
    ```bash
    curl -X POST http://localhost:8080/auth/register \
    -H "Content-Type: application/json" \
    -d '{"name":"Test","age":19,"email":"test@mail.ru","password":"12345678","weight":76,"height":184,"gender":"male","activeDays":3}'
    ```
    Response:

    ```json
    {
        "id": 1,
	    "name": "Test",
		"age": 19,
		"email": "test@mail.ru",
		"weight": 76,
		"height": 184,
		"gender": "male",
		"activeDays": 3,
		"caloriesNorm": 2950.16
    }
    ```
2. Login User
    ```bash
    curl -X POST http://localhost:8080/auth/login \
    -H "Content-Type: application/json" \
    -d '{"email": "test@mail.ru", "password": "12345678"}'
    ```
    Response:

    ```json
    {
        "token":"mock-jwt-token-12345",
        "token_type":"Bearer"
    }
    ```
3. Get User Profile (Auth Required)
    ```bash
    curl -X GET http://localhost:8080/me \
    -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
    ```
    Response:

    ```json
    {
        "id": 1,
        "name": "Test",
        "age": 19,
        "email": "test@mail.ru",
        "weight": 76,
        "height": 184,
        "gender": "male",
        "activeDays": 3,
        "caloriesNorm": 2950.16
    }
    ```
4. Add Meal to Diary
    ```bash
    curl -X POST http://localhost:8080/diary \
    -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
    -H "Content-Type: application/json" \
    -d '{"product_id": 123, "gramms": 100}'
    ```
    Response:

    ```json
    {
    "id": 1,
    "user_id": 1,
    "product_id": 123,
    "product": ...,
    "gramms": 100,
    "created_at": "2024-01-22T13:45:00Z"
    }
    ```
5. Get Today's Meals
    ```bash
    curl -X GET http://localhost:8080/diary \
    -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
    ```
    Response:

    ```json
    [
        {
            "id": 1,
            "user_id": 1,
            "product_id": 123,
            "product": [...],
            "gramms": 100,
            "created_at": "2024-01-22T13:45:00Z"
        }
    ]
    ```
6. Create Food Product
    ```bash
    curl -X POST http://localhost:8080/products \
    -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
    -H "Content-Type: application/json" \
    -d '{"name": "Avocado", "calories": 160, "proteins": 2, "fats": 15, "carbohydrates": 9}'
    ```
    Response:

    ```json
    {
    "id": 124,
    "name": "Avocado",
    "calories": 160,
    "proteins": 2,
    "fats": 15,
    "carbohydrates": 9,
    "created_at": "2024-01-22T15:00:00Z"
    }
    ```
7. Get Daily Summary
    ```bash
    curl -X GET http://localhost:8080/diary/summary \
    -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
    ```
    Response:

    ```json
    {
        "meals": [...],
        "total_calories": 1850,
        "total_proteins": 83,
        "total_fats": 57,
        "total_carbs": 188,
        "daily_norm": 2100,
        "remaining": 350
    }
    ```
