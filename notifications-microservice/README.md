# ***Notifications Microservice***

<!-- https://github.com/Ileriayo/markdown-badges -->
<!-- https://dev.to/envoy_/150-badges-for-github-pnk -->

<!-- [![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=things_things&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=things_things) -->

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![SonarCloud](https://img.shields.io/badge/Sonar%20cloud-F3702A?style=for-the-badge&logo=sonarcloud&logoColor=white)

This document outlines the endpoints provided by the Notification Microservice implemented in Golang.

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

