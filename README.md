# TinyMonitor
A headless monitoring system

### Running Locally

Make sure to start the database

`docker run -it --name=tm-postgres -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=postgres -p 5432:5432 postgres:15`

Inspect the database if need be with `docker exec -it tm-postgres psql -U postgres postgres`
### Configuration

All configuration is done through environment variables, refer to the section below for specifics.

#### Server Configuration

Available environment variables and their defaults:

```
TINYMONITOR_PORT=8080
TINYMONITOR_DOMAIN=localhost
TINYMONITOR_AUTO_TLS=false
TINYMONITOR_LOG_LEVEL=info
TINYMONITOR_LOG_FORMAT=text
```


### Adding new models with ent

Run `go run -mod=mod entgo.io/ent/cmd/ent init $MordelNameHere`

Generate the ORM code with `go generate ./ent`