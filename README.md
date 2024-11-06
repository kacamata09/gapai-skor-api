## setup

### Environment configuration
You can settings your environment in env.yaml  
### Database migration
1. Create database as in your env.yaml
2. Run migration with following docs in https://github.com/golang-migrate/migrate  
   Example :  
```
migrate -path ./migrations/mysql -database "mysql://root:password@tcp(127.0.0.1:3306)/dbname" up
```
3. You can change database to mysql if change package driver sql and config_connection at main.go  

### Run App
```
go run main.go
```
