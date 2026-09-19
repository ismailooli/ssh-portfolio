# SSH Portfolio

A terminal portfolio served over SSH.

## Project layout

- `cmd/ssh-portfolio`: application entry point.
- `internal/server`: Wish SSH server configuration and startup.
- `internal/tui`: Bubble Tea model, rendering, and embedded UI assets.

## Run locally

```sh
go run ./cmd/ssh-portfolio
```

Then connect with `ssh -p 2222 localhost`.
