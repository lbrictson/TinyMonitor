# TinyMonitor
A headless monitoring system, designed to be simple to use and easy to maintain

### Running Locally

Starting the web server:

`TINYMONITOR_TESTING=true go run cmd/server/main.go`

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

Note:  CLI can be configured via a config file which is automatically sourced from `$HOME/.tinymonitor/config.json`

An example config file is below:

```
{
    "server_url": "http://localhost:8080",
    "username": "admin",
    "api_key": "aaaabbbbcccceeeedddd"
}
```

You can also set your CLI configuration with environment variables as well

```
TINYMONITOR_API_KEY="aaaabbbbcccceeeedddd"
TINYMONITOR_USERNAME=admin 
TINYMONITOR_SERVER_URL=http://localhost:8080
```

If both a config file and environment variable is present the environment variable will be selected

### Adding new models with ent

Run `go run -mod=mod entgo.io/ent/cmd/ent init $ModelNameHere`

Generate the ORM code with `go generate ./ent`