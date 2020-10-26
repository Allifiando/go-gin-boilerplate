# Go Gin Boilerplate

> A starter project with Golang, Gin and MySql

[![Go Version][go-image]][go-url]

Golang Gin boilerplate with MySql resource.

![](header.jpg)

### Boilerplate structure

```
.
├── Makefile
├── README.md
├── controller
│   └── userController.go
├── helper
│   ├── pagination.go
│   ├── response.go
├── middleware
│   └── auth.go
│   └── cors.go
├── model
│   └── request
│       └── user_req.go
│   └── response
│       └── user_res.go
│   └── base.go
│   └── user.go
├── pkg
│   └── error
│       └── error.go
├── server
│   └── route.go
├── main.go
```

## Installation & Running

```
    - go mod vendor
    - make run-dev
```

## Usage example

`curl http://localhost:4000/`

[go-image]: https://img.shields.io/badge/Go--version-1.13-blue.svg
[go-url]: https://golang.org/doc/go1.13
