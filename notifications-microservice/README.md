# ***Notifications Microservice***

<!-- https://github.com/Ileriayo/markdown-badges -->
<!-- https://dev.to/envoy_/150-badges-for-github-pnk -->

<!-- [![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=things_things&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=things_things) -->

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![SonarCloud](https://img.shields.io/badge/Sonar%20cloud-F3702A?style=for-the-badge&logo=sonarcloud&logoColor=white)

## **Dependencies**

- [**github.com/joho/godotenv**](https://github.com/joho/godotenv) v1.5.1
  - **Description:** `godotenv` is a Go package that loads environment variables from a `.env` file into the process environment.
  - **Usage in Project:** Used for loading environment variables from `.env` files to configure the application.

- [**github.com/labstack/echo/v4**](https://github.com/labstack/echo/v4) v4.11.4
  - **Description:** `echo` is a high-performance, minimalist web framework for Go. Version 4 is a major release with significant improvements and changes over previous versions.
  - **Usage in Project:** Used for defining routes, handling HTTP requests, and building RESTful APIs.

- [**github.com/stretchr/testify**](https://github.com/stretchr/testify) v1.9.0
  - **Description:** `testify` is a toolkit with common assertions and mocks that plays nicely with the standard Go testing package.
  - **Usage in Project:** Used for writing unit tests and assertions in tests for the application.

- [**go.mongodb.org/mongo-driver**](https://go.mongodb.org/mongo-driver) v1.14.0
  - **Description:** `mongo-driver` is the official Go driver for MongoDB. It provides tools for interacting with MongoDB databases from Go applications.
  - **Usage in Project:** Used for connecting to MongoDB databases, performing CRUD operations, and interacting with MongoDB collections.



## **Endpoints**

### 1. **Get Notification by ID**

- **Method:** GET
- **Endpoint:** `/notifications/:notification_id`
- **Description:** Retrieves a notification by its ID.
- **Parameters:**
  - `notification_id`: The ID of the notification to retrieve.
- **Response:**
  - Status Code: 200 OK
  - Body: JSON representation of the notification.

### 2. **Get Notifications by User ID**

- **Method:** GET
- **Endpoint:** `/notifications/user/:user_id`
- **Description:** Retrieves all notifications associated with a specific user.
- **Parameters:**
  - `user_id`: The ID of the user to retrieve notifications for.
- **Response:**
  - Status Code: 200 OK
  - Body: JSON array of notification objects.

### 3. **Create Notification**

- **Method:** POST
- **Endpoint:** `/notifications`
- **Description:** Creates a new notification.
- **Request Body:**
  - JSON representation of the notification to create.
- **Response:**
  - Status Code: 201 Created
  - Body: JSON representation of the created notification.

### 4. **Mark Notification as Read**

- **Method:** PUT
- **Endpoint:** `/notifications/markAsRead/:notification_id`
- **Description:** Marks a notification as read by its ID.
- **Parameters:**
  - `notification_id`: The ID of the notification to mark as read.
- **Response:**
  - Status Code: 200 OK

### 5. **Mark All Notifications as Read for a User**

- **Method:** PUT
- **Endpoint:** `/notifications/markAllAsRead/:user_id`
- **Description:** Marks all notifications as read for a specific user.
- **Parameters:**
  - `user_id`: The ID of the user to mark all notifications as read for.
- **Response:**
  - Status Code: 200 OK

## **Dockerization**

- **Build Image:** Run the following command:
  ```bash
  docker build -t notifications-mcsv .
  ```
- **Run App Container:** Run the following command:
  ```bash
  docker run -d -p 8005:8005 notifications-mcsv
  ```
