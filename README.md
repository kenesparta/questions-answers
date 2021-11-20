# Question Answer Service

REST Service for questions and answers

# 1. Requirements

| Software         | Version | Importance                   |
| ---------------- | ------- | ---------------------------- |
| 🐳 Docker         | 20.10.11 | Required                     |
| 🐙 Docker Compose | 1.29.2  | Required                     |
| 🐃 GNU Make       | 4.2.1   | Optional                     |
| ‍🚀 Postman        | 9.1.5   | Optional                     |


# 2. Environment file (.env)

1. Copy the `.env-example` file to `.env` file:

```shell
cp .config-example .config
```

2. Fill out the variables with your own credentials
```
API_PORT=8087

POSTGRES_DB=mydb
POSTGRES_USER=user
POSTGRES_PASSWORD=secret
POSTGRES_DRIVER=postgres
POSTGRES_HOST=127.0.0.1
```

