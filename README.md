# Gouros

A command line tool that parses yml files and generates boilerplate for API endpoints. Written specifically for Go REST API apps utilizing Gin. \
It will generate:

* Model: The model corresponding to the API resource
* Repository: For handling database operations 
* Service: For handling the logic of the request
* Controller: For binding data and params to objects and variables, delegating the request to the service and returning the response (data and HTTP status)
* Router: For registering all the functions in the controller to appropriate endpoints

Possible extensions:
* DB Schema: Generate the schema for creating the db table (for a specific database, i.e Postgres)
* Client Code?
* Relationships?
* Protected Routes?

## Example input

```
models:
  - table: employee
    attributes:
      - column: id 
        type: int64
        pk: true
        auto: true
      - column: name 
        type: string
      - column: surname 
        type: string
      - column: active 
        type: bool

entity:
  resource: employee
  description: Api Endpoints related to employees
  base: /employees
  routes:
    - endpoint: /
      method: GET
      description: Get all employees
    - endpoint: /$id
      method: GET
      description: Get employee by ID
    - endpoint: /
      method: POST
      description: Create new employee
      body: {}
    - endpoint: /$id
      method: PUT
      description: Edit an existing employee
      body: {}
    - endpoint: /$id
      method: DELETE
      description: Delete employee with given ID
```
