# Gin Base

A template of clean architecture with gin framework, go-fx as dependency container, gorm as orm for database related operations.

To learn about project structure and dependency injection please go through [here](https://medium.com/wesionary-team/dependency-injection-with-go-fx-b698a6585cf0?source=friends_link&sk=26f391ae41c493946ee3434be2ed4971)

## Running the project

- Make sure you have docker installed.
- Copy `.env.example` to `.env`
- Run `docker-compose up -d`
- Go to `localhost:8080` to verify if the server works.
- [Pgadmin](https://www.pgadmin.org/) Database Management runs at `5051` .
- [Adminer](https://www.adminer.org/) Database Management runs at `8081` 

If you are running without docker be sure database configuration is provided in `.env` file and run `go run . app:serve`

#### Environment Variables

<details>
    <summary>Variables Defined in the project </summary>

| Key            | Value                    | Desc                                        |
| -------------- | ------------------------ | ------------------------------------------- |
| `SERVER_PORT`  | `8080`                   | Port at which app runs                      |
| `ENV`          | `development,production` | App running Environment                     |
| `LOG_OUTPUT`   | `./server.log`           | Output Directory to save logs               |
| `LOG_LEVEL`    | `info`                   | Level for logging (check lib/logger.go:172) |
| `DB_USER`      | `username`               | Database Username                           |
| `DB_PASS`      | `password`               | Database Password                           |
| `DB_HOST`      | `0.0.0.0`                | Database Host                               |
| `DB_PORT`      | `5432` or `3306`         | Database Port                               |
| `DB_NAME`      | `test`                   | Database Name                               |
| `JWT_SECRET`   | `secret`                 | JWT Token Secret key                        |
| `ADMINER_PORT` | `5051` or `8081`         | Adminer DB Port                             |
| `DEBUG_PORT`   | `8082`                   | Port that delve debugger runs in            |

</details>

#### Migration Commands

> ⚓️ &nbsp; Add argument `p=host` if you want to run the migration runner from the host environment instead of docker environment.
> Check [#19](https://github.com/thinhhuynh/gin-base/issues/19) for more details. eg; `make p=host migrate-up`

<details>
    <summary>Migration commands available</summary>

| Command             | Desc                                           |
| ------------------- | ---------------------------------------------- |
| `make migrate-up`   | runs migration up command                      |
| `make migrate-down` | runs migration down command                    |
| `make force`        | Set particular version but don't run migration |
| `make goto`         | Migrate to particular version                  |
| `make drop`         | Drop everything inside database                |
| `make create`       | Create new migration file(up & down)           |

</details>

## Implemented Features

- Dependency Injection (go-fx)
- Routing (gin web framework)
- Environment Files
- Logging (file saving on `production`) [zap](https://github.com/uber-go/zap)
- Middlewares (cors)
- Database Setup (postgres or mysql)
- Models Setup and Automigrate (gorm)
- Repositories
- Implementing Basic CRUD Operation
- Authentication (JWT)
- Migration Runner Implementation
- Live code refresh
- Dockerize Application with Debugging Support Enabled. Debugger runs at `8082`. Vs code configuration is at `.vscode/launch.json` which will attach debugger to remote application. [Learn More](https://medium.com/wesionary-team/docker-debug-environment-for-go-and-gin-framework-36df80e061ac?source=friends_link&sk=35c9d856852944083dd30059200d87f0)
- Cobra Commander CLI Support. try: `go run . --help`

## Contributing

Please open issues if you want the template to add some features that is not in todos. 🙇‍♂️

Create a PR with relevant information if you want to contribute in this template.

## License
This project is licensed under the terms of the [MIT license](https://github.com/thinhhuynh/gin-base/blob/main/LICENSE).
