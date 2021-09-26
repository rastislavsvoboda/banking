# Banking

from: REST based microservices API development in Go

[original sources](https://github.com/ashishjuyal/banking)

init module

```
go mod init github.com/rastislavsvoboda/banking
```

build

```
go build
```

run without build

```
$env:SERVER_ADDRESS="localhost"; $env:SERVER_PORT="8000"; $env:DB_USER="root"; $env:DB_PASSWORD="codecamp"; $env:DB_ADDRESS="localhost"; $env:DB_PORT="3306"; $env:DB_NAME="banking"; go run .\main.go
```

or after build

```
banking.exe
```

## resources

- https://hub.docker.com/_/mysql/
- https://github.com/gorilla/mux
- https://github.com/uber-go/zap
- https://github.com/jmoiron/sqlx
