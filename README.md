# TinyMonitor
A headless monitoring system

### Running Locally

`go run cmd/server/main.go`

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
TINYMONITOR_DB_LOCATION=data/
```


### Adding new models with ent

Run `go run -mod=mod entgo.io/ent/cmd/ent init $MordelNameHere`

Generate the ORM code with `go generate ./ent`