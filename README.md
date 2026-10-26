# 👻 GhostOps

A lightweight Go CLI tool that analyzes log files, summarizes log statistics, and detects potential incidents using keyword-based rules.

## Features

- Reads log files line by line.
- Counts processed lines and malformed entries.
- Summarizes valid entries by severity and service.
- Detects incidents from `ERROR` and `FATAL` log messages.
- Groups incidents by incident type and affected service.
- Prints a report in the terminal and can optionally export it to JSON.
- Uses only the Go standard library.

## Requirements

- Go version compatible with the version declared in `go.mod`.

## Usage

Run the analysis from the repository root:

```bash
go run . sample.log
```

Save the report to JSON as well:

```bash
go run . sample.log -o report.json
```

Build an executable:

```bash
go build -o ghostops .
```

Then run it with the same arguments:

```bash
./ghostops sample.log
./ghostops sample.log -o report.json
```

The input file must be the first argument. The optional `-o` flag is followed by the output JSON filename.

## Input Format

Each valid log line is expected to have this structure:

```text
TIMESTAMP SEVERITY SERVICE MESSAGE
```

Example:

```text
2026-10-02T12:20:47Z FATAL user-service Database connection pool exhausted and crashed system
2026-10-02T10:22:03Z ERROR payment-service Payment gateway timeout - transaction aborted
```

- **Timestamp:** RFC3339 timestamp, such as `2026-10-02T12:20:47Z`.
- **Severity:** `INFO`, `WARN`, `ERROR`, or `FATAL`.
- **Service:** service name ending in `-service`, such as `payment-service`.
- **Message:** the rest of the line after the first three fields.

Malformed lines are counted and skipped so analysis can continue. The parser expects a timestamp, a recognized severity, a valid service name, and a message.

## Incident Detection

Incident detection uses predefined keywords. The message is matched case-insensitively against the incident rules. In the current implementation, only entries with `ERROR` or `FATAL` severity are checked for incidents.

Detected incidents are grouped by incident type and affected service. Each incident records its frequency and first/last observed timestamps. Example incident types include:

- `database_connection_failure`
- `payment_gateway_timeout`
- `service_crash`

Keyword matching is heuristic: a matching phrase indicates a potential incident, not proof of a root cause.

## Project Structure

```text
GhostOps/
├── analyzer/       # Streams through the file, updates statistics, aggregates incidents
├── models/
│   ├── incident/   # Incident model and matching logic
│   ├── logentry/   # Log entry and severity definitions
│   ├── report/     # Report model and terminal/JSON output
│   └── stat/       # Summary statistics model
├── parser/         # Parses and validates individual log lines
├── utils/          # Shared helpers and incident keywords
├── main.go         # CLI arguments and application entry point
├── sample.log      # Sample input log
└── data.json       # Example/generated JSON report
```

### Main Data Structures

- **`LogEntry`** represents a parsed log line: timestamp, severity, service, and message.
- **`Stat`** stores the number of processed lines, malformed entries, and counts by severity and service.
- **`Incident`** stores the incident type, observed severities, frequency, affected service, and time range.
- **`Report`** combines statistics and the list of incidents and provides terminal and JSON output methods.

## Development

Run the package checks with:

```bash
go test ./...
```

Add or expand tests for malformed input, empty files, overlapping incident keywords, out-of-order timestamps, and JSON output as the project evolves.

## Possible Improvements

- Validate the number of fields before indexing parsed log fields, and test lines with missing fields.
- Make incident-rule matching deterministic and decide explicitly whether one log entry may match multiple incident types.
- Sort statistics and incidents before output so reports have a stable order.
- Add unit tests for parsing, incident matching, aggregation, and JSON writing.
- Consider a small output interface only if alternative reporters or easier isolated testing become useful; the current project does not need interfaces just for their own sake.
