# GoLang Workshop

This project requires you to use a GoLang template to implement a RESTful API, perform unit testing, containerize the application with Docker, perform load testing, and profile the application using `pprof`. This document guides you through each task.

## Prerequisites

- GoLang
- Docker
- Grafana
- Make
- Basic knowledge about REST APIs, Unit Testing, Load Testing, and Profiling

## Tasks

### 1. Implement REST Endpoints

In the `app-no-implementation/internal/usecase` package, implement the REST endpoints.

### 2. Write Unit Tests

Create a new file `app-no-implementation/internal/usecase/user_test.go` and write the unit tests for the REST endpoints you have created in the previous task.

### 3. Dockerize the App

Use the Docker Compose command to build and run your app:

```shell
docker-compose up --build -d
```

### 4. Load Testing

Modify the app-no-implementation/scripts/load-test.sh script as necessary to perform load testing.


### 5. Grafana Monitoring

Open Grafana to monitor the load test results at http://localhost:8088/grafana. Analyze the logs and dashboards to understand the performance of your app.


### 6. Profiling Using pprof

Follow the pprof guide on the Go blog - https://go.dev/blog/pprof, implement profiling for your GoLang app and analyze the data.


## Acceptance Criteria:
- All APIs are working correctly.
- No errors are returned when you run the make lint command. 
- All tests pass when you run go test -v ./....