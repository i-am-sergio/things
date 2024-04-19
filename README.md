# **_Things_**

[![auth-microservice CI](https://github.com/i-am-sergio/things/actions/workflows/auth-microservice.yml/badge.svg)](https://github.com/i-am-sergio/things/actions/workflows/auth-microservice.yml)
[![notifications-microservice CI](https://github.com/i-am-sergio/things/actions/workflows/notifications-microservice.yml/badge.svg)](https://github.com/i-am-sergio/things/actions/workflows/notifications-microservice.yml)
[![product-microservice CI](https://github.com/i-am-sergio/things/actions/workflows/product-microservice.yml/badge.svg)](https://github.com/i-am-sergio/things/actions/workflows/product-microservice.yml)

![SonarCloud](https://img.shields.io/badge/Sonar%20cloud-F3702A?style=for-the-badge&logo=sonarcloud&logoColor=white)
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![GraphQL](https://img.shields.io/badge/-GraphQL-E10098?style=for-the-badge&logo=graphql&logoColor=white)
![React Native](https://img.shields.io/badge/React_Native-20232A?style=for-the-badge&logo=react&logoColor=61DAFB)
![Vue](https://img.shields.io/badge/Vue.js-35495E?style=for-the-badge&logo=vue.js&logoColor=4FC08D)
![TypeScript](https://img.shields.io/badge/typescript-%23007ACC.svg?style=for-the-badge&logo=typescript&logoColor=white)

## **_Server Endpoints_**

### 1. **_[Auth & Users Microservice Endpoints](./auth-microservice/README.md)_**

### 2. **_[Products Microservice Endpoints](./product-microservice/README.md)_**

### 3. **_[Notifications-Microservice Endpoints](./notifications-microservice/README.md)_**

### 4. **_[Ads Microservice Endpoints](documentation.md)_**

### 5. **_[Chat Microservice Endpoints](documentation.md)_**

## **_Docker Compose_**

- **Build**
    ```bash
    docker compose create    
    ```
- **Build and Initialize**
    ```bash
    docker compose up -d   
    ```
- **Only Initialize**
    ```bash
    docker compose start    
    ```
- **Stopping**
    ```bash
    docker compose stop  
    ```
- **Stop and Delete Multi-container**
    ```bash
    docker compose down
    ```