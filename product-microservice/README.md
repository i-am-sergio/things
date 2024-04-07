# Product Microservice

## Overview

This microservice manages product data, allowing operations such as creating, retrieving, updating, and deleting products. It integrates Golang with Echo for the RESTful API, MySQL for data storage, and Cloudinary for image handling.

## Technologies

- **Programming Language**: Golang
- **Web Framework**: Echo
- **Database**: MySQL

## Getting Started

### How to Run

Execute the following command in the root directory of the project:

```bash
    go run main.go
```

## API Endpoints

### Create a Product

**POST** `/products`  
Creates a new product entry. Use a multipart/form-data request with the following fields:

- `UserID` (type: integer)
- `Name` (type: text)
- `Description` (type: text)
- `Price` (type: number)
- `Category` (type: text)
- `Ubication` (type: text)
- `image` (type: file)

Example:

```plaintext
    UserID: 1
    Name: Hammer
    Description: NEW
    Price: 19.99
    Category: Home
    Ubication: Virginia, USA
    image: [Attach image file]
```

### Get All Products

**GET** `/products`  
Retrieves all products.

### Get Product by ID

**GET** `/products/:id`  
Retrieves a single product by its ID.

### Get Products by Category

**GET** `/products?category=:category`  
Fetches products filtered by the specified category.

### Search Products

**GET** `/products/search?q=:query`  
Performs a search among all products by name and description based on the query provided.

### Update a Product

**PUT** `/products/:id`  
Updates the product specified by ID. Use a multipart/form-data request with any of the fields you wish to update, including the image file if necessary.

Example:

```plaintext
    UserID: 2
    Name: Hammer Update
    Description: NEW Update
    Price: 15.99
    Category: Home Update
    Ubication: Virginia, USA Update
    image: [Attach new image file]
```

### Delete a Product

**DELETE** `/products/:id`  
Deletes the product and its associated comments based on the product ID.

### Set Premium Status for a Product by ID

**PUT** `/products/:id/premium`  
Toggles the premium status of a product specified by ID. This endpoint switches a product's premium status to the opposite of its current state (i.e., from premium to non-premium or vice versa).
