# Example of API made with Golang and JWT (JSON Web Token)
Based on [this YouTube video](https://www.youtube.com/watch?v=YA6cVebkwJE).

## Is it necessary to create a [.env file](#example-of-required-parameters-inside-env-file) in the root folder
Example of main struture files
```
golang_jwt
└───api
    ├───auto
    ├───config
    ├───controllers
    │   └───users
    ├───database
    ├───models
    ├───repository
    │   └───crud
    ├───responses
    ├───router
    │   └───routes
    ├───security
    └───utils
        ├───channels
        └───console
└───.env
└───.gitignore
└───go.mod
└───go.sum
└───main.go
└───README.md
```

## Example of required parameters inside .env file
Just create a file called `.env` in the root folder with these following parameters:
```
API_PORT=3000
DB_DRIVER=mysql
DB_HOST=localhost
DB_USER=user
DB_PASS=pass
DB_NAME=database_test
```
