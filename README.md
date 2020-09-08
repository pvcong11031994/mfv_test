# Test Money Forward VietNam

[![N|Solid](https://careers.moneyforward.vn/assets/logo-footer-b2812cede8e018c3938b930cafa1bc51efd3891bf1035014c7006e9de6e0772d.png)](https://careers.moneyforward.vn/)

### Requirement
- Implement a Rest API with CRUD functionality.
- Database: MySQL or PostgreSQL (Project current use PostgreSQL)
- Unit test as much as you can.
- Set up service with docker compse.

### Prepare
1. Install Docker
2. Install PostgreSQL 
3. VSCode
4. Data sample (Folder MFV_TEST/seed/seeder.go)

### New Features
- Get info user transaction by `user_id` and `user_account_id` (user_id is the ID of the table `users` and user_account_id is the ID of the table `user_accounts`)
- Insert one user transaction by user_id.
- Delete users by user_id (Delete soft)
- Update information users (The sample project is going to update the name)

### Note
Since the time is limited, the above source code will only handle happly cases.
Haven't handled auth cases (Source: SetMiddlewareAuthentication and auth/token.go is TODO).

### Installation
- Setup database:
Config file .env
```sh
# Postgres 
API_SECRET=98hbun98h #Used when creating a JWT. It can be anything
DB_HOST=127.0.0.1
DB_DRIVER=postgres
DB_USER=username
DB_PASSWORD=password
DB_NAME=postgres
DB_PORT=5432 #Default postgres port
```

```sh
$ go run main.go
```
- Sample CURL:

```sh
$ curl --location --request GET 'localhost:8080/api/users/1/transactions'
```

```sh
$ curl --location --request POST 'localhost:8080/api/users/1/transactions' \
--header 'Content-Type: application/json' \
--data-raw '{
  "account_id": 2,
  "amount": 100000.00,
  "transaction_type": "deposit"
}'
```

```sh
curl --location --request DELETE 'localhost:8080/api/users/1/transactions'
```

```sh
curl --location --request PUT 'localhost:8080/api/users/2/transactions' \
--header 'Content-Type: application/json' \
--data-raw '{
  "name": "CONGPV"
}'
```

### Document
1. https://docs.docker.com/get-started/part2/
2. https://medium.com/@hackintoshrao/hello-world-go-dockercompose-38e0f28618dc
3. https://levelup.gitconnected.com/dockerized-crud-restful-api-with-go-gorm-jwt-postgresql-mysql-and-testing-61d731430bd8