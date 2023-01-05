
# Clean Atchitecture With Golang

Clean architecture with gin framework, air hot reload library, gorm as orm for database related operations.

## Orverview

A sample project to learn about Golang, Docker and Clean Architecture.

## Requirement

- macOS(Apple silicon)
- Docker version 20.10.21


## Running the project

- Make sure you have docker installed.
- Run `git clone https://github.com/takeuchi-shogo/clean-architecture-golang.git`
- Copy .env.sample to .env
- Copy ./app/src/infrastructure/config/config.go.sample to ./app/src/infrastructure/config/config.go
- Run `docker-compose --env-file ./.env -f docker/docker-compose.yml up -d`
- Go to `localhost:8081`

## Description

### Project tree
<pre>
.
├── README.md
├── app
│   ├── air.toml
│   ├── bin
│   │   └── dev
│   ├── go.mod
│   ├── go.sum
│   ├── lib
│   │   └── env.go
│   ├── main.go
│   ├── src
│   │   ├── adapters
│   │   │   ├── controllers
│   │   │   │   ├── context.go
│   │   │   │   ├── h.go
│   │   │   │   ├── logger.go
│   │   │   │   └── product
│   │   │   │       └── users_controller.go
│   │   │   ├── gateways
│   │   │   │   ├── gateways
│   │   │   │   └── repositories
│   │   │   │       └── user_repository.go
│   │   │   └── presenters
│   │   │       └── product
│   │   │           └── users_presenter.go
│   │   ├── application
│   │   │   ├── repositories
│   │   │   │   └── user_repository.go
│   │   │   ├── usecases
│   │   │   │   └── product
│   │   │   │       └── user_interactor.go
│   │   │   └── utilities
│   │   │       └── format.go
│   │   ├── entities
│   │   │   ├── errors.go
│   │   │   ├── response.go
│   │   │   └── users.go
│   │   └── infrastructure
│   │       ├── config
│   │       │   ├── config.go
│   │       │   └── config.go.sample
│   │       ├── database
│   │       │   └── db.go
│   │       ├── middleware
│   │       │   ├── cors.go
│   │       │   ├── logger.go
│   │       │   └── request_handler.go
│   │       ├── route
│   │       │   ├── routing.go
│   │       │   └── user_routes.go
│   │       └── server
│   │           └── server.go
│   ├── tasks
│   └── tmp
├── docker
│   ├── app
│   │   └── Dockerfile
│   ├── docker-compose.yml
│   ├── mysql
│   │   └── Dockerfile
│   └── nginx
│       └── Dockerfile
├── mysql
│   ├── config
│   │   └── my.conf
│   ├── data
│   └── migrations
│       └── 01_init_db.sql
└── nginx
    └── config
        └── default.conf
</pre>


#### About /app
Folders related to the application itself.

* bin

	binary folder.

* lib

	load environment variables.

* /src

	Contains the project source code.

	- /adapters

		interface layer.

	- /applications

		business logic layer.

	- /entities

		domain layer.

	- /infrastructure

		infra layer.

* tasks

* tmp

* air.toml

	`https://github.com/cosmtrek/air` of setting file	

* go.mod, go.sum

	golang modules.

* main.go

	root file.


#### About /docker

Contains the docker.

#### About /mysql

Contains the database.

#### About /nginx

Contains the nginx.

#### About etc...

* .env

	Application-wide environment variable configuration file.


## TODO

。。。。。



## Author

[My Twitter Account](https://twitter.com/shogo_mthr123)
