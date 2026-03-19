# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project

Personal portfolio website built with Go (Gin), Templ templates, Tailwind CSS, Alpine.js, and HTMX. Deployed via Docker Swarm with Traefik reverse proxy.

## Development Commands

```bash
# Run locally (generates templ, then starts server)
make run

# Watch Tailwind CSS changes
npm run watch

# Generate templ templates manually
templ generate

# Run Go server directly
go run main.go -host localhost -port 8080 -static ./static

# Docker build
docker build -t mattcollecutt.com:latest .
```

The `-docker` flag binds to `0.0.0.0:8080` instead of `localhost:8080`.

## Architecture

- **main.go** — Gin server setup, routing, graceful shutdown. Routes: `/` (HTML), `/api/health`, `/api/ping`, `/static/*`
- **handler/** — HTTP handlers. `renderer.go` adapts Templ components to Gin's HTML renderer interface. `root.go` handles the main page
- **view/** — Templ templates organized as `layout/` (base HTML shell), `component/` (reusable UI pieces), `root/` (page content)
- **model/** — Data structures (media types for API responses)
- **static/** — Compiled CSS, JS libraries (Alpine.js, HTMX, lazysizes), images, favicon
- **input.css** — Tailwind source with custom animations/styles; compiles to `static/output.css`

## Key Patterns

- Templ generates `*_templ.go` files (gitignored) — always run `templ generate` after editing `.templ` files
- Multi-stage Dockerfile: fetch deps → generate templ → compile Go → build Tailwind → package in distroless
- Tailwind v3 with `input.css` as entry point

## Deployment

- CI/CD via `.github/workflows/deploy-service.yaml`
- Branch mapping: `dev` → dev, `master` → staging, release tag → prod
- Multi-platform builds (amd64 + arm64) pushed to GHCR
- Docker Swarm stack (`docker-stack.yaml`) with Traefik labels for routing and TLS
