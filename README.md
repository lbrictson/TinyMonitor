# TinyMonitor
A headless monitoring system, designed to be simple to use and easy to maintain

### Running Locally with Docker-Compose

`docker-compose up -d --build`

If this is your first time launching get your API key from the logs.  It's only shown once

`docker-compose logs tinymonitor`

Create your CLI config file in `~/.tinymonitor/config.json`

```json
{
"server_url": "http://localhost:8080",
"username": "admin",
"api_key": "API-Key-From-Logs"
}
```

Build the CLI

`go build -o=tmp/tinymonitor cmd/tinymonitor/main.go`

Interact with the Service

`tmp/tinymonitor monitor list`

Create a metric sink for the influxdb container

`tmp/tinymonitor sink apply -f example/inflxudb_v1_sink.json`

Create some sample monitors

```
tmp/tinymonitor monitor apply -f example/simple_browser.json
tmp/tinymonitor monitor apply -f example/simple_http.json
tmp/tinymonitor monitor apply -f example/simple_ping.json
```

Optional:  Setup InfluxDB in Grafana

1.  Navigate to localhost:3000 (admin/admin)
2.  Add a new datasource:  Type Influxdb, host: http://influxdb:8086, database: test


### Running Locally (Non Docker)

Starting the web server:

`TINYMONITOR_TESTING=true go run cmd/tinymonitor/main.go server`

Running the CLI

`go run cmd/tinymonitor/main.go`

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