# email-dispatcher

A concurrent bulk email sender built with Go. Reads recipients from a CSV file and dispatches emails in parallel using a producer-consumer worker pool.

## How it works

1. **Producer** reads `data/emails.csv` and pushes recipients into a channel
2. **Workers** (default: 5) consume from the channel and send emails via SMTP concurrently

## Project structure

```
cmd/email-dispatcher/   Entry point
internal/
  config/               SMTP and app settings
  email/                Worker pool logic
  producer/             CSV reader
data/
  emails.csv            Recipient list (Name, Email)
```

## Usage

Start a local SMTP server (e.g. MailHog, MailPit) on port 1025, then:

```bash
go run ./cmd/email-dispatcher
```

## Configuration

Edit `internal/config/config.go` to change SMTP host, port, sender address, worker count, etc.
