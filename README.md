# Question Answer Service

REST Service for questions and answers

## Database modeling

![model](model.svg)

# 1. Requirements

| Software         | Version | Importance                   |
| ---------------- | ------- | ---------------------------- |
| 🐳 Docker         | 20.10.11 | Required                     |
| 🐙 Docker Compose | 1.29.2  | Required                     |
| 🐃 GNU Make       | 4.2.1   | Optional                     |
| ‍🚀 Postman        | 9.1.5   | Optional                     |


# 2. Prepare the environment file (.env)

1. Copy the `.env-example` file to `.env` file:

```shell
cp .config-example .config
```

2. Fill out the variables with your own credentials, If you want to start quickly, 
these are the default values (do not modify):
```
API_PORT=8087

POSTGRES_DB=qas_db
POSTGRES_USER=qas
POSTGRES_PASSWORD=secret
POSTGRES_DRIVER=postgres
POSTGRES_HOST=qas_pgsql
```

# 3. Run the service
1. Please look at the `docker-compose.yml`

2. Execute the command `make l/up`, if you are on Windows execute:
```shell
docker-compose down --remove-orphans --rmi all
```
```shell
docker-compose --env-file ./.env up --detach --remove-orphans --force-recreate --build
```

3. Execute the **sql scripts** on the `./sql` directory by these commands: