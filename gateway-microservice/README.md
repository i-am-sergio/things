# GraphQL Gateway-Microservice

This project implements a GraphQL gateway microservice using Go, GraphQL, and Apollo Server.

## Setup

### Install gqlgen
```bash
go get github.com/99designs/gqlgen
```
### Setup Tools
```bash
printf '// +build tools\npackage tools\nimport _ "github.com/99designs/gqlgen"' | gofmt > tools.go
go mod tidy
```
### Initialize gqlgen Configuration
```bash
go run github.com/99designs/gqlgen init
```
This command initializes gqlgen configuration, creating necessary files for schema generation and resolver implementation.

## Running the Server
This command starts the GraphQL server, which serves as the gateway microservice for handling GraphQL queries and mutations.

```bash
go run server.go
```    

## File Structure

- `server.go`: Entry point of the application. Initializes and starts the GraphQL server.
- `graphql/schema.graphqls`: GraphQL schema definition file. Contains the schema definition for the API.
- `graphql/schema.resolvers.go`: Auto-generated file by gqlgen. Contains resolver implementations.
- `graphql/generated.go`: Auto-generated file by gqlgen. Contains type definitions and GraphQL schema registration.
- `resolver.go`: Custom resolver implementations for handling GraphQL queries and mutations.
- `model.go`: Definition of data models used in the application.
- `db`: Directory containing database connection logic and configurations.
- `tools.go`: File used for tooling purposes. Ensures that gqlgen is included as a dependency only for development.

## Usage

1. Install dependencies:

```bash
go mod tidy
```    

2. Run the server:

```bash
go run server.go
```    
 

