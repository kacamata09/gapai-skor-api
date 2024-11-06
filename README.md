my clean architecture using golang with echo inspired by bxcodec's go clean architecture
## setup

#### migrations
1. Create database as in your env
2. Execute this command in your database if you using postgresql, if you using mysql you don't need exec this
```CREATE EXTENSION IF NOT EXISTS "uuid-ossp";```
3. Exec migrations:
install migrate module mysql:
`go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest`
exec migration:
migrate -path ./migrations/mysql -database "mysql://root:password@tcp(127.0.0.1:3306)/dbname" up

