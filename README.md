# TinyMonitor
A headless monitoring system

### Running Locally

Starting the web server:

`go run cmd/server/main.go`

Running the CLI

`go run cmd/cli/main.go`

Hot reload for the web server is provided through air

`air`

### Configuration

All configuration is done through environment variables, refer to the section below for specifics.

#### Server Configuration

Available environment variables and their defaults:

```
TINYMONITOR_TESTING=false
TINYMONITOR_PORT=8080
TINYMONITOR_DOMAIN=localhost
TINYMONITOR_AUTO_TLS=false
TINYMONITOR_LOG_LEVEL=info
TINYMONITOR_LOG_FORMAT=text
TINYMONITOR_DB_LOCATION=data/
```

#### CLI Configuration

Note:  CLI can be configured via a config file (TODO)

```
TINYMONITOR_API_KEY="aaaabbbbcccceeeedddd"
TINYMONITOR_USERNAME=admin 
TINYMONITOR_SERVER_URL=http://localhost:8080
```

### Adding new models with ent

Run `go run -mod=mod entgo.io/ent/cmd/ent init $MordelNameHere`

Generate the ORM code with `go generate ./ent`