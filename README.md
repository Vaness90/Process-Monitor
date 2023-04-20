# Process Monitor

Process Monitor is a simple Golang application that monitors specified processes and exports their statuses as Prometheus metrics.

## Prerequisites

You need to have Golang installed on your machine. If you don't have it installed, follow the instructions on the [official Golang website](https://golang.org/doc/install).

You also need to install the following packages:

```bash
go get github.com/mitchellh/go-ps
go get github.com/prometheus/client_golang/prometheus
go get github.com/prometheus/client_golang/prometheus/promhttp
go get github.com/spf13/pflag
```
## Compilation

To compile the project, run the following command in the project directory:

```bash

go build -o process-monitor main.go
```

This will create an executable file named process-monitor.
## Usage

To run the compiled binary, use the following command:

```bash

./process-monitor -p <process_name> [-p <another_process_name>]
```
Replace <process_name> and <another_process_name> with the names of the processes you want to monitor.

For example:

```bash
./process-monitor -p yandex -p chrome
```
The application will start an HTTP server on port 8081 and expose the metrics at /metrics. You can then use tools like Prometheus and Grafana to scrape and visualize the metrics.
## Metrics

The application exposes the following metric:

    process_exists{process_name="<process_name>"}: Indicates whether the specified process is running (1) or not running (0).

## Contributing

If you'd like to contribute to this project, feel free to submit a pull request or report issues.