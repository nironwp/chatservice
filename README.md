# ChatService microservice

The ChatService microservice interacts with ChatGPT through the OpenAI API, storing chat message history in MySQL.

## Instalando os pré-requisitos

Go 1.20\
https://go.dev/dl/

Docker\
https://docs.docker.com/get-docker/

sqlc
```bash
sudo snap install sqlc
```

migrate to Go
```bash
curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | sudo bash
sudo apt-get update
sudo apt-get install migrate
```

protoc\
https://grpc.io/docs/protoc-installation/

## How to run the project

Make a copy of the `env.example` file named `.env` inside the `chatservice` folder. Enter your OpenAI API Key in `OPENAI_API_KEY` inside the `.env` file. You can get an API Key from OpenAI [by clicking here](https://platform.openai.com/account/api-keys).

### Using Docker

```bash
cd folder/of/project/chatservice
docker-compose up -d
```
> *If you choose to run the chatservice microservice using Docker, be sure to change the value of DB_HOST to `DB_HOST=mysql` inside your `.env file`*

### Running locally

```bash
cd folder/of/project/chatservice
go run .\cmd\chatservice\main.go
```

> *If you choose to run the chatservice microservice locally, be sure to change the DB_HOST value to `DB_HOST=localhost` inside your `.env file`*

### migrate

In the first execution it will be necessary to apply the `migrate` to create the tables in the MySQL database, through the `Makefile`.

```bash
cd folder/of/project/chatservice
make migrate
```
> *When doing `make migrate` make sure that the MySQL connection string inside the `Makefile` file points to *mysql:3306* when the chatservice is running in Docker, or *localhost:3306* when the chatservice is running locally.*

### About the tiktoken-go package

For instructions on how to compile the [tiktoken-go](https://github.com/j178/tiktoken-go) package, watch Lecture 02 starting at [Minute 1:14](https://youtu.be/ lstRv2q-sOI?t=4446)

## Additional Information

On Windows, use the Ubuntu terminal with WSL 2 to run the commands.\
For more details, see Full Cycle's [WSL2 + Docker Quick Start](https://github.com/codeedu/wsl2-docker-quickstart).
