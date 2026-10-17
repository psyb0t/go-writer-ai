# Changelog

All notable changes per release. Versions follow [semver](https://semver.org).

## v0.1.2, 2026-10-17

- Documentation only. The agent skill explains how to look up every endpoint, parameter and model with `go doc` or on pkg.go.dev, and declares its license.
- The skill no longer lists the API credentials under OpenClaw's `requires.env`, which kept OpenClaw from loading it until they were set. It requires only `go`; the credentials are still documented as what your program reads.
- The README points to the full API reference, and the examples in the README and the skill were checked against the current code.

## v0.1.1, 2026-10-17

- Documentation only. The README and `AGENTS.md` describe the daily regeneration workflow: what it checks, how it fixes failing checks, how it picks the next version and which paths a fix may not touch.

## v0.1.0, 2026-10-16

- Added `writer`, a REST client generated with oapi-codegen from the OpenAPI document Writer publishes for its SDKs (Writer API, spec version 1.0, OpenAPI 3.0.3). It covers all 21 paths: chat, completions, models, files, Knowledge Graphs, no-code applications and their jobs, the PDF parser, web search and vision.
- Number fields are `float64` instead of oapi-codegen's default `float32`.
- `make generate` reads the current spec URL from the `.stats.yml` of writer/writer-python, downloads the spec from Stainless and regenerates the client. The spec is not committed.
- Added the `go-writer-ai` agent skill: `claude plugin install go-writer-ai@psyb0t` (Claude Code), `codex plugin add go-writer-ai@psyb0t` (Codex), `openclaw skills install @psyb0t/go-writer-ai` (ClawHub).
