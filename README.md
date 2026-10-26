# GhostOps

GhostOps is a command-line log analyzer written in Go. It reads a log file line by line, collects statistics about the entries, and groups `ERROR` and `FATAL` entries into incidents using keyword rules. The result is printed to the console and can also be saved as a JSON file.

## Features

- Counts total entries and malformed entries.
- Counts entries per service and per severity level.
- Skips lines that do not match the expected format and counts them as malformed.
- Detects incidents in `ERROR` and `FATAL` entries by case-insensitive keyword matching.
- Merges incidents with the same type and service, tracking frequency, severities, and first and last occurrence.
- Prints a report to the console.
- Optionally writes the report to a JSON file with the `-o` flag.
- Uses only the Go standard library.

## Requirements

- Go 1.27.2 or newer, as declared in `go.mod`.

## Running the program

Run from the repository root:

```bash
go run . sample.log
```

To also save the report as JSON:

```bash
go run . sample.log -o report.json
```

To build a binary:

```bash
go build -o ghostops .
./ghostops sample.log -o report.json
```

Arguments:

- The input log file must be the first argument.
- `-o <file>` is optional and writes the report as indented JSON to `<file>`.

The program prints the input file name, then the statistics and the list of incidents. `data.json` in the repository is an example JSON report produced from `sample.log`.

## Input format

Each line is split on whitespace and must have this form:

```text
<timestamp> <severity> <service> <message>
2026-10-02T10:22:03Z ERROR payment-service Payment gateway timeout - transaction aborted
```

- `timestamp` must be in RFC 3339 format, for example `2026-10-02T10:22:03Z`.
- `severity` must be one of the values in `logentry.SeverityMap`: `INFO`, `WARN`, `ERROR`, `FATAL` (or `INVALID_LEVEL`).
- `service` must end with `-service`. The suffix is removed in the report, so `payment-service` is reported as `payment`.
- `message` is the rest of the line.

Lines that do not follow this format are counted as malformed and skipped.

### Incident rules

Incident types and their keywords are defined in `utils/rules.go`:

- `database_connection_failure`: `database connection`, `connection pool exhausted`
- `payment_gateway_timeout`: `payment gateway timeout`, `transaction aborted`
- `authentication_failure`: `authentication failed`, `login failed`, `unauthorized`
- `network_timeout`: `connection timeout`, `request timeout`, `i/o timeout`
- `service_crash`: `crashed`, `panic`, `fatal error`
- `resource_exhaustion`: `out of memory`, `resource exhausted`, `too many open files`

Only entries with severity `ERROR` or `FATAL` are checked. An entry that contains none of the keywords is not reported as an incident.

## Project structure

```text
GhostOps/
├── main.go                     Entry point and command-line argument handling
├── go.mod                      Module definition (sg/ghostops)
├── sample.log                  Sample input log
├── data.json                   Example JSON report generated from sample.log
├── analyzer/
│   └── analyzer.go             FileAnalysis: reads the file and builds the report
├── parser/
│   ├── parser.go               ParseLine: converts one line into a LogEntry
│   └── errors.go               Parsing error values
├── models/
│   ├── logentry/log_entry.go   LogEntry, SeverityEnum, and severity maps
│   ├── incident/incident.go    Incident, MatchIncident, and console output
│   ├── stat/stat.go            Stat counters and console output
│   └── report/report.go        Report, JSON output, console output, and lookup helpers
└── utils/
    ├── utils.go                Service name extraction and time comparison helpers
    └── rules.go                Incident keyword rules
```

Processing flow:

1. `main.go` reads the arguments and calls `analyzer.FileAnalysis`.
2. `FileAnalysis` reads the file line by line and calls `parser.ParseLine` for each line.
3. Valid entries update the statistics. `ERROR` and `FATAL` entries are passed to `incident.MatchIncident`.
4. Matched incidents are added to the report or merged with an existing incident of the same type and service.
5. `main.go` prints the report with `ShowResultCL` and, if `-o` is given, saves it with `WriteJson`.

