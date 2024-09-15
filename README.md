# Aneka Zoo
Aneka Zoo is a simple API application that allows users to get a list of animals and their details. The application is built using the Go and uses a PostgreSql database to store the data.

## Installation
1. Clone the repository
2. Run `go mod download` to install the dependencies
3. Run `docker-compose -f docker-storage.yaml up -d` to create container for the PostgreSql database, default database name is `aneka_zoo`
4. Run `go run main.go` to start the application
5. The application will run on `localhost:8080`

## API Documentation
### 1. Get All Animals
- URL: `http://localhost:8080/animals`
- Method: `GET`
- Response:
  ```json
  {
    "success": true,
    "message": "Success list animal",
    "trace_id": "ff15211a-62b2-40c1-afea-29ad2544ac33",
    "data": [
        {
            "id": 1,
            "name": "Sapi",
            "class": "mamal",
            "legs": 2,
            "created_at": "2024-09-15T01:47:26.35606Z"
        }
    ],
    "meta": {
        "current_page": 1,
        "last_page": 1,
        "total": 1,
        "per_page": 10
    }
  }
  ```
  
### 2. Get Animal By ID
- URL: `http://localhost:8080/animals/{id}`
- Method: `GET`
- Response:
  ```json
  {
    "success": true,
    "message": "Success get animal",
    "trace_id": "7aa9affa-190c-4444-a9f2-3d38dbf1c103",
    "data": {
        "id": 1,
        "name": "Sapi",
        "class": "mamal",
        "legs": 2
    }
  }
  ```

### 3. Create Animal
- URL: `http://localhost:8080/animals`
- Method: `POST`
- Request:
  ```json
  {
    "name": "Kucing",
    "class": "mamal",
    "legs": 4
  }
  ```
- Response:
  ```json
  {
    "success": true,
    "message": "Success create animal",
    "trace_id": "6ed92357-d854-4d6c-b430-8f96f64d1c8a"
  }
  ```
  
### 4. Update Animal
- URL: `http://localhost:8080/animals/{id}`
- Method: `PUT`
- Request:
  ```json
  {
    "name": "Kucing",
    "class": "mamal",
    "legs": 11
  }
  ```
- Response:
  ```json
  {
    "success": true,
    "message": "Success update animal",
    "trace_id": "5e3d4c78-a7a6-4ef3-bd87-dbd6fe431351"
  }
  ```
  
### 5. Delete Animal
- URL: `http://localhost:8080/animals/{id}`
- Method: `DELETE`
- Response:
  ```json
  {
    "success": false,
    "message": "Failed delete animal",
    "trace_id": "c1df911f-a78c-46e6-9acc-fa401b12e10a"
  }
  ```
  
## OR
- You can import the Postman collection from `Animal.postman_collection.json` to Postman to test the API


## Thank You
Thank you for test Aneka Zoo. If you have any questions or need help, please contact me at [email](mailto:m.bimagusta@gmail.com)