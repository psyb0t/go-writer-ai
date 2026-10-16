# Changelog

All notable changes per release. Versions follow [semver](https://semver.org).

## v0.1.0, 2026-10-16

- Added `writer`, a REST client generated with oapi-codegen from the OpenAPI document Writer publishes for its SDKs (Writer API, spec version 1.0, OpenAPI 3.0.3). It covers all 21 paths: chat, completions, models, files, Knowledge Graphs, no-code applications and their jobs, the PDF parser, web search and vision.
- Number fields are `float64` instead of oapi-codegen's default `float32`.
- `make generate` reads the current spec URL from the `.stats.yml` of writer/writer-python, downloads the spec from Stainless and regenerates the client. The spec is not committed.
- Added the `go-writer-ai` agent skill: `claude plugin install go-writer-ai@psyb0t` (Claude Code), `codex plugin add go-writer-ai@psyb0t` (Codex), `openclaw skills install @psyb0t/go-writer-ai` (ClawHub).
